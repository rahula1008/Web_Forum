import axios from "axios";

export const api = axios.create({
    baseURL: import.meta.env.VITE_BACKEND_URL,
    withCredentials: true, // send the cookie on every request
});

api.interceptors.response.use(
    (res) => res,
    (err) => {
        if (err.response?.status === 401) {
            // In case decide to take unauthorised users somewhere else
        }
        return Promise.reject(err);
    }
);
