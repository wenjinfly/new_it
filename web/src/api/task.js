import http from '../utils/http';
//

const GetTaskListByName = (params) => {
    return http.post(
        '/task/GetTaskListByName',
        params)
}


export default {
    GetTaskListByName,
}