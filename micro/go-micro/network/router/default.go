package router

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"micro_learn/micro/go-micro/network/router/table"
	"micro_learn/micro/go-micro/registry"
)

const (
	// AdvertiseEventsTick is time interval in which the router advertises route updates
	AdvertiseEventsTick = 5 * time.Second
	// AdvertiseTableTick is time interval in which router advertises all routes found in routing table
	AdvertiseTableTick = 1 * time.Minute
	// AdvertSuppress is advert suppression threshold
	AdvertSuppress = 2000.0
	// AdvertRecover is advert recovery threshold
	AdvertRecover = 750.0
	// DefaultAdvertTTL is default advertisement TTL
	DefaultAdvertTTL = 1 * time.Minute
	// DeletePenalty penalises route deletion
	DeletePenalty = 1000.0
	// UpdatePenalty penalises route updates
	UpdatePenalty = 500.0
	// PenaltyHalfLife is the time the advert penalty decays to half its value
	PenaltyHalfLife = 2.0
	// MaxSuppressTime defines time after which the suppressed advert is deleted
	MaxSuppressTime = 5 * time.Minute
)

var (
	// PenaltyDecay is a coefficient which controls the speed the advert penalty decays
	PenaltyDecay = math.Log(2) / PenaltyHalfLife
)

// router provides default router implementation
type router struct {
	// embed the table
	table.Table
	opts       Options
	status     Status
	exit       chan struct{}
	eventChan  chan *table.Event
	advertChan chan *Advert
	advertWg   *sync.WaitGroup
	wg         *sync.WaitGroup
	sync.RWMutex
}

// newRouter creates a new router and returns it
func newRouter(opts ...Option) Router {
	// get default options
	options := DefaultOptions()

	// apply requested options
	for _, o := range opts {
		o(&options)
	}

	return &router{
		Table:      options.Table,
		opts:       options,
		status:     Status{Error: nil, Code: Stopped},
		exit:       make(chan struct{}),
		eventChan:  make(chan *table.Event),
		advertChan: make(chan *Advert),
		advertWg:   &sync.WaitGroup{},
		wg:         &sync.WaitGroup{},
	}
}

// Init initializes router with given options
func (r *router) Init(opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}
	return nil
}

// Options returns router options
func (r *router) Options() Options {
	return r.opts
}

// manageRoute applies route action on the routing table
func (r *router) manageRoute(route table.Route, action string) error {
	switch action {
	case "create":
		if err := r.Create(route); err != nil && err != table.ErrDuplicateRoute {
			return fmt.Errorf("failed adding route for service %s: %s", route.Service, err)
		}
	case "update":
		if err := r.Update(route); err != nil && err != table.ErrDuplicateRoute {
			return fmt.Errorf("failed updating route for service %s: %s", route.Service, err)
		}
	case "delete":
		if err := r.Delete(route); err != nil && err != table.ErrRouteNotFound {
			return fmt.Errorf("failed deleting route for service %s: %s", route.Service, err)
		}
	default:
		return fmt.Errorf("failed to manage route for service %s. Unknown action: %s", route.Service, action)
	}

	return nil
}

// manageServiceRoutes manages routes for a given service.
// It returns error of the routing table action fails.
func (r *router) manageServiceRoutes(service *registry.Service, action string) error {
	// action is the routing table action
	action = strings.ToLower(action)

	// take route action on each service node
	for _, node := range service.Nodes {
		route := table.Route{
			Service: service.Name,
			Address: node.Address,
			Gateway: "",
			Network: r.opts.Network,
			Link:    table.DefaultLink,
			Metric:  table.DefaultLocalMetric,
		}

		if err := r.manageRoute(route, action); err != nil {
			return err
		}
	}

	return nil
}

// manageRegistryRoutes manages routes for each service found in the registry.
// It returns error if either the services failed to be listed or the routing table action fails.
func (r *router) manageRegistryRoutes(reg registry.Registry, action string) error {
	services, err := reg.ListServices()
	if err != nil {
		return fmt.Errorf("failed listing services: %v", err)
	}

	// add each service node as a separate route
	for _, service := range services {
		// get the service to retrieve all its info
		srvs, err := reg.GetService(service.Name)
		if err != nil {
			continue
		}
		// manage the routes for all returned services
		for _, srv := range srvs {
			if err := r.manageServiceRoutes(srv, action); err != nil {
				return err
			}
		}
	}

	return nil
}

// watchRegistry watches registry and updates routing table based on the received events.
// It returns error if either the registry watcher fails with error or if the routing table update fails.
func (r *router) watchRegistry(w registry.Watcher) error {
	// wait in the background for the router to stop
	// when the router stops, stop the watcher and exit
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		<-r.exit
		w.Stop()
	}()

	var watchErr error

	for {
		res, err := w.Next()
		if err != nil {
			if err != registry.ErrWatcherStopped {
				watchErr = err
			}
			break
		}

		if err := r.manageServiceRoutes(res.Service, res.Action); err != nil {
			return err
		}
	}

	return watchErr
}

