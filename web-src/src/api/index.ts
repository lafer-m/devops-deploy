import axios, {AxiosRequestConfig, Method} from 'axios'

const host = window.location.hostname
const baseURL = 'http://' + host + ':8082'
const axiosInstance = axios.create({
    baseURL: baseURL,
    timeout: 60000
})

// 添加请求拦截器, 例如添加一个 auth token到请求头部
axiosInstance.interceptors.request.use((config: AxiosRequestConfig) => {
    config.headers['Authorization'] = 'test'
    return config
}, err => {
    console.log(err)
    return Promise.reject(err)
})

// 添加响应拦截器
axiosInstance.interceptors.response.use(resp => {
    return resp
}, err => {
    console.log(err)
    return Promise.reject(err)
})

interface FetchConfig {
    path?: string;
    body?: object;
    query?: object;
    headers?: object;
}

/*
 * request 封装axios http请求
 * @Param  method  get/post/put/delete..
 * @Param  axios params
 */
export default function request(method: Method, config: FetchConfig = {path: '', body: {}, query: {}, headers: {}}) {
    const url = baseURL + config.path
    return axiosInstance({
        url,
        method,
        headers: {
            ...config.headers
        },
        data: config.body,
        params: {
            ...config.query
        }
    })
}

