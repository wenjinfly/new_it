const LOCAL_STORAGE_KEY = "searchHistory";



export const saveHistory = (arr) => {
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(arr));
}

export const loadHistory = () => JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY))

export const removeAllHistory = () => { localStorage.removeItem(LOCAL_STORAGE_KEY) }


export default {
    saveHistory,
    loadHistory,
    removeAllHistory
};