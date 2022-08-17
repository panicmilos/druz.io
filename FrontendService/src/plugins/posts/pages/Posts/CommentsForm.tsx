import { FC } from "react";
import { Button, extractErrorMessage, Form, FormTextAreaInput, useNotificationService } from "../../imports";
import { Comment } from "../../models/Post";
import * as Yup from 'yup';
import { createUseStyles } from "react-jss";
import { useCommentsService } from "../../services";
import { useMutation } from "react-query";
import { usePostsResult } from "../../hooks";
import { AxiosError } from "axios";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '1em',
  }
});

type Props = {
  postId: string,
  existingComment?: Comment,
  isEdit: boolean
}

export const CommentsForm: FC<Props> = ({ postId, existingComment = undefined, isEdit = false}) => {

  const commentsService = useCommentsService(postId);
  const notificationService = useNotificationService();

  const schema = Yup.object().shape({
    text: Yup.string()
      .required(() => ({ text: "Text must be provided." }))
  });

  const { setResult } = usePostsResult();

  const addCommentMutation = useMutation((createComment: any) => commentsService.add(createComment), {
    onSuccess: () => {
      notificationService.success('You have successfully wrote a new comment.');
      setResult({ status: 'OK', type: 'ADD_COMMENT' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'OK', type: 'ADD_COMMENT' });
    }
  });
  const addComment = (createComment: any) => addCommentMutation.mutate(createComment);

  const updateCommentMutation = useMutation((updateComment: any) => commentsService.update(existingComment?.id ?? '', updateComment), {
    onSuccess: () => {
      notificationService.success('You have successfully updated an existing comment.');
      setResult({ status: 'OK', type: 'UPDATE_COMMENT' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'OK', type: 'UPDATE_COMMENT' });
    }
  });
  const updateComment = (updateComment: any) => updateCommentMutation.mutate(updateComment);


  const classes = useStyles();

  return (
    <>
      <Form
          initialValue={existingComment}
          schema={schema}
          onSubmit={isEdit ? updateComment : addComment}
        >
          <FormTextAreaInput label="Text" name="text" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </>
  );
}