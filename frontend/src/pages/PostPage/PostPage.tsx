import { useEffect, useState } from "react";
import { useParams } from "react-router";
import "./PostPage.css";
import type { Post } from "../../types/post";
import type { Comment } from "../../types/comment";
import CommentItem from "./components/CommentItem";
import { api } from "../../auth/client";

export default function PostPage() {
    const { id, postId } = useParams();
    const [post, setPost] = useState<Post | null>(null);
    const [comments, setComments] = useState<Comment[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    const getPostDetailsEndpoint = `/posts/${postId}`;

    const getCommentsEndpoint = `/posts/${postId}/comments`;

    useEffect(() => {
        const loadPostAndComments = async () => {
            setIsLoading(true);
            try {
                const postResponse = await api.get(getPostDetailsEndpoint);
                setPost(postResponse.data.data);

                const commentsResponse = await api.get(getCommentsEndpoint);
                setComments(commentsResponse.data.data);

            } finally {
                setIsLoading(false);
            }
        };

        loadPostAndComments();
    }, [getPostDetailsEndpoint, getCommentsEndpoint]);

    const createdLabel = post?.created_at
        ? new Date(post.created_at).toLocaleDateString()
        : "";

    return (
        <div className="post-page">
            <header className="post-hero">
                <p className="topic-kicker">Post</p>
                <h1 className="post-title">{post?.title ?? "Untitled Post"}</h1>
                <div className="post-meta">
                    <span>Topic {id ?? "-"}</span>
                    <span>By User {post?.creator_id ?? "-"}</span>
                    <span>{createdLabel}</span>
                </div>
            </header>

            <section className="post-body">
                {isLoading && <p className="post-state">Loading post...</p>}
                {!isLoading && !post && (
                    <p className="post-state">Post not found.</p>
                )}
                {!isLoading && post && <p>{post.body}</p>}
            </section>

            <section className="comments-panel">
                <div className="comments-header">
                    <h2>Comments</h2>
                    <button className="comment-button" type="button">
                        Add Comment
                    </button>
                </div>
                <div className="comments-list">
                    {isLoading && <p className="post-state">Loading comments...</p>}
                    {!isLoading && comments?.length === 0 && (
                        <p className="post-state">No comments yet.</p>
                    )}
                    {!isLoading &&
                        comments?.map((comment) => (
                            <CommentItem key={comment.id} comment={comment} />
                        ))}
                </div>
            </section>
        </div>
    );
}
