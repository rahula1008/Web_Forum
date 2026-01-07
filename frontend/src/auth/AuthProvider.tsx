import React, { useEffect, useMemo, useState } from "react";
import type { User } from "../types/user";
import { getMe, login as apiLogin, logout as apiLogout } from "./authApi";
import { AuthContext } from "./useAuth";


export function AuthProvider({ children }: { children: React.ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    //sets the user using getme, which requires authentication 
    const refreshMe = async () => {
        try {
            const me = await getMe();
            setUser(me);
        } catch {
            setUser(null);
        }
    };

    //
    useEffect(() => {
        const loadMe = async () => {
            setIsLoading(true);
            await refreshMe();
            setIsLoading(false);
        };
        loadMe();
    }, []);

    const login = async (username: string, password: string) => {
        const me = await apiLogin(username, password); // cookie set by backend
        setUser(me);
    };

    const logout = async () => {
        await apiLogout(); // backend clears cookie
        setUser(null);
    };

    const value = useMemo(
        () => ({
            user,
            isLoading,
            isAuthed: !!user,
            refreshMe,
            login,
            logout,
        }),
        [user, isLoading]
    );

    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}


