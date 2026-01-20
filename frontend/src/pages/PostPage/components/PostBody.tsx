import type { Post } from "../../../types/post";
import type { User } from "../../../types/user";
import { canEditResource } from "../../../utils/canEditResource";
import './PostBody.css';

type PostBodyProps = {
    isLoading: boolean,
    post: Post,
    user: User | null;
};


export default function PostBody({ isLoading, post, user }: PostBodyProps) {


    const canEdit = canEditResource(user, post.creator_id);


    const createdLabel = post?.created_at
        ? new Date(post.created_at).toLocaleDateString()
        : "";

    return (
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
                                <button className="edit-button">Edit</button>
                                <button className="delete-button">Delete</button>
                            </>
                        )}

                    </>
                </>
            }

        </section>
    )
}
