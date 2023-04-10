这是一个新项目基于service weaver

weaver generate .
weaver generate ./app/configApp/


SERVICEWEAVER_CONFIG=weaver.toml go run .

weaver multi deploy weaver.toml

使用命令
curl "localhost:12345/greeter?name=66188"


#mysql config