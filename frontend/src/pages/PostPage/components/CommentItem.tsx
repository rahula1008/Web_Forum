import "../PostPage.css";
import type { Comment } from "../../../types/comment";
import type { User } from "../../../types/user";
import { canEditResource } from "../../../utils/canEditResource";
import { useState } from "react";
import { api } from "../../../auth/client";

type CommentItemProps = {
    comment: Comment,
    user: User | null;
};

export default function CommentItem({ comment, user }: CommentItemProps) {

    const canEdit = canEditResource(user, comment.creator_id);

    const [isEditing, setIsEditing] = useState(false);
    const [draftBody, setDraftBody] = useState(comment.body);
    const [isSaving, setIsSaving] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [deleteError, setDeleteError] = useState<string | null>(null);

    async function handleSave() {
        // Cannot have an empty comment:

        const trimmed = draftBody.trim();
        if (trimmed.length === 0) {
            setError("Comment cannot be empty");
            return;
        }

        setError(null);
        setIsSaving(true);

        //Store previous in case need to revert
        const prev = comment;
        const newComment: Comment = { ...comment, body: trimmed };

        try {
            const res = await api.put(`/comments/${comment.id}`, newComment);
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

    function handleCancel() {
        setDraftBody(comment.body);
        setError(null);
        setIsEditing(false);
    }

    async function deleteComment() {
        try {
            const res = await api.delete(`/comments/${comment.id}`);
            if (res.status === 204) {
                alert("Your comment has successfully been deleted. Refresh the page to see changes");
            }
        } catch {
            setDeleteError("Failed to delete comment.")
        }
    }

    return (
        <article className="comment-card">

            {isEditing ? (
                <>
                    <textarea
                        className="comment-edit"
                        value={draftBody}
                        onChange={(e) => setDraftBody(e.target.value)}
                        disabled={isSaving}
                        rows={3}
                    />
                    <div className="comment-actions">
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
                    <p className="comment-body">{comment.body}</p>
                    <div className="comment-meta">
                        <span>By User {comment.creator_id}</span>
                        {canEdit && (
                            <>
                                <button onClick={() => setIsEditing(true)}>Edit</button>
                                <button onClick={deleteComment}>Delete</button>
                            
                                {deleteError && <p>{deleteError}</p>}
                            </>
                        )}
                    </div>
                </>
            )}


        </article>
    );
}