// watchTable watches routing table entries and either adds or deletes locally registered service to/from network registry
// It returns error if the locally registered services either fails to be added/deleted to/from network registry.
func (r *router) watchTable(w table.Watcher) error {
	// wait in the background for the router to stop
	// when the router stops, stop the watcher and exit
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		<-r.exit
		w.Stop()
	}()

	var watchErr error

	for {
		event, err := w.Next()
		if err != nil {
			if err != table.ErrWatcherStopped {
				watchErr = err
			}
			break
		}

		select {
		case <-r.exit:
			close(r.eventChan)
			return nil
		case r.eventChan <- event:
		}
	}

	// close event channel on error
	close(r.eventChan)

	return watchErr
}

// advertiseEvents advertises events to event subscribers
func (r *router) advertiseEvents(advType AdvertType, events []*table.Event) {
	defer r.advertWg.Done()

	a := &Advert{
		Id:        r.opts.Id,
		Type:      advType,
		TTL:       DefaultAdvertTTL,
		Timestamp: time.Now(),
		Events:    events,
	}

	select {
	case r.advertChan <- a:
	case <-r.exit:
		return
	}
}

// advertiseTable advertises the whole routing table to the network
func (r *router) advertiseTable() error {
	// create table advertisement ticker
	ticker := time.NewTicker(AdvertiseTableTick)

	for {
		select {
		case <-ticker.C:
			// list routing table routes to announce
			routes, err := r.List()
			if err != nil {
				return fmt.Errorf("failed listing routes: %s", err)
			}
			// collect all the added routes before we attempt to add default gateway
			events := make([]*table.Event, len(routes))
			for i, route := range routes {
				event := &table.Event{
					Type:      table.Update,
					Timestamp: time.Now(),
					Route:     route,
				}
				events[i] = event
			}

			// advertise all routes as Update events to subscribers
			if len(events) > 0 {
				r.advertWg.Add(1)
				go r.advertiseEvents(Update, events)
			}
		case <-r.exit:
			return nil
		}
	}
}

// routeAdvert contains a list of route events to be advertised
type routeAdvert struct {
	events []*table.Event
	// lastUpdate records the time of the last advert update
	lastUpdate time.Time
	// penalty is current advert penalty
	penalty float64
	// isSuppressed flags the advert suppression
	isSuppressed bool
	// suppressTime records the time interval the advert has been suppressed for
	suppressTime time.Time
}

// processEvents processes routing table events.
// It suppresses unhealthy flapping events and advertises healthy events upstream.
func (r *router) processEvents() error {
	// ticker to periodically scan event for advertising
	ticker := time.NewTicker(AdvertiseEventsTick)
	// advertMap is a map of advert events
	advertMap := make(map[uint64]*routeAdvert)

	for {
		select {
		case <-ticker.C:
			var events []*table.Event
			// collect all events which are not flapping
			for key, advert := range advertMap {
				// decay the event penalty
				delta := time.Since(advert.lastUpdate).Seconds()
				advert.penalty = advert.penalty * math.Exp(-delta*PenaltyDecay)

				// suppress/recover the event based on its penalty level
				switch {
				case advert.penalty > AdvertSuppress && !advert.isSuppressed:
					advert.isSuppressed = true
					advert.suppressTime = time.Now()
				case advert.penalty < AdvertRecover && advert.isSuppressed:
					advert.isSuppressed = false
				}

				// max suppression time threshold has been reached, delete the advert
				if advert.isSuppressed {
					if time.Since(advert.suppressTime) > MaxSuppressTime {
						delete(advertMap, key)
						continue
					}
				}

				if !advert.isSuppressed {
					for _, event := range advert.events {
						e := new(table.Event)
						*e = *event
						events = append(events, e)
						// delete the advert from the advertMap
						delete(advertMap, key)
					}
				}
			}

			// advertise all Update events to subscribers
			if len(events) > 0 {
				r.advertWg.Add(1)
				go r.advertiseEvents(Update, events)
			}
		case e := <-r.eventChan:
			// if event is nil, continue
			if e == nil {
				continue
			}

			// determine the event penalty
			var penalty float64
			switch e.Type {
			case table.Update:
				penalty = UpdatePenalty
			case table.Delete:
				penalty = DeletePenalty
			}

			// check if we have already registered the route
			// we use the route hash as advertMap key
			hash := e.Route.Hash()
			advert, ok := advertMap[hash]
			if !ok {
				events := []*table.Event{e}
				advert = &routeAdvert{
					events:     events,
					penalty:    penalty,
					lastUpdate: time.Now(),
				}
				advertMap[hash] = advert
				continue
			}

			// attempt to squash last two events if possible
			lastEvent := advert.events[len(advert.events)-1]
			if lastEvent.Type == e.Type {
				advert.events[len(advert.events)-1] = e
			} else {
				advert.events = append(advert.events, e)
			}

			// update event penalty and recorded timestamp
			advert.lastUpdate = time.Now()
			advert.penalty += penalty

		case <-r.exit:
			// first wait for the advertiser to finish
			r.advertWg.Wait()
			// close the advert channel
			close(r.advertChan)
			return nil
		}
	}
}

