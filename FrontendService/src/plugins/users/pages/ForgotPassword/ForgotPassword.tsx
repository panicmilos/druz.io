import { useNavigate, useParams, useSearchParams } from "react-router-dom";
import { Button, Card, extractErrorMessage, Form, FormTextInput, useNotificationService } from "../../imports";
import { usePasswordRecoveryService } from "../../services/PasswordRecoveryService";
import * as Yup from 'yup';
import { useMutation } from "react-query";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});

export const ForgotPassword = () => {

  const nav = useNavigate();
  const { id } = useParams();
  let [searchParams] = useSearchParams();
  const token = searchParams.get("token");

  const passwordRecoveryService = usePasswordRecoveryService();
  const notificationService = useNotificationService();

  const schema = Yup.object().shape({
    NewPassword: Yup.string().required(() => ({NewPassword: "New password must be provided."}))
                              .matches(/[A-Z]/, ()=> ({NewPassword: "New Password must contain capital letters."}))
                              .matches(/[a-z]/, ()=> ({NewPassword: "New Password must contain lower letters."}))
                              .matches(/[0-9]/, ()=> ({NewPassword: "New Password must contain numbers."}))
                              .matches(/[^a-zA-Z0-9]/, ()=> ({NewPassword: "New Password must contain special characters."})),

    ConfirmPassword: Yup.string().oneOf([Yup.ref('NewPassword'), null], () => ({ConfirmPassword: "Passwords don't match!" })).required(() => ({ConfirmPassword: "Password must be provided."}))
  });

  const changePasswordMutation = useMutation((newPassword: string) => passwordRecoveryService.recover(id || '', token || '', newPassword), {
    onSuccess: () => {
      notificationService.success('You have successfully changed your password.');
      nav('/');
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      nav('/');
    }
  });
  const changePassword = (newPassword: string) => changePasswordMutation.mutate(newPassword);
  
  const classes = useStyles();

  return (
    <Card title="Change Password">
      <Form
          schema={schema}
          onSubmit={(values) => changePassword(values.NewPassword)}
        >
          <FormTextInput type="password" label="New Password" name="NewPassword" />
          <FormTextInput type="password" label="Confirm password" name="ConfirmPassword" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </Card>
  );
}