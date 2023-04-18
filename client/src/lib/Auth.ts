import API from "./AxiosWrapper";

export const registerUser = async(email: string, password: string) => {
    try {
        return await API.post(`/api/auth/register`, `{\n"email": "${email}",\n"password": "${password}"\n}`);
    } catch (error) {
        console.log(error);
        throw error;
    }
}

export const loginUser = async(email: string, password: string) => {
    try {
        return await API.post(`/api/auth/login`, `{\n"email": "${email}",\n"password": "${password}"\n}`);
    } catch (error) {
        console.error(error);
        throw error;
    }
}
