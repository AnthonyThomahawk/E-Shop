import API from "./AxiosWrapper";

export const getProducts = async (page:number, pageSize:number) => {
    try {
        return await API.get(`/api/products?page=${page}&page_size=${pageSize}`);
    } catch (error) {
        console.error(error)
    }
}

export const getProductCategories = async () => {
    try {
        return await API.get(`/api/categories`);
    } catch (error) {
        console.error(error)
    }
}

export const getProductsByCategoryId = async (page:number, pageSize:number, categoryId:number) => {
    try {
        return await API.get(`/api/products?page=${page}&page_size=${pageSize}&category=${categoryId}`);
    } catch (error) {
        console.error(error)
    }
}

export const getProductDetails = async (id:number) => {
    try {
        return await API.get(`/api/products/${id}`);
    } catch (error) {
        console.error(error);
    }
}