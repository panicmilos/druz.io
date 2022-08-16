import { FC } from "react";
import { Profile } from "../../models/User";
import * as Yup from 'yup';
import { useUserService } from "../../services";
import { Button, Card, extractErrorMessage, Form, FormTextInput, useNotificationService } from "../../imports";
import { useMutation } from "react-query";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});

type Props = {
  user?: Profile;
}

export const ChangePasswordForm: FC<Props> = ({ user }) => {

  const schema = Yup.object().shape({
    CurrentPassword : Yup.string().required(() => ({CurrentPassword: "Current password must be provided."})),
    NewPassword: Yup.string().required(() => ({NewPassword: "New password must be provided."}))
                              .matches(/[A-Z]/, ()=> ({NewPassword: "New Password must contain capital letters."}))
                              .matches(/[a-z]/, ()=> ({NewPassword: "New Password must contain lower letters."}))
                              .matches(/[0-9]/, ()=> ({NewPassword: "New Password must contain numbers."}))
                              .matches(/[^a-zA-Z0-9]/, ()=> ({NewPassword: "New Password must contain special characters."})),

    ConfirmPassword: Yup.string().oneOf([Yup.ref('NewPassword'), null], () => ({ConfirmPassword: "Passwords don't match!" })).required(() => ({ConfirmPassword: "Password must be provided."}))
  });

  
  const usersService = useUserService();
  const notificationService = useNotificationService();

  const changePasswordMutation = useMutation((changePassword: any) => usersService.changePassword(user?.ID ?? '', changePassword), {
    onSuccess: () => {
      notificationService.success('You have successfully changed your password.');
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const changePassword = (changePassword: any) => changePasswordMutation.mutate(changePassword);
  
  const classes = useStyles();

  return (
    <Card title="Change Password">
      <Form
          schema={schema}
          onSubmit={changePassword}
        >
          <FormTextInput type="password" label="Current Password" name="CurrentPassword" />
          <FormTextInput type="password" label="New Password" name="NewPassword" />
          <FormTextInput type="password" label="Confirm password" name="ConfirmPassword" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </Card>
  );
}