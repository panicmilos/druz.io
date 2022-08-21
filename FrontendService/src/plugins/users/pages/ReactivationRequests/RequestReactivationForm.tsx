import { FC } from "react";
import * as Yup from 'yup';
import { useUserReactivationService } from "../../services";
import { Button, Card, extractErrorMessage, Form, FormTextInput, useNotificationService, EMAIL_REGEX } from "../../imports";
import { useMutation } from "react-query";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});

export const RequestReactivationForm: FC = () => {

  const schema = Yup.object().shape({
    Email: Yup.string()
      .required(() => ({ Email: "Email must be provided." })) 
      .matches(EMAIL_REGEX, () => ({Email: "Must be a valid email."}))
  });

  
  const userReactivationService = useUserReactivationService();
  const notificationService = useNotificationService();

  const requestReactivationMutation = useMutation((email: string) => userReactivationService.request(email), {
    onSuccess: () => {
      notificationService.success('You have successfully requested profile reactivation.');
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const requestReactivation = (email: string) => requestReactivationMutation.mutate(email);
  
  const classes = useStyles();

  return (
    <Card title="Profile Reactivation Request">
      <Form
          schema={schema}
          onSubmit={(values) => requestReactivation(values.Email)}
        >
          <FormTextInput label="Email" name="Email" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </Card>
  );
}