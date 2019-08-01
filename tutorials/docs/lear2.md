## learn02

### auth-server
#### 生成代码
```
micro new --namespace=smtl.micro.learn --type=srv --alias=aurh micro_learn/tutorials/learn2/auth_server
```

#### proto
auth-server的proto，提供的rpc服务

```
protoc --proto_path=. --go_out=. --micro_out=. proto/auth/auth.proto
```
### 测试一下 auth服务
```
micro --registry=etcdv3  call  smtl.micro.learn.srv.auth Auth.MakeAccessToken '{"userName":"micro", "userId":"10001"}'
{
	"success": true,
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjcyNDQ2ODgsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTY0NjUyNjg4LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTY0NjUyNjg4LCJzdWIiOiIxMDAwMSJ9.fjKqChHHXA7aWXryhQz-aBpBV6uScW4qmjyY4Qe-IqI"
}
```
### 测试web
启动auth-server
```
go run main.go
```
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


curl --request POST   --url http://127.0.0.1:8091/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=123'
{"data":{"id":10001,"name":"micro","createdTime":1564470621,"updatedTime":1564470621},"ref":1564655251421648000,"success":true,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjcyNDcyNTEsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTY0NjU1MjUxLCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTY0NjU1MjUxLCJzdWIiOiIxMDAwMSJ9.3I4jnCTzjv7yyxAqJ7tCjTeVQ357FZM3sLScmmorTIs"}
```