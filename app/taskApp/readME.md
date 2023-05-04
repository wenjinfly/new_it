#task
##登录获取token
curl -H "Content-Type: application/json" -X POST -d '{"username": "test", "password":"2361bd3bf151f0ec52a3ee762bca3644" }' "http://127.0.0.1:12345/login"
##添加任务 

小程序开发->\u5c0f\u7a0b\u5e8f\u5f00\u53d1
我们需要做一个小程序->\u6211\u4eec\u9700\u8981\u505a\u4e00\u4e2a\u5c0f\u7a0b\u5e8f
可以保证按时开发完成小程序->\u53ef\u4ee5\u4fdd\u8bc1\u6309\u65f6\u5f00\u53d1\u5b8c\u6210\u5c0f\u7a0b\u5e8f
远程->\u8fdc\u7a0b
有开发小程序经验->\u6709\u5f00\u53d1\u5c0f\u7a0b\u5e8f\u7ecf\u9a8c

"TaskName":"小程序开发","TaskRewardMin":5000.21,"TaskRewardMax":6000.66,
"TaskStatus":"202","TaskRequiredPerson":3,"TaskDescribe":"我们需要做一个小程序",
"TaskPhase":"112","TaskDuty":"可以保证按时开发完成小程序","TaskGetQualification":"有开发小程序经验",
"TaskAddress":"远程","TaskBegin":"2023-05-01 10:00:05","TaskEnd":"2023-05-010 10:00:05"                             

"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug"

curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug" -d '{"TaskName":"\u5c0f\u7a0b\u5e8f\u5f00\u53d1","TaskRewardMin":5000.21,"TaskRewardMax":6000.66,"TaskStatus":"202","TaskRequiredPerson":3,"TaskDescribe":"\u6211\u4eec\u9700\u8981\u505a\u4e00\u4e2a\u5c0f\u7a0b\u5e8f","TaskPhase":"112","TaskDuty":"\u53ef\u4ee5\u4fdd\u8bc1\u6309\u65f6\u5f00\u53d1\u5b8c\u6210\u5c0f\u7a0b\u5e8f","TaskGetQualification":"\u6709\u5f00\u53d1\u5c0f\u7a0b\u5e8f\u7ecf\u9a8c","TaskAddress":"\u8fdc\u7a0b","TaskBegin":"2023-05-01 10:00:05","TaskEnd":"2023-05-10 10:00:05"   }' "http://127.0.0.1:12345/task/AddTask"

curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug" -d '{"TaskName":"\u5c0f\u7a0b\u5e8f\u5f00\u53d13","TaskRewardMin":5000.21,"TaskRewardMax":6000.66,"TaskStatus":"202","TaskRequiredPerson":3,"TaskDescribe":"\u6211\u4eec\u9700\u8981\u505a\u4e00\u4e2a\u5c0f\u7a0b\u5e8f","TaskPhase":"112","TaskDuty":"\u53ef\u4ee5\u4fdd\u8bc1\u6309\u65f6\u5f00\u53d1\u5b8c\u6210\u5c0f\u7a0b\u5e8f","TaskGetQualification":"\u6709\u5f00\u53d1\u5c0f\u7a0b\u5e8f\u7ecf\u9a8c","TaskAddress":"\u8fdc\u7a0b","TaskBegin":"2023-05-01 10:00:05","TaskEnd":"2023-05-10 10:00:05"   }' "http://127.0.0.1:12345/task/AddTask"


curl -H "Content-Type: application/json" -X POST -d '{"TaskId": 2 }' "http://127.0.0.1:12345/task/GetTaskByTaskID"

curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug" -d '{"page":1,"pageSize": 20 }' "http://127.0.0.1:12345/task/GetTaskListByUserId"


curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug" -d '{"TaskId": 4,"TaskPhase":"111" }' "http://127.0.0.1:12345/task/UpdateTaskPhase"

curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug" -d '{"TaskId": 4,"TaskStatus":"203" }' "http://127.0.0.1:12345/task/UpdateTaskStatus"

//DeleteTask

curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMxMjkyNjcsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzA0MTg2N30.8OjCHqgCr8NDv32FS3maJWsxDlAIOptHbHeY6UiL4ug" -d '{"TaskId": 6 }' "http://127.0.0.1:12345/task/DeleteTask"


#-----------------------------
##chat
curl -H "Content-Type: application/json" -X POST -d '{"username": "test", "password":"2361bd3bf151f0ec52a3ee762bca3644" }' "http://127.0.0.1:12345/login"

//test
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMyOTcyMDMsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzIwOTgwM30.3i1G0N2HRYUXbpJT6GK3tK2OpAKDEQdrN5gHa6ZfUdM

//admin
curl -H "Content-Type: application/json" -X POST -d '{"username": "admin", "password":"2361bd3bf151f0ec52a3ee762bca3644" }' "http://127.0.0.1:12345/login"

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTkyMTRkNjEtYjczZS00OTE0LTg2MmMtNTU4MDJmOGZhYTQzIiwiSUQiOjIsIlVzZXJuYW1lIjoiYWRtaW4iLCJBdXRob3JpdHlJZCI6Ijc3NyIsIkJ1ZmZlclRpbWUiOjYwNDgwMCwiZXhwIjoxNjgzMjk3ODEyLCJpc3MiOiJuZXctaXQiLCJuYmYiOjE2ODMyMTA0MTJ9.3fkXMcqkrxghLU--0NAEFzBJukNJ_K0JTlq__Sz5Z4U


curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMyOTcyMDMsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzIwOTgwM30.3i1G0N2HRYUXbpJT6GK3tK2OpAKDEQdrN5gHa6ZfUdM" -d '{"FromId": 3, "FromName":"\u6d4b\u8bd5\u4eba\u5458","ToId": 2, "ToName":"\u8d85\u7ea7\u7ba1\u7406\u5458","SendTime":"2023-05-10 10:00:05","Content":"\u60a8\u597d\uff0c\u5f53\u524d\u5c0f\u7a0b\u5e8f\u4efb\u52a1\u8fd8\u53ef\u4ee5\u63a5\u5355\u4e48\uff1f","ChatType":1,"ClassType":1 }' "http://127.0.0.1:12345/chat/AddChat"


curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTg4Y2Y4YjYtZDA4Ni00NmUxLWIwNTQtMjU3MTQxYWY2ZTlkIiwiSUQiOjMsIlVzZXJuYW1lIjoidGVzdCIsIkF1dGhvcml0eUlkIjoiNjY2IiwiQnVmZmVyVGltZSI6NjA0ODAwLCJleHAiOjE2ODMyOTcyMDMsImlzcyI6Im5ldy1pdCIsIm5iZiI6MTY4MzIwOTgwM30.3i1G0N2HRYUXbpJT6GK3tK2OpAKDEQdrN5gHa6ZfUdM" -d '{"FromId": 3, "FromName":"\u6d4b\u8bd5\u4eba\u5458","SendTime":"2023-05-10 09:00:05","Content":"\u5c0f\u7a0b\u5e8f\u9879\u76ee\u7ecf\u9a8c\u4e30\u5bcc\uff0c\u5df2\u7ecf\u6309\u65f6\u4fdd\u91cf\u5b8c\u6210100\u4e2a\u4e86","ChatType":1,"ClassType":2,"TaskID":1,"TaskName":"\u5c0f\u7a0b\u5e8f\u5f00\u53d1" }' "http://127.0.0.1:12345/chat/AddChat"



curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTkyMTRkNjEtYjczZS00OTE0LTg2MmMtNTU4MDJmOGZhYTQzIiwiSUQiOjIsIlVzZXJuYW1lIjoiYWRtaW4iLCJBdXRob3JpdHlJZCI6Ijc3NyIsIkJ1ZmZlclRpbWUiOjYwNDgwMCwiZXhwIjoxNjgzMjk3ODEyLCJpc3MiOiJuZXctaXQiLCJuYmYiOjE2ODMyMTA0MTJ9.3fkXMcqkrxghLU--0NAEFzBJukNJ_K0JTlq__Sz5Z4U" -d '{"FromId": 2, "FromName":"\u8d85\u7ea7\u7ba1\u7406\u5458","ToId": 3, "ToName":"\u6d4b\u8bd5\u4eba\u5458","SendTime":"2023-05-10 10:00:09","Content":"\u73b0\u5728\u8fd8\u53ef\u4ee5\u63a5\u5355\uff0c\u4f60\u4e4b\u524d\u505a\u8fc7\u7c7b\u4f3c\u7684\u9879\u76ee\u4e48\uff1f","ChatType":1,"ClassType":1 }' "http://127.0.0.1:12345/chat/AddChat"


curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTkyMTRkNjEtYjczZS00OTE0LTg2MmMtNTU4MDJmOGZhYTQzIiwiSUQiOjIsIlVzZXJuYW1lIjoiYWRtaW4iLCJBdXRob3JpdHlJZCI6Ijc3NyIsIkJ1ZmZlclRpbWUiOjYwNDgwMCwiZXhwIjoxNjgzMjk3ODEyLCJpc3MiOiJuZXctaXQiLCJuYmYiOjE2ODMyMTA0MTJ9.3fkXMcqkrxghLU--0NAEFzBJukNJ_K0JTlq__Sz5Z4U" -d '{"FromId": 2, "ToId": 3, "ClassType":1 }' "http://127.0.0.1:12345/chat/GetChatListByIds"


curl -H "Content-Type: application/json" -X POST --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTkyMTRkNjEtYjczZS00OTE0LTg2MmMtNTU4MDJmOGZhYTQzIiwiSUQiOjIsIlVzZXJuYW1lIjoiYWRtaW4iLCJBdXRob3JpdHlJZCI6Ijc3NyIsIkJ1ZmZlclRpbWUiOjYwNDgwMCwiZXhwIjoxNjgzMjk3ODEyLCJpc3MiOiJuZXctaXQiLCJuYmYiOjE2ODMyMTA0MTJ9.3fkXMcqkrxghLU--0NAEFzBJukNJ_K0JTlq__Sz5Z4U" -d '{"FromId": 3, "ToId": 2, "ClassType":1 }' "http://127.0.0.1:12345/chat/GetChatListByIds"

