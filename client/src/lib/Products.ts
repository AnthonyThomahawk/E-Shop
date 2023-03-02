import API from "./AxiosWrapper";

export const getProducts = async () => {
    try {
        return await API.get('/products');
    } catch (error) {
        console.error(error)
    }
}