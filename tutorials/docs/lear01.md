## learn01

### user-server
#### 生成代码
```
micro new --namespace=smtl.micro.learn --type=srv --alias=user micro_learn/tutorials/learn1/user_server
```

#### proto
user-server的proto，提供的rpc服务

```
syntax = "proto3";

package smtl.micro.learn.srv.user;

service User {
	rpc QueryUserByName (Request) returns (Response) {
	}
}

message user {
	int64 id = 1;
	string name = 2;
	string pwd = 3;
	uint64 createdTime = 4;
	uint64 updatedTime = 5;
}

message Error {
	int32 code = 1;
	string detail = 2;
}

message Request {
	string userID = 1;
	string userName = 2;
	string userPwd = 3;
}

message Response {
	bool success = 1;
	Error error = 2;
	user user = 3;
}
```

生成micro的rpc文件

```
protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto
```
#### sql
```
CREATE TABLE `user`
(
    `id`           int(10) unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned                                                       DEFAULT NULL COMMENT '用户id',
    `user_name`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户名',
    `pwd`          varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
    `created_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_user_name_uindex` (`user_name`),
    UNIQUE KEY `user_user_id_uindex` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户表';
```
### user-web
#### 生成代码
命令生成代码
```
micro new --namespace=smtl.micro.learn --type=web --alias=user micro_learn/tutorials/learn1/user_web
```

主要handler实现

```
func UserLogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	userName := r.Form.Get("userName")

	// call the backend service
	userClient := user.NewUserService("smtl.micro.learn.srv.user", client.DefaultClient)
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
```

main 函数
```
func main() {

	// register
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{"127.0.0.1:2379"}
	})

	// create new web service
	service := web.NewService(
		web.Name("smtl.micro.learn.web.user"),
		web.Version("latest"),
		web.Registry(reg),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/user/login", handler.UserLogIn)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### 测试
启动user-server
```
go run main.go
# 启动成功信息
019-07-31 15:27:44.497173 I | Transport [http] Listening on [::]:56538
2019-07-31 15:27:44.497324 I | Broker [http] Connected to [::]:56540
2019-07-31 15:27:44.497995 I | Registry [etcdv3] Registering node: smtl.micro.learn.srv.user-a3824072-0e24-4c47-85f0-78f5d69bbf65

```
启动user-web
```
go run main.go
# 启动成功信息
2019-07-31 15:08:17.800764 I | Listening on [::]:54469
```
curl测试
```
curl --request POST   --url http://127.0.0.1:54469/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=1234'
{"error":{"detail":"密码错误"},"ref":1564558076031178000,"success":false}


curl --request POST   --url http://127.0.0.1:54469/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=123'
{"data":{"id":10001,"name":"micro","createdTime":1564470621,"updatedTime":1564470621},"ref":1564558145349977000,"success":true}
```

### micro命令 支持etcdv3
直接go get下来的micro命令不支持etcdv3
```
micro --registry=etcdv3 call smtl.micro.learn.srv.user User.QueryUserByName '{"userName":"micro"}'
命令一直不运行，原因是micro默认不支持etcdv3，需要手动重新编译

```
#### 步骤
- 进入micro项目目录，建立一个plugin.go 
```
package main

import (
	_ "github.com/micro/go-plugins/registry/etcdv3"
)

```
- 然后重新编译
```
go build  -o micro main.go plugin.go
```
- 拷贝到GOPATH/BIN下
```
copy micro $GOAPTH/bin
```