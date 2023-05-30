import http from '../utils/http';
//

const GetTaskListByName = (params) => {
    return http.post(
        '/task/GetTaskListByName',
        params)
}
const GetTaskByTaskID = (params) => {
    return http.post(
        '/task/GetTaskByTaskID',
        params)
}

export default {
    GetTaskListByName,
    GetTaskByTaskID
}