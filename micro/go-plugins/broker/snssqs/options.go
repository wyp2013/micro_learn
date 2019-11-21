package snssqs

import (
	"micro_learn/micro/go-micro/broker"
)

type maxMessagesKey struct{}

// MaxReceiveMessages indicates how many messages a receive operation should pull
// during any single call
func MaxReceiveMessages(max int64) broker.SubscribeOption {
	return setSubscribeOption(maxMessagesKey{}, max)
}

type visibilityTimeoutKey struct{}

// VisibilityTimeout controls how long a message is hidden from other queue consumers
// before being put back. If a consumer does not delete the message, it will be put back
// even if it was "processed"
func VisibilityTimeout(seconds int64) broker.SubscribeOption {
	return setSubscribeOption(visibilityTimeoutKey{}, seconds)
}

type waitTimeSecondsKey struct{}

// WaitTimeSeconds controls the length of long polling for available messages
func WaitTimeSeconds(seconds int64) broker.SubscribeOption {
	return setSubscribeOption(waitTimeSecondsKey{}, seconds)
}

type validateOnPublishKey struct{}

// ValidateOnPublish determines whether to pre-validate messages before they're published
// This has a significant performance impact
func ValidateOnPublish(validate bool) broker.PublishOption {
	return setPublishOption(validateOnPublishKey{}, validate)
}
