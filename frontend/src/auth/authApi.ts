import type { User } from "../types/user";
import { api } from "./client";


type ApiResponse<T> = {
    success: boolean;
    data: T;
};

const loginEndpoint = "/users/login";
const logoutEndpoint = "/users/logout";
const getMeEndpoint = "/users/me";
const registerEndpoint = "/users/signup";

export async function login(email: string, password: string): Promise<void> {
    await api.post(loginEndpoint, { email, password });
}

export async function logout(): Promise<void> {
    await api.post(logoutEndpoint);
}

export async function getMe(): Promise<User> {
    const res = await api.get<ApiResponse<User>>(getMeEndpoint);
    return res.data.data;
}

export async function register(
    username: string,
    email: string,
    password: string
): Promise<void> {
    await api.post(registerEndpoint, { username, email, password });
}
