import { AxiosError } from "axios"
import { FC, useContext } from "react"
import { useMutation } from "react-query"
import { usePostsResult, useUsersMap } from "../../hooks"
import { AuthContext, Button, extractErrorMessage, useNotificationService } from "../../imports"
import { Post } from "../../models/Post"
import { useLikesService } from "../../services"

type Props = {
  post: Post
}


export const LikesList: FC<Props> = ({ post }) => {

  const { user } = useContext(AuthContext);
  
  const likesService = useLikesService(post.id);
  const notificationService = useNotificationService();
  const usersMap = useUsersMap();
  const { setResult } = usePostsResult();

  const likeMutation = useMutation(() => likesService.like(), {
    onSuccess: () => {
      setResult({ status: 'OK', type: 'LIKE' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'LIKE' });
    }
  });
  const like = () => likeMutation.mutate();

  const dislikeMutation = useMutation(() => likesService.dislike(), {
    onSuccess: () => {
      setResult({ status: 'OK', type: 'DISLIKE' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'DISLIKE' });
    }
  });
  const dislike = () => dislikeMutation.mutate();

  
  return (
    <>

     {
      post?.likedBy.map(l => usersMap[l]).join(', ')
     }

     {
      !post.likedBy.includes(user?.ID ?? '') ?
        <Button onClick={like} type="submit">Like</Button> :
        <Button onClick={dislike} type="submit">Dislike</Button>
     }

    </>
  )
}