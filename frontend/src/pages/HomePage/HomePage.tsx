import { useEffect, useState } from "react";
import { Link } from "react-router";
import "./HomePage.css";
import axios from "axios";
import type { Topic } from "../../types/topic";
import { useAuth } from "../../auth/useAuth";



const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;
const getAllTopicsURL = `${BACKEND_URL}/topics`;


export default function HomePage() {
    const {isAuthed, user} = useAuth();
    const [topics, setTopics] = useState<Topic[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const loadTopics = async () => {
            setIsLoading(true);
            try {
                const response = await axios.get(getAllTopicsURL);

                //console.log("Response: ", response);
                setTopics(response.data.data);

            } finally {
                setIsLoading(false);
            }
        };

        loadTopics();
    }, []);

    return (
        <div className="home-page">
            <header className="home-header">
                <h1>WEB FORUM</h1>
            </header>
            <div>
                <p>Test user: {isAuthed ?? user?.email}</p>
            </div>
            <section className="topics-panel">
                <h2 className="topics-title">Topics</h2>
                <div className="topics-list">
                    {isLoading && <p className="topics-state">Loading topics...</p>}
                    {!isLoading && topics.length === 0 && (
                        <p className="topics-state">No topics yet.</p>
                    )}
                    {!isLoading &&
                        topics.map((topic) => (
                            <div key={topic.id} className="topic-item">
                                <Link className="topic-link" to={`/topics/${topic.id}/posts`}>
                                    {topic.title}
                                </Link>
                            </div>
                        ))}
                </div>
            </section>
        </div>
    );
}
