import { Link, useParams } from "react-router";
import type { Post } from "../../../types/post";
import "../TopicPage.css";


type PostItemProps = {
    post: Post;
};

export default function PostItem({ post }: PostItemProps) {
    const { id } = useParams();

    return (
        <article className="post-card">
            <div className="post-card-main">
                <h3 className="post-card-title">
                    <Link className="post-card-link" to={`/topics/${id}/posts/${post.id}`}>
                        {post.title}
                    </Link>
                </h3>
                <p className="post-card-excerpt">{post.body}</p>
            </div>
            <div className="post-card-meta">
                <span>By {post.creator_id}</span>
                <span>{post.created_at.toString()}</span>
            </div>
        </article>
    );
}
