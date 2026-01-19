import { useEffect, useState } from "react";
import { Link } from "react-router";
import "./HomePage.css";
import type { Topic } from "../../types/topic";
import { useAuth } from "../../auth/useAuth";
import Header from "../../components/Header/Header";
import { api } from "../../auth/client";


const getAllTopicsURL = `/topics`;


export default function HomePage() {
    const { isAuthed } = useAuth();
    const [topics, setTopics] = useState<Topic[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [isCreating, setIsCreating] = useState(false);
    const [topicTitle, setTopicTitle] = useState("");
    const [topicDescription, setTopicDescription] = useState("");
    const [createError, setCreateError] = useState("");
    const [isSubmitting, setIsSubmitting] = useState(false);

    useEffect(() => {
        const loadTopics = async () => {
            setIsLoading(true);
            try {
                const response = await api.get(getAllTopicsURL);

                //console.log("Response: ", response);
                setTopics(response.data.data);

            } finally {
                setIsLoading(false);
            }
        };

        loadTopics();
    }, []);

    const handleCreateTopic = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (!topicTitle.trim() || !topicDescription.trim()) {
            setCreateError("Please fill in both the title and description.");
            return;
        }

        setCreateError("");
        setIsSubmitting(true);
        try {
            const response = await api.post("/topics", {
                title: topicTitle.trim(),
                description: topicDescription.trim(),
            });
            const createdTopic = response.data.data as Topic;
            setTopics((prevTopics) => [createdTopic, ...prevTopics]);
            setTopicTitle("");
            setTopicDescription("");
            setIsCreating(false);
        } catch (error) {
            const message =
                error instanceof Error
                    ? error.message
                    : "Unable to create topic. Try again.";
            setCreateError(message);
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="home-page">
            
            <Header />
            
            <section className="topics-panel">
                <div className="topics-header">
                    <h2 className="topics-title">Topics</h2>
                    {isAuthed && (
                        <button
                            className="create-topic-toggle"
                            type="button"
                            onClick={() => setIsCreating((prev) => !prev)}
                        >
                            {isCreating ? "Cancel" : "Create Topic"}
                        </button>
                    )}
                </div>
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
            {isAuthed && isCreating && (
                <div
                    className="create-topic-backdrop"
                    role="dialog"
                    aria-modal="true"
                >
                    <div className="create-topic-modal">
                        <div className="create-topic-modal-header">
                            <h3>Create a topic</h3>
                            <button
                                className="create-topic-close"
                                type="button"
                                onClick={() => setIsCreating(false)}
                                aria-label="Close create topic form"
                            >
                                ✕
                            </button>
                        </div>
                        <form className="create-topic-form" onSubmit={handleCreateTopic}>
                            <label className="create-topic-field">
                                <span>Title</span>
                                <input
                                    type="text"
                                    name="title"
                                    placeholder="Enter a topic title"
                                    value={topicTitle}
                                    onChange={(event) => setTopicTitle(event.target.value)}
                                />
                            </label>
                            <label className="create-topic-field">
                                <span>Description</span>
                                <textarea
                                    name="description"
                                    placeholder="Describe what this topic is about"
                                    rows={4}
                                    value={topicDescription}
                                    onChange={(event) =>
                                        setTopicDescription(event.target.value)
                                    }
                                />
                            </label>
                            {createError && (
                                <p className="create-topic-error">{createError}</p>
                            )}
                            <div className="create-topic-actions">
                                <button
                                    className="create-topic-cancel"
                                    type="button"
                                    onClick={() => setIsCreating(false)}
                                >
                                    Cancel
                                </button>
                                <button
                                    className="create-topic-submit"
                                    type="submit"
                                    disabled={isSubmitting}
                                >
                                    {isSubmitting ? "Creating..." : "Create Topic"}
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </div>
    );
}
