import { useEffect, useState } from "react";
import { useParams } from "react-router";
import "./TopicPage.css";
import PostItem from "./components/PostItem";
import type { Post } from "../../types/post";
import type { Topic } from "../../types/topic";
import axios from "axios";

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;


export default function TopicPage() {
    const { id } = useParams();
    const getAllPostsURL = `${BACKEND_URL}/topics/${id}/posts`;
    const getTopicDetails = `${BACKEND_URL}/topics/${id}`;
    const [posts, setPosts] = useState<Post[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [topic, setTopic] = useState<Topic>();

    useEffect(() => {
        const loadPostsAndTopic = async () => {
            setIsLoading(true);
            try {
                const response = await axios.get(getAllPostsURL)
                setPosts(response.data.data)

                const topicData = await axios.get(getTopicDetails)
                setTopic(topicData.data.data)
            } finally {
                setIsLoading(false);
            }
        };

        loadPostsAndTopic();
    }, [id, getAllPostsURL, getTopicDetails]);

    return (
        <div className="topic-page">
            <header className="topic-header">
                <p className="topic-kicker">Topic</p>
                <h1 className="topic-title">{topic?.title ?? "Untitled"}</h1>
            </header>
            <section className="posts-panel">
                <div className="posts-header">
                    <h2>Posts</h2>
                    <button className="new-post-button" type="button">
                        New Post
                    </button>
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
        </div>
    );
}
