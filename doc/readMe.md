这是一个新项目基于service weaver

weaver generate .
weaver generate ./app/configApp/

继承过weaver里边的东西都需要生成,比如configAPP就需要执行两次
一次是：weaver generate ./app/configApp/model
另外是：weaver generate ./app/configApp/

weaver generate ./app/usercentApp/

SERVICEWEAVER_CONFIG=weaver.toml go run .

weaver multi deploy weaver.toml

使用命令

curl -H "Content-Type: application/json" -X POST -d '{"username": "123", "password":"2361bd3bf151f0ec52a3ee762bca3644" }' "http://127.0.0.1:12345/login"

curl -H "Content-Type: application/json" -X POST -d '{"username": "aaaaaa", "password":"56778899","nickname":"yyds","authorityId":"1" }' "http://127.0.0.1:12345/user/register"


curl -H "Content-Type: application/json" -X POST -d '{"username": "123", "password":"2361bd3bf151f0ec52a3ee762bca3644","newPassword":"345" }' "http://127.0.0.1:12345/user/changePassword"


curl -H "Content-Type: application/json" -X POST -d '{"UserId": 1 }' "http://127.0.0.1:12345/user/resetPassword"

#mysql config

#数据模型
文件是newit.pdma.json
使用PDManer打开

