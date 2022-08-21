import { FC } from "react";
import * as Yup from 'yup';
import { Button, Card, extractErrorMessage, Form, FormTextInput, useNotificationService, EMAIL_REGEX } from "../../imports";
import { useMutation } from "react-query";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";
import { usePasswordRecoveryService } from "../../services/PasswordRecoveryService";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});

export const ForgotPasswordForm: FC = () => {

  const schema = Yup.object().shape({
    Email: Yup.string()
      .required(() => ({ Email: "Email must be provided." })) 
      .matches(EMAIL_REGEX, () => ({Email: "Must be a valid email."}))
  });

  
  const passwordRecoveryService = usePasswordRecoveryService();
  const notificationService = useNotificationService();

  const requestPasswordRecoveryMutation = useMutation((email: string) => passwordRecoveryService.request(email), {
    onSuccess: () => {
      notificationService.success('You have successfully requested password recovery.');
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const requestPasswordRecovery = (email: string) => requestPasswordRecoveryMutation.mutate(email);
  
  const classes = useStyles();

  return (
    <Card title="Password Recovery Request">
      <Form
          schema={schema}
          onSubmit={(values) => requestPasswordRecovery(values.Email)}
        >
          <FormTextInput label="Email" name="Email" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </Card>
  );
}