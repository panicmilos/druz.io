import { AxiosError } from "axios"
import moment from "moment"
import { FC, useContext, useEffect, useState } from "react"
import { createUseStyles } from "react-jss"
import { useMutation } from "react-query"
import { useNavigate } from "react-router-dom"
import { usePostsResult, useUserImagesMap, useUsersMap } from "../../hooks"
import { AuthContext, Button, Card, ConfirmationModal, extractErrorMessage, Modal, useNotificationService } from "../../imports"
import { Post, Comment } from "../../models/Post"
import { useCommentsService } from "../../services"
import { CommentsForm } from "./CommentsForm"


type Props = {
  post: Post
}

const useStyles = createUseStyles({
  container: {
    marginTop: '-8%',
    '& button': {
      margin: '0.5em 0.5em 0.5em 0.5em'
    },
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginTop: '20px'
  },
  nameContainer: {
    display: 'flex',
    alignItems: 'center',
    '& p': {
      marginLeft: '5px',
    },
    '& img': {
      width: '36px',
      height: '36px',
      borderRadius: '50%'
    }
  },
});


export const CommentsList: FC<Props> = ({ post }) => {

  const { user } = useContext(AuthContext);

  const [isCommentModalOpen, setIsCommentModalOpen] = useState(false);
  const [isDeletePostOpen, setIsDeletePostOpen] = useState(false);
  const [selectedComment, setSelectedComment] = useState<Comment|undefined>();

  const commentsService = useCommentsService(post.id);
  const notificationService = useNotificationService();
  const usersMap = useUsersMap();
  const userImagesMap = useUserImagesMap();
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

    if (result.status === 'OK' && result.type === 'DELETE_COMMENT') {
      setIsDeletePostOpen(false);
    }
  }, [result]);

  const nav = useNavigate();
  const navigateToUser = (userId: number) => {
    nav(`/users/${userId}/`)
  }
  
  const classes = useStyles();

  return (
    <div className={classes.container}>

      <Modal title={!!selectedComment ? "Update Comment" : "Write Comment"} open={isCommentModalOpen} onClose={() => setIsCommentModalOpen(false)}>
        <CommentsForm postId={post.id} existingComment={selectedComment} isEdit={!!selectedComment} />
      </Modal>

      <div className={classes.buttons}>
        <Button onClick={() => { setSelectedComment(undefined); setIsCommentModalOpen(true)} }>Comment</Button>         
      </div>

      <ConfirmationModal title="Delete Comment" open={isDeletePostOpen} onClose={() => setIsDeletePostOpen(false)} onYes={deleteComment}>
        <p>Are you sure you want to delete this comment?</p>
      </ConfirmationModal>

      {
        post.comments?.map((comment: Comment) =>
          <Card key={comment.id}>

            <div style={{display: 'flex'}}>
              <div style={{flexGrow: 1, marginTop: '1em'}}>
                <div onClick={() => navigateToUser(+post.writtenBy)} className={classes.nameContainer}>
                  <img src={(user?.ID + '' === comment.userId ? user?.Image : userImagesMap[comment.userId]) || '/images/no-image.png' } />
                  <p>
                    {
                      user?.ID + '' === comment.userId ? 'Your Comment' : usersMap[comment.userId]} @ {moment(comment.createdAt).format('yyyy-MM-DD HH:mm')
                    }
                  </p>
                </div>
              </div>
              <div style={{float: 'right'}}>
                {
                  user?.ID + '' === comment.userId ?
                    <>
                      <Button onClick={() => { setSelectedComment(comment); setIsCommentModalOpen(true)} }>Update</Button>         
                      <Button onClick={() => { setSelectedComment(comment); setIsDeletePostOpen(true)} }>Delete</Button>         
                    </> : <></>
                }
              </div>
            </div>

            <p>{comment.text}</p>
          </Card>
        )
      }
    </div>
  )
}