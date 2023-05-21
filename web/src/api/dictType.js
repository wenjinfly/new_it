import http from '../utils/http';

const getDictTypeList = (params) => {

    return http.post(
        '/dict/GetDictTypeList',
        params
    )
}

const AddDictType = (params) => {
    return http.post(
        '/dict/AddDictType',
        params
    )
}

const GetDictTypeByCode = (params) => {
    return http.post(
        '/dict/GetDictTypeByCode',
        params)
}

const UpdateDictType = (params) => {
    return http.post(
        '/dict/UpdateDictType',
        params)
}

const DeleteDictType = (params) => {
    return http.post(
        '/dict/DeleteDictType',
        params)
}

export default {
    getDictTypeList,
    AddDictType,
    GetDictTypeByCode,
    UpdateDictType,
    DeleteDictType

};