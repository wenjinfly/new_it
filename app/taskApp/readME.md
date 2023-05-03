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


curl -H "Content-Type: application/json" -X POST -d '{"TaskId": 6 }' "http://127.0.0.1:12345/task/GetTaskByTaskID"