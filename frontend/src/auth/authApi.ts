import type { User } from "../types/user";
import { api } from "./client";


type ApiResponse<T> = {
    success: boolean;
    data: T;
};

const loginEndpoint = "/users/login";
const logoutEndpoint = "/users/logout";
const getMeEndpoint = "/users/me";

export async function login(username: string, password: string): Promise<User> {
    // Backend sets cookie; response can return user (recommended)
    const res = await api.post<ApiResponse<User>>(loginEndpoint, { username, password });
    return res.data.data;
}

export async function logout(): Promise<void> {
    await api.post(logoutEndpoint);
}

export async function getMe(): Promise<User> {
    const res = await api.get<ApiResponse<User>>(getMeEndpoint);
    return res.data.data;
}
