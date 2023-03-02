import API from "./AxiosWrapper";

export const getProducts = async (page:number, pageSize:number) => {
    try {
        return await API.get(`/api/products?page=${page}&page_size=${pageSize}`);
    } catch (error) {
        console.error(error)
    }
}
