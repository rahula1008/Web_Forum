import { useEffect, useState } from "react";
import { useParams } from "react-router";
import "./TopicPage.css";
import PostItem from "./components/PostItem";
import type { Post } from "../../types/post";
import type { Topic } from "../../types/topic";
import { api } from "../../auth/client";
import Header from "../../components/Header/Header";
import { useAuth } from "../../auth/useAuth";



export default function TopicPage() {
    const { isAuthed, user } = useAuth();
    const { id } = useParams();
    const getAllPostsURL = `/topics/${id}/posts`;
    const getTopicDetails = `/topics/${id}`;
    const [posts, setPosts] = useState<Post[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [topic, setTopic] = useState<Topic>();
    const [isCreating, setIsCreating] = useState(false);
    const [postTitle, setPostTitle] = useState("");
    const [postBody, setPostBody] = useState("");
    const [createError, setCreateError] = useState("");
    const [isSubmitting, setIsSubmitting] = useState(false);

    useEffect(() => {
        const loadPostsAndTopic = async () => {
            setIsLoading(true);
            try {
                const response = await api.get(getAllPostsURL)
                setPosts(response.data.data)

                const topicData = await api.get(getTopicDetails)
                setTopic(topicData.data.data)
            } finally {
                setIsLoading(false);
            }
        };

        loadPostsAndTopic();
    }, [id, getAllPostsURL, getTopicDetails]);

    const handleCreatePost = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const topicIdValue = Number(id);
        if (!postTitle.trim() || !postBody.trim()) {
            setCreateError("Please fill in both the title and body.");
            return;
        }
        if (!Number.isFinite(topicIdValue)) {
            setCreateError("Topic is missing. Please refresh and try again.");
            return;
        }
        if (!user?.id) {
            setCreateError("You must be logged in to create a post.");
            return;
        }

        setCreateError("");
        setIsSubmitting(true);
        try {
            const response = await api.post("/posts", {
                title: postTitle.trim(),
                body: postBody.trim(),
                topic_id: topicIdValue,
                creator_id: Number(user.id),
            });
            const createdPost = response.data.data as Post;
            setPosts((prevPosts) => [createdPost, ...prevPosts]);
            setPostTitle("");
            setPostBody("");
            setIsCreating(false);
        } catch (error) {
            const message =
                error instanceof Error
                    ? error.message
                    : "Unable to create post. Try again.";
            setCreateError(message);
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="topic-page">
            <Header typeOfPage="Topic" />
            
            <section className="posts-panel">
                <div className="posts-header">
                    <h1>Topic: {topic?.title}</h1>
                    <h2>Posts</h2>
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
                <div className="posts-list">
                    {isLoading && <p className="posts-state">Loading posts...</p>}
                    {!isLoading && posts.length === 0 && (
                        <p className="posts-state">No posts yet. Be the first to post.</p>
                    )}
                    {!isLoading &&
                        posts.map((post) => (
                            <PostItem key={post.id} post={post} />
                        ))}
                </div>
            </section>
            {isAuthed && isCreating && (
                <div className="create-post-backdrop" role="dialog" aria-modal="true">
                    <div className="create-post-modal">
                        <div className="create-post-modal-header">
                            <h3>Create a post</h3>
                            <button
                                className="create-post-close"
                                type="button"
                                onClick={() => setIsCreating(false)}
                                aria-label="Close create post form"
                            >
                                ✕
                            </button>
                        </div>
                        <form className="create-post-form" onSubmit={handleCreatePost}>
                            <label className="create-post-field">
                                <span>Title</span>
                                <input
                                    type="text"
                                    name="title"
                                    placeholder="Enter a post title"
                                    value={postTitle}
                                    onChange={(event) => setPostTitle(event.target.value)}
                                />
                            </label>
                            <label className="create-post-field">
                                <span>Body</span>
                                <textarea
                                    name="body"
                                    placeholder="Write your post"
                                    rows={5}
                                    value={postBody}
                                    onChange={(event) => setPostBody(event.target.value)}
                                />
                            </label>
                            {createError && (
                                <p className="create-post-error">{createError}</p>
                            )}
                            <div className="create-post-actions">
                                <button
                                    className="create-post-cancel"
                                    type="button"
                                    onClick={() => setIsCreating(false)}
                                >
                                    Cancel
                                </button>
                                <button
                                    className="create-post-submit"
                                    type="submit"
                                    disabled={isSubmitting}
                                >
                                    {isSubmitting ? "Creating..." : "Create Post"}
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </div>
    );
}
