package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	r "github.com/go-redis/redis"
	"micro_learn/tutorials/learn2/basic"
	"sync"
)

var (
	s  *service
	ca *r.Client
	m  sync.RWMutex
)

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}


// 把jwt的claim转成claims
func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{
		Subject: claims["sub"].(string),
	}

	return jC
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	ca = basic.GetRedisClient()

	s = &service{}
}
