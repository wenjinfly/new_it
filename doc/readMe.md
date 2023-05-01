go mod tidy

这是一个新项目基于service weaver

weaver generate .
weaver generate ./app/configApp/
weaver generate ./app/usercentApp/

继承过weaver里边的东西都需要生成,比如configAPP就需要执行两次
一次是：weaver generate ./app/configApp/model
另外是：weaver generate ./app/configApp/

weaver generate ./app/usercentApp/

SERVICEWEAVER_CONFIG=weaver.toml go run .

weaver multi deploy weaver.toml

使用命令

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYTgxNTQ3YjktYTRkOC1kZjNmLWM0NzAtYWZlYWExMWVjNTY0IiwiSUQiOjEsIlVzZXJuYW1lIjoiMTIzIiwiQXV0aG9yaXR5SWQiOiIxIiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODIyMzM5NTUsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MjE0NjU1NX0.OrRw8hfpnedD4hOhGGcwvFPdyC3zfcv41O0qt_qKglQ

curl -H "Content-Type: application/json" -X POST -d '{"username": "admin", "password":"2361bd3bf151f0ec52a3ee762bca3644" }' "http://127.0.0.1:12345/login"

curl -H "Content-Type: application/json" -X POST -d '{"username": "test", "password":"2361bd3bf151f0ec52a3ee762bca3644","nickname":"yyds","authorityId":"888" }' "http://127.0.0.1:12345/user/register"

curl -H "Content-Type: application/json" -X POST -d '{"username": "test", "password":"2361bd3bf151f0ec52a3ee762bca3644","nickname":"yyds","authorityId":"2221" }' "http://127.0.0.1:12345/user/register"

curl -H "Content-Type: application/json" -X POST -d '{"username": "test", "password":"2361bd3bf151f0ec52a3ee762bca3644" }' "http://127.0.0.1:12345/login"

curl -H "Content-Type: application/json" --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOTNiODc0MTItMmUyYy00NmE0LTliYTMtY2M5NjM5MTdlNzZhIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODI5NjA1NDEsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4Mjg3MzE0MX0.EJQLxHO4CHR-xF2Oz6HV6r4BUbwgg_fNEqT5gzjvcZk" -X POST -d '{"nickname":"\u6d4b\u8bd5","Phone":"18112056621" }' "http://127.0.0.1:12345/user/SetSelfInfo"

///setUserAuthority
curl -H "Content-Type: application/json" -X POST -d '{"UserId": 3, "authorityIds":["666","777"] }' "http://127.0.0.1:12345/user/setUserAuthorities"

curl -H "Content-Type: application/json" --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTkyMTRkNjEtYjczZS00OTE0LTg2MmMtNTU4MDJmOGZhYTQzIiwiSUQiOjIsIlVzZXJuYW1lIjoiYWRtaW4iLCJBdXRob3JpdHlJZCI6Ijc3NyIsIkJ1ZmZlclRpbWUiOjYwNDgwMCwiZXhwIjoxNjgzMDMwNjQ5LCJpc3MiOiJuZXctaXQiLCJuYmYiOjE2ODI5NDMyNDl9.vfi2nguknGvqHy2JGpx0A32DL_w3NpLswi74tNXaGKE" -X POST -d '{"UserId": 3, "authorityId":"777" }' "http://127.0.0.1:12345/user/setUserAuthority"

curl -H "Content-Type: application/json" -X POST -d '{"username": "123", "password":"2361bd3bf151f0ec52a3ee762bca3644","newPassword":"345" }' "http://127.0.0.1:12345/user/changePassword"




curl -H "Content-Type: application/json" -X POST -d '{"UserId": 1 }' "http://127.0.0.1:12345/user/resetPassword"

//
curl -H "Content-Type: application/json" -X POST -d '{"page": 1 ,"pageSize":10}' "http://127.0.0.1:12345/user/getUserList"


#没有token
curl -H "Content-Type: application/json" -X GET "http://127.0.0.1:12345/user/getUserInfo"
#有token
curl -H "Content-Type: application/json" --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOTNiODc0MTItMmUyYy00NmE0LTliYTMtY2M5NjM5MTdlNzZhIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODI5NjA1NDEsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4Mjg3MzE0MX0.EJQLxHO4CHR-xF2Oz6HV6r4BUbwgg_fNEqT5gzjvcZk" -X GET "http://127.0.0.1:12345/user/getUserInfo"

