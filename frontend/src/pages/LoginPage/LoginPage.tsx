import { useState } from "react";
import { Link, useNavigate } from "react-router";
import { useAuth } from "../../auth/useAuth";
import "./LoginPage.css";
import Header from "../../components/Header/Header";

export default function LoginPage() {
    const navigate = useNavigate();
    const { login, isAuthed } = useAuth();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [isSubmitting, setIsSubmitting] = useState(false);

    
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (!email || !password) {
            setError("Please enter your email and password.");
            return;
        }

        setError("");
        setIsSubmitting(true);
        try {
            await login(email, password);
            navigate("/");
        } catch (err) {
            const message =
                err instanceof Error ? err.message : "Unable to log in. Try again.";
            setError(message);
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        isAuthed ? <p>Already Logged in</p> :
        <main className="login-page">
            <Header />
            <section className="login-hero">
                <div className="login-hero-content">
                    <p className="login-eyebrow">Welcome back</p>
                    <h1 className="login-title">Log in to Web Forum</h1>
                    <p className="login-subtitle">
                        Join the conversation, share ideas, and keep track of your
                        favorite topics.
                    </p>
                </div>
                <div className="login-card">
                    <form className="login-form" onSubmit={handleSubmit}>
                        <label className="login-field">
                            <span>Email address</span>
                            <input
                                type="email"
                                name="email"
                                autoComplete="email"
                                placeholder="you@example.com"
                                value={email}
                                onChange={(event) => setEmail(event.target.value)}
                            />
                        </label>
                        <label className="login-field">
                            <span>Password</span>
                            <input
                                type="password"
                                name="password"
                                autoComplete="current-password"
                                placeholder="Enter your password"
                                value={password}
                                onChange={(event) => setPassword(event.target.value)}
                            />
                        </label>
                        {error && <p className="login-error">{error}</p>}
                        <button
                            className="login-submit"
                            type="submit"
                            disabled={isSubmitting}
                        >
                            {isSubmitting ? "Logging in..." : "Log in"}
                        </button>
                    </form>
                    <p className="login-footer">
                        New here? <Link to="/register">Create an account</Link>
                    </p>
                </div>
            </section>
        </main>
    );
}
