import axios from 'axios';// 引入axios
import { useUsersStore } from '@/pinia/modules/user'

const http = axios.create({
    // baseURL: import.meta.env.VITE_APP_BASE_URL, //在.env中
    baseURL: '/api',
});

// 添加请求拦截器
http.interceptors.request.use(
    function (config) {
        // 在发送请求之前做些什么
        /**
          1、比如添加token之类的请求头信息
          2、添加每次请求loading等
        */
        const userStore = useUsersStore()

        config.headers = {
            'Content-Type': 'application/json',
            'Authorization': userStore.token,
            //'Access-Control-Allow-Origin': true,
            ...config.headers
        }
        return config;
    },
    function (error) {
        // 对请求错误做些什么
        console.log("#####error###")

        console.log(error)
        return Promise.reject(error);
    });

// 添加响应拦截器
http.interceptors.response.use(
    function (response) {
        // 对响应数据做点什么
        /**
          1、集中处理响应数据（如错误码处理）
        */

        return response;
    },
    function (error) {
        // 对响应错误做点什么
        console.log("#####error 2222###")
        console.log(error)

        return Promise.reject(error);
    });

export default http;
