import axios, { type AxiosInstance } from 'axios'

const axiosAPI: AxiosInstance = axios.create({
    baseURL: 'http://localhost:5000'
})

const apiRequest = async (method: string, url: string, request?: {}) => {
    const headers = {
        authorization: ''
    }

    try {
        const res = await axiosAPI({
            method,
            url,
            data: request,
            headers
        })

        return await Promise.resolve(res.data)
    } catch (err) {
        return await Promise.reject(err)
    }
}

const get = (url: string) => apiRequest('get', url)
const post = (url: string, request: {}) => apiRequest('post', url, request)
const put = (url: string, request: {}) => apiRequest('put', url, request)
const del = (url: string, request: {}) => apiRequest('delete', url, request)
const patch = (url: string, request: {}) => apiRequest('patch', url, request)

const API = {
    get,
    delete: del,
    post,
    put,
    patch
}
export default API