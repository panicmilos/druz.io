import { AxiosError } from "axios";
import { FC, useContext, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useMutation } from "react-query";
import { usePostsResult, useUsersMap } from "../../hooks";
import { AuthContext, Button, Card, ConfirmationModal, extractErrorMessage, Modal, useNotificationService } from "../../imports";
import { Post } from "../../models/Post";
import { usePostsService } from "../../services";
import { CommentsList } from "./CommentsList";
import { LikesList } from "./LikesList";
import { PostsForm } from "./PostsForm";

type Props = {
  posts: Post[]
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


export const PostsList:FC<Props> = ({ posts }) => {

  const { user } = useContext(AuthContext);
  
  const [isPostModalOpen, setIsPostModalOpen] = useState(false);
  const [isDeletePostOpen, setIsDeletePostOpen] = useState(false);
  const [selectedPost, setSelectedPost] = useState<Post|undefined>();

  const postsService = usePostsService();
  const notificationService = useNotificationService();
  const usersMap = useUsersMap();
  const { result, setResult } = usePostsResult();

  const deletePostMutation = useMutation(() => postsService.delete(selectedPost?.id ?? ''), {
    onSuccess: () => {
      notificationService.success('You have successfully deleted a post.');
      setResult({ status: 'OK', type: 'DELETE_POST' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'DELETE_POST' });
    }
  });
  const deletePost = () => deletePostMutation.mutate();

  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && ['ADD_POST', 'UPDATE_POST'].includes(result.type)) {
      setIsPostModalOpen(false);
    }

    if (result.status === 'OK' && result.type === 'DELETE_POST') {
      setIsDeletePostOpen(false);
    }
  }, [result]);

  const classes = useStyles();

  return (
    <div className={classes.container}>

      <Modal title="Write Post" open={isPostModalOpen} onClose={() => setIsPostModalOpen(false)}>
        <PostsForm existingPost={selectedPost} isEdit={!!selectedPost} />
      </Modal>

      <div className={classes.buttons}>
        <Button onClick={() => { setSelectedPost(undefined); setIsPostModalOpen(true)} }>Write</Button>         
      </div>

      <ConfirmationModal title="Delete Post" open={isDeletePostOpen} onClose={() => setIsDeletePostOpen(false)} onYes={deletePost}>
        <p>Are you sure you want to delete this post?</p>
      </ConfirmationModal>

      {
        posts?.map((post: Post) => 
          <Card>

            {
              user?.ID === post.writtenBy ?
                <>
                  <Button onClick={() => { setSelectedPost(post); setIsPostModalOpen(true)} }>Update</Button>
                  <Button onClick={() => { setSelectedPost(post); setIsDeletePostOpen(true)} }>Delete</Button>
                </> : <></>
            }

            <p>{post.text}</p>
            <p>{post.createdAt}</p>
            <p>{usersMap[post.writtenBy]}</p>

            <LikesList post={post} />
            <CommentsList post={post} />
          </Card>
        )
      }

    </div>
  ); 
  
}