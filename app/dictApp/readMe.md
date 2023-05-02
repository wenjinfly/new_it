#用于测试

性别->\u6027\u522b
男->\u7537
女->\u5973

curl -H "Content-Type: application/json" -X POST -d '{"TypeCode": "400", "TypeName":"Gender","TypeCNName":"\u6027\u522b" }' "http://127.0.0.1:12345/dict/AddDictType"


curl -H "Content-Type: application/json" -X POST -d '{"TypeCode": "400", "TypeCNName":"\u6027\u522b\u7537" }' "http://127.0.0.1:12345/dict/UpdateDictType"

curl -H "Content-Type: application/json" -X POST -d '{"TypeCode": "400" }' "http://127.0.0.1:12345/dict/GetDictTypeByCode"


curl -H "Content-Type: application/json" -X POST -d '{"page":1,"pageSize": 20}' "http://127.0.0.1:12345/dict/GetDictTypeList"

curl -H "Content-Type: application/json" -X POST -d '{"TypeCode": "400" }' "http://127.0.0.1:12345/dict/DeleteDictType"



#--------------

curl -H "Content-Type: application/json" -X POST -d '{"Code": "401", "Name":"female","CNName":"\u5973","Fixed":true,"TypeCode":"400" }' "http://127.0.0.1:12345/dict/AddDictInfo"


curl -H "Content-Type: application/json" -X POST -d '{"Code": "400" }' "http://127.0.0.1:12345/dict/DeleteDictInfo"

curl -H "Content-Type: application/json" -X POST -d '{"Code": "401", "Name":"female","CNName":"\u5973","Fixed":true,"TypeCode":"400" }' "http://127.0.0.1:12345/dict/UpdateDictInfo"

curl -H "Content-Type: application/json" -X POST -d '{"Code": "400" }' "http://127.0.0.1:12345/dict/GetDictInfoByCode"


curl -H "Content-Type: application/json" -X POST -d '{"page":1,"pageSize": 20, "TypeCode": "300"}' "http://127.0.0.1:12345/dict/GetDictInfoListByTypeCode"