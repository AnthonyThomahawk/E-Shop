export function setLocalStorage(key, value) {
    localStorage.setItem(key, JSON.stringify(value));
}

export function getLocalStorage(key) {
    const value = localStorage.getItem(key);
    if (value) {
        return JSON.parse(value);
    }
}