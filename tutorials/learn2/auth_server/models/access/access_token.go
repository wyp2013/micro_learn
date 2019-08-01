package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"time"
)

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"

	tokenExpiredTopic = "smtl.micro.learn.topic.auth.tokenExpired"
)

// Subject token 持有者
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}


// service 服务
type service struct {
}

// createTokenClaims Claims
func (s *service) createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.ID,
		IssuedAt:  now.Unix(),
		Issuer:    "book.micro.mu",
		Subject:   subject.ID,
	}

	return
}

// saveTokenToCache 保存token到缓存
func (s *service) saveTokenToCache(subject *Subject, val string) (err error) {
	//保存
	if err = ca.Set(tokenIDKeyPrefix+subject.ID, val, tokenExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveTokenToCache] 保存token到缓存发生错误，err:" + err.Error())
	}
	return
}

// delTokenFromCache 清空token
func (s *service) delTokenFromCache(subject *Subject) (err error) {
	//保存
	if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
		return fmt.Errorf("[delTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
	}
	return
}

// getTokenFromCache 从缓存获取token
func (s *service) getTokenFromCache(subject *Subject) (token string, err error) {
	// 获取
	tokenCached, err := ca.Get(tokenIDKeyPrefix + subject.ID).Result()
	if err != nil {
		return token, fmt.Errorf("[getTokenFromCache] token不存在 %s", err)
	}

	return string(tokenCached), nil
}

// parseToken 解析token
func (s *service) parseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("不合法的token格式: %v", token.Header["alg"])
		}
		return []byte(GetSecretKey()), nil
	})

	// jwt 框架自带了一些检测，如过期，发布者错误等
	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] 不合法的token, err:%s", err)
	}

	// 检测合法
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] 不合法的token")
	}

	return mapClaimToJwClaim(claims), nil
}

// MakeAccessToken 生成token并保存到redis
func (s *service) MakeAccessToken(subject *Subject) (ret string, err error) {
	m, err := s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token Claim 失败，err: %s", err)
	}

	// 创建
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	ret, err = token.SignedString([]byte(GetSecretKey()))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token失败，err: %s", err)
	}

	// 保存到redis
	err = s.saveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}

	return
}

// GetCachedAccessToken 获取token
func (s *service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}

	return
}

// DelUserAccessToken 清除用户token
func (s *service) DelUserAccessToken(tk string) (err error) {
	// 解析token字符串
	claims, err := s.parseToken(tk)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
	}

	// 通过解析到的用户id删除
	err = s.delTokenFromCache(&Subject{
		ID: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
	}

	// 广播删除
	msg := &broker.Message{
		Body: []byte(claims.Subject),
	}
	if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
		log.Logf("[pub] 发布token删除消息失败： %v", err)
	} else {
		fmt.Println("[pub] 发布token删除消息：", string(msg.Body))
	}

	return
}

func GetSecretKey() string {
	return "W6VjDud2W1kMG3BicbMNlGgI4ZfcoHtMGLWr"
}
