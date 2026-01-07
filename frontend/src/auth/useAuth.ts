import { createContext, useContext } from "react";
import type { User } from "../types/user";


type AuthContextValue = {
    user: User | null;
    isLoading: boolean;
    isAuthed: boolean;
    refreshMe: () => Promise<void>;
    login: (username: string, password: string) => Promise<void>;
    logout: () => Promise<void>;
};

export const AuthContext = createContext<AuthContextValue | null>(null);

export function useAuth() {
    const ctx = useContext(AuthContext);
    if (!ctx) throw new Error("useAuth must be used within AuthProvider");
    return ctx;
}
