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

curl -H "Content-Type: application/json" -X POST -d '{"username": "123", "password":"100" }' "http://127.0.0.1:12345/login"



#mysql config