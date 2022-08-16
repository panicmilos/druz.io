import { AxiosError } from 'axios';
import { createUseStyles } from 'react-jss';
import { useMutation } from 'react-query';
import { useParams } from 'react-router-dom';
import * as Yup from 'yup';
import { useReportsResult } from '../../hooks';
import { ALPHANUMERIC_REGEX, Button, extractErrorMessage, Form, FormTextInput, useNotificationService } from '../../imports';
import { useUserService } from '../../services';


const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});


export const ReportUserForm = () => {

  const { id } = useParams();

  const userService = useUserService();
  const notificationService = useNotificationService();
  
  const { setResult } = useReportsResult();

  const schema = Yup.object().shape({
    Reason: Yup.string()
      .required(() => ({ name: "Reason must be provided." })) 
      .matches(ALPHANUMERIC_REGEX, () => ({Reason: "Must be a valid reason."}))
  });

  const reportUserMutation = useMutation((Reason: string) => userService.report(id || '', Reason), {
    onSuccess: () => {
      notificationService.success('You have successfully reported that user.');
      setResult({ status: 'OK', type: 'REPORT_USER' })
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'REPORT_USER' })
    }
  });
  const reportUser = ({ Reason }: any) => reportUserMutation.mutate(Reason);

  
  const classes = useStyles();

  return (
    <Form
        schema={schema}
        onSubmit={reportUser}
      >
        <FormTextInput label="Reason" name="Reason" />

        <Button className={classes.submitButton} type="submit">Submit</Button>

      </Form>
  );
}