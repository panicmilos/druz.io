import { Button, extractErrorMessage, Form, FormTextAreaInput, useNotificationService } from "../../imports";
import { usePostsService } from "../../services"

import * as Yup from 'yup';
import { useMutation } from "react-query";
import { usePostsResult } from "../../hooks";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";
import { Post } from "../../models/Post";
import { FC } from "react";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '1em',
  }
});

type Props = {
  existingPost?: Post,
  isEdit: boolean
}

export const PostsForm:FC<Props> = ({ existingPost = undefined, isEdit = false}) => {

  const postsService = usePostsService();
  const notificationService = useNotificationService();

  const schema = Yup.object().shape({
    text: Yup.string()
      .required(() => ({ text: "Text must be provided." }))
  });

  const { setResult } = usePostsResult();

  const addPostMutation = useMutation((createPost: any) => postsService.add(createPost), {
    onSuccess: () => {
      notificationService.success('You have successfully wrote a new post.');
      setResult({ status: 'OK', type: 'ADD_POST' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'OK', type: 'ADD_POST' });
    }
  });
  const addPost = (createPost: any) => addPostMutation.mutate(createPost);

  const updatePostMutation = useMutation((updatePost: any) => postsService.update(existingPost?.id ?? '', updatePost), {
    onSuccess: () => {
      notificationService.success('You have successfully updated an existing post.');
      setResult({ status: 'OK', type: 'UPDATE_POST' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'OK', type: 'UPDATE_POST' });
    }
  });
  const updatePost = (updatePost: any) => updatePostMutation.mutate(updatePost);

  const classes = useStyles();

  return (
    <>
      <Form
          initialValue={existingPost}
          schema={schema}
          onSubmit={isEdit ? updatePost : addPost}
        >
          <FormTextAreaInput label="Text" name="text" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </>
  )
}