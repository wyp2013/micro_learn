package handler

import (
	"context"
	"encoding/json"
	"micro_learn/tutorials/learn2/user_web/client"
	"net/http"
	"time"

	"github.com/micro/go-micro/util/log"
	auth "micro_learn/tutorials/learn2/auth_server/proto/auth"
	user "micro_learn/tutorials/learn2/user_server/proto/user"
)

func UserLogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	userName := r.Form.Get("userName")

	// call the backend service
	ctx := context.TODO()
	userClient := client.GetUserService()
	rsp, err := userClient.QueryUserByName(ctx, &user.Request{
		UserName: userName,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Pwd != r.Form.Get("pwd") {
		response["success"] = false
		response["error"] = &user.Error{
			Detail: "密码错误",
		}
	} else {
		response["success"] = true
		rsp.User.Pwd = ""
		response["data"] = rsp.User

		authClient := client.GetAuthService()
		arsp, err := authClient.MakeAccessToken(ctx, &auth.Request{UserId: uint64(rsp.User.Id), UserName: userName})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		response["token"] = arsp.Token
		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: arsp.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
