import { useEffect, useState } from "react";
import { useParams } from "react-router";
import "./PostPage.css";
import type { Post } from "../../types/post";
import type { Comment } from "../../types/comment";
import CommentItem from "./components/CommentItem";
import { api } from "../../auth/client";
import Header from "../../components/Header/Header";
import { useAuth } from "../../auth/useAuth";

export default function PostPage() {
	const { isAuthed, user } = useAuth();
	const { id, postId } = useParams();
	const [post, setPost] = useState<Post | null>(null);
	const [comments, setComments] = useState<Comment[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [commentBody, setCommentBody] = useState("");
	const [commentError, setCommentError] = useState("");
	const [isSubmitting, setIsSubmitting] = useState(false);

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

	const handleCreateComment = async (
		event: React.FormEvent<HTMLFormElement>
	) => {
		event.preventDefault();
		const postIdValue = Number(postId);
		if (!commentBody.trim()) {
			setCommentError("Please enter a comment.");
			return;
		}
		if (!Number.isFinite(postIdValue)) {
			setCommentError("Post is missing. Please refresh and try again.");
			return;
		}
		if (!user?.id) {
			setCommentError("You must be logged in to comment.");
			return;
		}

		setCommentError("");
		setIsSubmitting(true);
		try {
			const response = await api.post("/comments", {
				body: commentBody.trim(),
				post_id: postIdValue,
				creator_id: Number(user.id),
			});
			const createdComment = response.data.data as Comment;
			setComments((prevComments) => [createdComment, ...prevComments]);
			setCommentBody("");
		} catch (error) {
			const message =
				error instanceof Error
					? error.message
					: "Unable to create comment. Try again.";
			setCommentError(message);
		} finally {
			setIsSubmitting(false);
		}
	};

	const typeOfPage = "Post";


	return (
		<div className="post-page">
			<Header typeOfPage={typeOfPage} centerText=" " />


			<section className="post-body">
				{isLoading && <p className="post-state">Loading post...</p>}
				{!isLoading && !post && (
					<p className="post-state">Post not found.</p>
				)}
				{!isLoading && post && <p className="post-title">{post.title}</p>}
				{!isLoading && post &&
					<>
						<p>{post.body}</p>
						<p>Created: {createdLabel}</p>
						<p>Author: {id}</p>
					</>
				}

			</section>

			<section className="comments-panel">
				<div className="comments-header">
					<h2>Comments</h2>
				</div>

				<div className="comments-list">
					{isLoading && <p className="post-state">Loading comments...</p>}
					{!isLoading && comments?.length === 0 && (
						<p className="post-state">No comments yet.</p>
					)}
					{!isLoading &&
						comments?.map((comment) => (
							<CommentItem key={comment.id} user={user} comment={comment} />
						))}
				</div>


			</section>
			{isAuthed && (
				<form className="comment-form" onSubmit={handleCreateComment}>
					<label className="comment-field">
						<span>Add a comment</span>
						<textarea
							name="comment"
							placeholder="Share your thoughts"
							rows={3}
							value={commentBody}
							onChange={(event) =>
								setCommentBody(event.target.value)
							}
						/>
					</label>
					{commentError && (
						<p className="comment-error">{commentError}</p>
					)}
					<button
						className="comment-submit"
						type="submit"
						disabled={isSubmitting}
					>
						{isSubmitting ? "Posting..." : "Post comment"}
					</button>
				</form>
			)}
			{!isAuthed && (
				<>
					<p className="login-request">Please log in to post a comment / reply to this post</p>
				</>
			)}
		</div>
	);
}
