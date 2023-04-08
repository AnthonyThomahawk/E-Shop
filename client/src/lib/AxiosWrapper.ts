import axios, {type AxiosInstance} from 'axios'
import {getLocalStorage} from "./LocalStorage";

const baseURL = 'http://localhost:5000'

interface IUserAuth {
    Email: string;
    Token: string;
}

async function makeGetRequest(url: string, headers: Record<string, string>): Promise<any> {
    return new Promise((resolve, reject) => {
        const xhr = new XMLHttpRequest();
        xhr.open('GET', url);
        Object.entries(headers).forEach(([key, value]) => {
            xhr.setRequestHeader(key, value);
        });
        xhr.onreadystatechange = function() {
            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    try {
                        const responseJson = JSON.parse(xhr.responseText);
                        resolve(responseJson);
                    } catch (error) {
                        reject(new Error(`Error parsing JSON response: ${error.message}`));
                    }
                } else {
                    reject(new Error(`Error: ${xhr.statusText}`));
                }
            }
        };
        xhr.send();
    });
}

const apiRequest = async (inMethod: string, inUrl: string, request?: {}) => {
    let userData = getLocalStorage("UserData");

    if (userData == undefined) {
        userData = {} as IUserAuth;
        userData.Email = '';
        userData.Token = '';
    }

    if (request == undefined) {
        request = '';
    }

    const tok = userData.Token;

    try {
        const res = await axios({
            baseURL: 'http://localhost:5000',
            method: `${inMethod}`,
            url: `${inUrl}`,
            data: request,
            headers: {
                credentials: "include",
                'Authorization': 'Bearer ' + tok,
            }

        })

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