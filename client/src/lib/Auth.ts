import API from "./AxiosWrapper";

interface IUser {
    Email: string;
    Password: string;
}

export const registerUser = async(User: IUser) => {
    try {
        return await API.post(`/api/auth/register`, User);
    } catch (error) {
        console.error(error)
    }
}

export const loginUser = async(User: IUser) => {
    try {
        return await API.post(`/api/auth/login`, User);
    } catch (error) {
        console.error(error)
    }
}
