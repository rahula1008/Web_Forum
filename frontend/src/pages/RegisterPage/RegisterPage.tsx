import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router";
import { register } from "../../auth/authApi";
import "./RegisterPage.css";

export default function RegisterPage() {
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [isSubmitting, setIsSubmitting] = useState(false);

    useEffect(() => {
        document.title = "Register";
    }, []);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (!username || !email || !password) {
            setError("Please fill in all fields.");
            return;
        }

        setError("");
        setIsSubmitting(true);
        try {
            await register(username, email, password);
            navigate("/login");
        } catch (err) {
            const message =
                err instanceof Error ? err.message : "Unable to register. Try again.";
            setError(message);
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <main className="register-page">
            <section className="register-hero">
                <div className="register-hero-content">
                    <p className="register-eyebrow">Create an account</p>
                    <h1 className="register-title">Join the Web Forum</h1>
                    <p className="register-subtitle">
                        Set up your profile and jump into the latest discussions.
                    </p>
                </div>
                <div className="register-card">
                    <form className="register-form" onSubmit={handleSubmit}>
                        <label className="register-field">
                            <span>Username</span>
                            <input
                                type="text"
                                name="username"
                                autoComplete="username"
                                placeholder="forumfan"
                                value={username}
                                onChange={(event) => setUsername(event.target.value)}
                            />
                        </label>
                        <label className="register-field">
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
                        <label className="register-field">
                            <span>Password</span>
                            <input
                                type="password"
                                name="password"
                                autoComplete="new-password"
                                placeholder="Create a password"
                                value={password}
                                onChange={(event) => setPassword(event.target.value)}
                            />
                        </label>
                        {error && <p className="register-error">{error}</p>}
                        <button
                            className="register-submit"
                            type="submit"
                            disabled={isSubmitting}
                        >
                            {isSubmitting ? "Creating account..." : "Sign up"}
                        </button>
                    </form>
                    <p className="register-footer">
                        Already have an account? <Link to="/login">Log in</Link>
                    </p>
                </div>
            </section>
        </main>
    );
}
