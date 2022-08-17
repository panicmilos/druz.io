import { AxiosError } from "axios"
import { FC, useEffect, useState } from "react"
import { createUseStyles } from "react-jss"
import { useMutation } from "react-query"
import { usePostsResult } from "../../hooks"
import { Button, Card, ConfirmationModal, extractErrorMessage, Modal, useNotificationService } from "../../imports"
import { Post, Comment } from "../../models/Post"
import { useCommentsService } from "../../services"
import { CommentsForm } from "./CommentsForm"


type Props = {
  post: Post
}

const useStyles = createUseStyles({
  container: {
    margin: '2% 3% 0% 3%',
    '& button': {
      margin: '0.5em 0.5em 0.5em 0.5em'
    },
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginTop: '20px'
  },
});


export const CommentsList: FC<Props> = ({ post }) => {

  const [isCommentModalOpen, setIsCommentModalOpen] = useState(false);
  const [isDeletePostOpen, setIsDeletePostOpen] = useState(false);
  const [selectedComment, setSelectedComment] = useState<Comment|undefined>();

  const commentsService = useCommentsService(post.id);
  const notificationService = useNotificationService();
  const { result, setResult } = usePostsResult();

  const deleteCommentMutation = useMutation(() => commentsService.delete(selectedComment?.id ?? ''), {
    onSuccess: () => {
      notificationService.success('You have successfully deleted a comment.');
      setResult({ status: 'OK', type: 'DELETE_COMMENT' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'DELETE_COMMENT' });
    }
  });
  const deleteComment = () => deleteCommentMutation.mutate();

  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && ['ADD_COMMENT', 'UPDATE_COMMENT'].includes(result.type)) {
      setIsCommentModalOpen(false);
    }
  }, [result]);

  const classes = useStyles();

  return (
    <div className={classes.container}>

      <Modal title="Write Comment" open={isCommentModalOpen} onClose={() => setIsCommentModalOpen(false)}>
        <CommentsForm postId={post.id} existingComment={selectedComment} isEdit={!!selectedComment} />
      </Modal>

      <div className={classes.buttons}>
        <Button onClick={() => { setSelectedComment(undefined); setIsCommentModalOpen(true)} }>Write</Button>         
      </div>

      <ConfirmationModal title="Delete Comment" open={isDeletePostOpen} onClose={() => setIsDeletePostOpen(false)} onYes={deleteComment}>
        <p>Are you sure you want to delete this comment?</p>
      </ConfirmationModal>

      {
        post.comments?.map((comment: Comment) =>
          <Card>
            <Button onClick={() => { setSelectedComment(comment); setIsCommentModalOpen(true)} }>Update</Button>         
            <Button onClick={() => { setSelectedComment(comment); setIsDeletePostOpen(true)} }>Delete</Button>         

            <p>{comment.text}</p>
            <p>{comment.createdAt}</p>
            <p>{comment.userId}</p>
          </Card>
        )
      }
    </div>
  )
}