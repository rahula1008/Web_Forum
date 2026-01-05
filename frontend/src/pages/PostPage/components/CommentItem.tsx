import "../PostPage.css";
import type { Comment } from "../../../types/comment";

type CommentItemProps = {
    comment: Comment;
};

export default function CommentItem({ comment }: CommentItemProps) {
    return (
        <article className="comment-card">
            <p className="comment-body">{comment.body}</p>
            <div className="comment-meta">
                <span>By User {comment.creator_id}</span>
                <span>Comment #{comment.id}</span>
            </div>
        </article>
    );
}
