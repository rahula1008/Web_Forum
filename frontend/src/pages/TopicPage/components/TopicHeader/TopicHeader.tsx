import { useState } from "react";
import type { Topic } from "../../../../types/topic"
import "./TopicHeader.css";
import { api } from "../../../../auth/client";

type TopicHeaderProps = {
    topic: Topic | null,
    isAuthed: boolean,
    setIsCreating: React.Dispatch<React.SetStateAction<boolean>>,
    isCreating: boolean;
}

export default function TopicHeader({ topic, isAuthed, setIsCreating, isCreating }: TopicHeaderProps) {

    const [isEditing, setIsEditing] = useState(false);
    const [draftDescription, setDraftDescription] = useState(topic?.description)
    const [isSaving, setIsSaving] = useState(false);
    const [error, setError] = useState<string | null>(null);
    // const [deleteError, setDeleteError] = useState<string | null>(null);

    async function handleSave() {
        if (!topic) {
            setError("Error loading topic");
            return;
        }

        const trimmed = (draftDescription ?? "").trim();
        if (trimmed.length === 0) {
            setError("Description cannot be empty");
            return;
        }

        setError(null);
        setIsSaving(true);

        const prev = topic;
        const newTopic: Topic = { ...topic, description: trimmed };

        try {
            const res = await api.put(`/topics/${topic.id}`, newTopic);
            if (res.status === 400) { 
                setError("Failed to save."); 
                return;
            }
            setIsEditing(false);
            alert("Please refresh the page to see your edits");
        } catch {
            setError("Failed to save. Please try again.");
            setDraftDescription(prev.description);
        } finally {
            setIsSaving(false);
        }
    }

    function handleCancel() {
        setDraftDescription(topic?.description);
        setError(null);
        setIsEditing(false);
    }

    return (
        isEditing ? (
            <>
                <div className="topic-titles">
                    <h1>Topic: {topic?.title}</h1>
                    <p>Description: </p>
                </div>
                <textarea
                    className="topic-edit"
                    value={draftDescription}
                    onChange={(e) => setDraftDescription(e.target.value)}
                    disabled={isSaving}
                    rows={3}
                />
                <div className="topic-actions">
                    <button onClick={handleSave} disabled={isSaving}>
                        {isSaving ? "Saving..." : "Save"}
                    </button>
                    <button onClick={handleCancel} disabled={isSaving}>
                        Cancel
                    </button>
                </div>
                {error && <p className="error-text">{error}</p>}
            </>
        ) : (
            <div className="posts-header">
                <div className="topic-titles">
                    <h1>Topic: {topic?.title}</h1>
                    <h2>Posts</h2>
                    <p>Description: {topic?.description}</p>
                </div>
                <div className="topic-actions">
                    {isAuthed && (
                        <button className="edit-topic-button" type="button" onClick={() => setIsEditing(true)}>
                            Edit
                        </button>
                    )}
                    {isAuthed && (
                        <button
                            className="new-post-button"
                            type="button"
                            onClick={() => setIsCreating((prev) => !prev)}
                        >
                            {isCreating ? "Cancel" : "New Post"}
                        </button>
                    )}
                </div>
            </div>
        )


    )
}