#角色
curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "222", "AuthorityName":"\u6d4b\u8bd5","ParentId":"0", "DefaultRouter":"dashboard"}' "http://127.0.0.1:12345/authority/createAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "888", "AuthorityName":"\u666e\u901a\u7528\u6237","ParentId":"0", "DefaultRouter":"dashboard"}' "http://127.0.0.1:12345/authority/createAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "999", "AuthorityName":"\u6d4b\u8bd5\u89d2\u8272","ParentId":"0", "DefaultRouter":"dashboard"}' "http://127.0.0.1:12345/authority/createAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "2221", "AuthorityName":"\u6d4b\u8bd5\u7528\u6237","ParentId":"222", "DefaultRouter":"dashboard"}' "http://127.0.0.1:12345/authority/createAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "2222", "AuthorityName":"\u6d4b\u8bd5\u7528\u6237","ParentId":"222", "DefaultRouter":"dashboard"}' "http://127.0.0.1:12345/authority/createAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"page":1,"pageSize": 20}' "http://127.0.0.1:12345/authority/getAuthorityList"

curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "2221", "AuthorityName":"\u6d4b\u8bd5\u7528\u6237","ParentId":"333", "DefaultRouter":"dashboard"}' "http://127.0.0.1:12345/authority/updateAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"AuthorityId": "2222"}' "http://127.0.0.1:12345/authority/deleteAuthority"

//deleteAuthority
#目录                                                  

curl -H "Content-Type: application/json" -X POST  -d '{"MenuLevel":0,"ParentId": 3, "Path":"dashboard","Name":"dashboard", "Component":"view/dashboard/index.vue","Sort":34,"Hidden":false,"Title":"\u4eea\u8868\u76d8"}' "http://127.0.0.1:12345/menu/addBaseMenu"


curl -H "Content-Type: application/json" -X POST  -d '{"page":1,"pageSize": 20}' "http://127.0.0.1:12345/menu/getMenuList"


curl -H "Content-Type: application/json" -X POST  -d '{"authorityId": "222", "menuId":3}' "http://127.0.0.1:12345/menu/addMenuAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"authorityId": "2221", "menuId":1}' "http://127.0.0.1:12345/menu/addMenuAuthority"


curl -H "Content-Type: application/json" -X POST  -d '{"MenuId":3}' "http://127.0.0.1:12345/menu/getBaseMenuById"

curl -H "Content-Type: application/json" -X POST  -d '{"MenuId":11}' "http://127.0.0.1:12345/menu/deleteBaseMenu"

curl -H "Content-Type: application/json" -X POST  -d '{"authorityId":"222"}' "http://127.0.0.1:12345/menu/getMenusByAuthority"

curl -H "Content-Type: application/json" -X POST  -d '{"authorityId":"222"}' "http://127.0.0.1:12345/menu/getBaseMenuTree"

curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTkyMTRkNjEtYjczZS00OTE0LTg2MmMtNTU4MDJmOGZhYTQzIiwiSUQiOjIsIlVzZXJuYW1lIjoiYWRtaW4iLCJBdXRob3JpdHlJZCI6Ijc3NyIsIkJ1ZmZlclRpbWUiOjYwNDgwMCwiZXhwIjoxNjgzMDMwNjQ5LCJpc3MiOiJuZXctaXQiLCJuYmYiOjE2ODI5NDMyNDl9.vfi2nguknGvqHy2JGpx0A32DL_w3NpLswi74tNXaGKE"  "http://127.0.0.1:12345/menu/getViewMenu"
//getViewMenu

curl -H "Content-Type: application/json" -X POST  -d '{"MenuId":11,"MenuLevel":0,"ParentId": 3, "Path":"dashboardtest","Name":"dashboardtest", "Component":"view/task/index.vue","Sort":35,"Hidden":false,"Title":"\u4e3b\u9875\u9762"}' "http://127.0.0.1:12345/menu/updateBaseMenu"



#转码
研发人员 \u7814\u53d1\u4eba\u5458
普通用户 \u666e\u901a\u7528\u6237
仪表盘 \u4eea\u8868\u76d8

测试角色 \u6d4b\u8bd5\u89d2\u8272
测试 \u6d4b\u8bd5

#mysql config

#数据模型
文件是newit.pdma.json
使用PDManer打开

