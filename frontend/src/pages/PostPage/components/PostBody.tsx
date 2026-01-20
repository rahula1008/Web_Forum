import { useState } from "react";
import type { Post } from "../../../types/post";
import type { User } from "../../../types/user";
import { canEditResource } from "../../../utils/canEditResource";
import './PostBody.css';
import { api } from "../../../auth/client";
import { useNavigate } from "react-router";

type PostBodyProps = {
    isLoading: boolean,
    post: Post,
    user: User | null;
};


export default function PostBody({ isLoading, post, user }: PostBodyProps) {

    const canEdit = canEditResource(user, post.creator_id);
    const navigate = useNavigate();

    const [isEditing, setIsEditing] = useState(false);
    const [draftBody, setDraftBody] = useState(post.body);
    const [isSaving, setIsSaving] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [deleteError, setDeleteError] = useState<string | null>(null);


    const createdLabel = post?.created_at
        ? new Date(post.created_at).toLocaleDateString()
        : "";


    async function handleSave() {
        const trimmed = draftBody.trim();
        if (trimmed.length === 0) {
            setError("Post body cannot be empty");
            return;
        }

        setError(null);
        setIsSaving(true);

        //Store previous in case need to revert
        const prev = post;
        const newPost: Post = { ...post, body: trimmed };

        try {
            const res = await api.put(`/posts/${post.id}`, newPost);
            if (res.status === 400) {
                setError("Failed to save.");
                console.log(res);
                return;
            }
            setIsEditing(false);
            alert("Please refresh the page to see your edits");

        } catch {
            setError("Failed to save. Please try again.");
            setDraftBody(prev.body);
        } finally {
            setIsSaving(false);
        }
    }

    async function handleCancel() {
        setDraftBody(post.body);
        setError(null);
        setIsEditing(false);
    }

    async function deletePost() {
        try {
            const res = await api.delete(`/posts/${post.id}`);
            if (res.status === 204) {
                alert("Your post has successfully been deleted. Taking you to the home page");
            }
            navigate("/");
        } catch {
            setDeleteError("Failed to delete post.")
        } 
    }

    return (
        <>
            {isEditing ? (
                <>
                    <textarea
                        className="post-edit"
                        value={draftBody}
                        onChange={(e) => setDraftBody(e.target.value)}
                        disabled={isSaving}
                        rows={3}
                    />
                    <div className="post-actions">
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
                <>
                    <section className="post-body">
                        {isLoading && <p className="post-state">Loading post...</p>}
                        {!isLoading && !post && (
                            <p className="post-state">Post not found.</p>
                        )}
                        {!isLoading && post && <p className="post-title">{post.title}</p>}
                        {!isLoading && post &&
                            <>
                                <>
                                    <p>{post.body}</p>
                                    <p>Created: {createdLabel}</p>
                                    <p>Author: {post.creator_id}</p>
                                </>
                                <>
                                    {canEdit && (
                                        <>
                                            <button className="edit-button" onClick={() => setIsEditing(true)}>Edit</button>
                                            <button className="delete-button" onClick={deletePost}>Delete</button>

                                            {deleteError && <p>{deleteError}</p>}
                                        </>
                                    )}

                                </>
                            </>
                        }

                    </section>
                </>
            )}
        </>

    )
}
