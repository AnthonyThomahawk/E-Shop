import axios, {type AxiosInstance} from 'axios'
import {getLocalStorage} from "./LocalStorage";

interface IUserAuth {
    Email: string;
    Token: string;
}


const axiosAPI: AxiosInstance = axios.create({
    baseURL: 'http://localhost:5000'
})

const apiRequest = async (method: string, url: string, request?: {}) => {
    let userData = getLocalStorage("UserData");

    if (userData == undefined) {
        userData = {} as IUserAuth;
        userData.Email = '';
        userData.Token = '';
    }

    if (request == undefined) {
        request = '';
    }

    axiosAPI.defaults.headers.common['Authorization'] = 'Bearer ' + userData.Token;

    try {
        const res = await axiosAPI({
            method,
            url,
            data: request
        });

        return [await Promise.resolve(res.data), await res.headers];
    } catch (err) {
        return await Promise.reject(err);
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