// watchErrors watches router errors and takes appropriate actions
func (r *router) watchErrors(errChan <-chan error) {
	defer r.wg.Done()

	var code StatusCode
	var err error

	select {
	case <-r.exit:
		code = Stopped
	case err = <-errChan:
		code = Error
	}

	r.Lock()
	defer r.Unlock()
	status := Status{
		Code:  code,
		Error: err,
	}
	r.status = status

	// stop the router if some error happened
	if err != nil && code != Stopped {
		// this will stop watchers which will close r.advertChan
		close(r.exit)
		// drain the advertise channel
		for range r.advertChan {
		}
		// drain the event channel
		for range r.eventChan {
		}
	}

}

// Advertise advertises the routes to the network.
// It returns error if any of the launched goroutines fail with error.
func (r *router) Advertise() (<-chan *Advert, error) {
	r.Lock()
	defer r.Unlock()

	if r.status.Code != Running {
		// add all local service routes into the routing table
		if err := r.manageRegistryRoutes(r.opts.Registry, "create"); err != nil {
			return nil, fmt.Errorf("failed adding routes: %s", err)
		}

		// list routing table routes to announce
		routes, err := r.List()
		if err != nil {
			return nil, fmt.Errorf("failed listing routes: %s", err)
		}
		// collect all the added routes before we attempt to add default gateway
		events := make([]*table.Event, len(routes))
		for i, route := range routes {
			event := &table.Event{
				Type:      table.Create,
				Timestamp: time.Now(),
				Route:     route,
			}
			events[i] = event
		}

		// add default gateway into routing table
		if r.opts.Gateway != "" {
			// note, the only non-default value is the gateway
			route := table.Route{
				Service: "*",
				Address: "*",
				Gateway: r.opts.Gateway,
				Network: "*",
				Metric:  table.DefaultLocalMetric,
			}
			if err := r.Create(route); err != nil {
				return nil, fmt.Errorf("failed adding default gateway route: %s", err)
			}
		}

		// NOTE: we only need to recreate these if the router errored or was stopped
		// TODO: These probably dont need to be struct members
		if r.status.Code == Error || r.status.Code == Stopped {
			r.exit = make(chan struct{})
			r.eventChan = make(chan *table.Event)
			r.advertChan = make(chan *Advert)
		}

		// routing table watcher
		tableWatcher, err := r.Watch()
		if err != nil {
			return nil, fmt.Errorf("failed creating routing table watcher: %v", err)
		}

		// registry watcher
		regWatcher, err := r.opts.Registry.Watch()
		if err != nil {
			return nil, fmt.Errorf("failed creating service registry watcher: %v", err)
		}

		// error channel collecting goroutine errors
		errChan := make(chan error, 4)

		r.wg.Add(1)
		go func() {
			defer r.wg.Done()
			// watch local registry and register routes in routine table
			errChan <- r.watchRegistry(regWatcher)
		}()

		r.wg.Add(1)
		go func() {
			defer r.wg.Done()
			// watch local registry and register routes in routing table
			errChan <- r.watchTable(tableWatcher)
		}()

		r.wg.Add(1)
		go func() {
			defer r.wg.Done()
			// watch routing table events and process them
			errChan <- r.processEvents()
		}()

		r.advertWg.Add(1)
		go func() {
			defer r.advertWg.Done()
			// advertise the whole routing table
			errChan <- r.advertiseTable()
		}()

		// advertise your presence
		r.advertWg.Add(1)
		go r.advertiseEvents(Announce, events)

		// watch for errors and cleanup
		r.wg.Add(1)
		go r.watchErrors(errChan)

		// mark router as running and set its Error to nil
		r.status = Status{Code: Running, Error: nil}
	}

	return r.advertChan, nil
}

// Process updates the routing table using the advertised values
func (r *router) Process(a *Advert) error {
	// NOTE: event sorting might not be necessary
	// copy update events intp new slices
	events := make([]*table.Event, len(a.Events))
	copy(events, a.Events)
	// sort events by timestamp
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Before(events[j].Timestamp)
	})

	for _, event := range events {
		// create a copy of the route
		route := event.Route
		action := event.Type
		if err := r.manageRoute(route, fmt.Sprintf("%s", action)); err != nil {
			return fmt.Errorf("failed applying action %s to routing table: %s", action, err)
		}
	}

	return nil
}

// Status returns router status
func (r *router) Status() Status {
	r.RLock()
	defer r.RUnlock()

	// make a copy of the status
	status := r.status

	return status
}

// Stop stops the router
func (r *router) Stop() error {
	r.RLock()
	// only close the channel if the router is running
	if r.status.Code == Running {
		// notify all goroutines to finish
		close(r.exit)
		// drain the advertise channel
		for range r.advertChan {
		}
		// drain the event channel
		for range r.eventChan {
		}
	}
	r.RUnlock()

	// wait for all goroutines to finish
	r.wg.Wait()

	return nil
}

// String prints debugging information about router
func (r *router) String() string {
	return "default"
}
