
export function setLocalStorage(key, value) {
    try {
        localStorage.setItem(key, JSON.stringify(value));
    }
    catch (error) {
        console.log(error);
    }
}

export function getLocalStorage(key) {
    try {
        const value = localStorage.getItem(key);
        if (value) {
            return JSON.parse(value);
        }
    }
    catch (error) {
        console.log(error);
    }
}