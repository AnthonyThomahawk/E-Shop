import API from "./AxiosWrapper";

export const getCart = async (page: number, pageSize: number) => {
    try {
        return await API.get(`/api/cart?page=${page}&page_size=${pageSize}`);
    } catch (error) {
        console.error(error);
    }
}

export const updateItem = async(productID: number, quantity: number) => {
    try {
        return await API.put(`/api/cart/${productID}`, `{\n"Quantity":${quantity}\n}`);
    } catch (error) {
        console.error(error);
    }
}

export const deleteItem = async(productID: number) => {
    try {
        return await API.delete(`/api/cart/${productID}`, '');
    } catch (error) {
        console.error(error);
    }
}

export const insertItem = async(productID: number, quantity: number) => {
    try {
        return await API.post(`/api/cart`, `{\n` +
            `    "ProductId": ${productID},\n` +
            `    "Quantity": ${quantity}\n` +
            `}`);
    } catch (error) {
        console.error(error);
    }
}