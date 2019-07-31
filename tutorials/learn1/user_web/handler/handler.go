package handler

import (
	"context"
	"encoding/json"
	"micro_learn/tutorials/learn1/user_web/client"
	"net/http"
	"time"

	user "micro_learn/tutorials/learn1/user_server/proto/user"
)

func UserLogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	userName := r.Form.Get("userName")

	// call the backend service
	userClient := client.GetUserService()
	rsp, err := userClient.QueryUserByName(context.TODO(), &user.Request{
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
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
