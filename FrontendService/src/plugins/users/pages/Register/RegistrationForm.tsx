import { FC } from "react";
import { useNotificationService, ALPHANUMERIC_REGEX, EMAIL_REGEX, extractErrorMessage, Card, FormTextInput, Button, Form, FormDateInput, FormSelectOptionInput } from "../../imports";
import { useUserService } from "../../services";
import * as Yup from 'yup';
import { useMutation } from "react-query";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";
import moment from "moment";


function subtractYears(numOfYears: number, date = new Date()) {
  date.setFullYear(date.getFullYear() - numOfYears);

  return date;
}

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '1em',
  }
});


const GenderOptions = [
  { label: 'Male', value: 0 },
  { label: 'Female', value: 1 },
  { label: 'Other', value: 2 }
]

export const RegistrationForm: FC = () => {
  
  const usersService = useUserService();
  const notificationService = useNotificationService();


  const schema = Yup.object().shape({
    Email: Yup.string()
      .required(() => ({ Email: "Email must be provided." })) 
      .matches(EMAIL_REGEX, () => ({Email: "Must be a valid email."})),
    Password: Yup.string().required(() => ({Password: "Password must be provided."}))
      .matches(/[A-Z]/, ()=> ({Password: "Password must contain capital letters."}))
      .matches(/[a-z]/, ()=> ({Password: "Password must contain lower letters."}))
      .matches(/[0-9]/, ()=> ({Password: "Password must contain numbers."}))
      .matches(/[^a-zA-Z0-9]/, ()=> ({Password: "Password must contain special characters."})),
    ConfirmPassword: Yup.string().oneOf(
      [Yup.ref('Password'), null],
      () => ({ConfirmPassword: "Passwords don't match!" })
      ).required(() => ({ConfirmPassword: "Password must be provided."})),
    FirstName: Yup.string()
      .required(() => ({ FirstName: "First name must be provided." })) 
      .matches(ALPHANUMERIC_REGEX, () => ({FirstName: "Must be a valid first name."})),
    LastName: Yup.string()
      .required(() => ({ LastName: "Last name must be provided." })) 
      .matches(ALPHANUMERIC_REGEX, () => ({LastName: "Must be a valid last name."})),
    Birthday: Yup.date()
      .required(() => ({ Birthday: "Birthday must be provided." })) 
      .max(subtractYears(13), () => ({Birthday: "Must be at least 13 year old."})),
    Gender: Yup.number()
      .required(() => ({ Gender: "Gender must be provided." })) 
  });

  const registerUserMutation = useMutation((createUser: any) => usersService.add(createUser), {
    onSuccess: () => {
      notificationService.success('You have successfully created new profile.');
    },
    onError: (error: AxiosError) => {
      console.log(error.response?.data);
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const registerUser = (createUser: any) => registerUserMutation.mutate(createUser);

  const classes = useStyles();

  return (
    <Card title="Register">
      <Form
          schema={schema}
          onSubmit={(values: any) => {
            const createUser = {
              Email: values.Email,
              Password: values.Password,
              Profile: {
                FirstName: values.FirstName,
                LastName: values.LastName,
                Birthday: moment(values.Birthday).format(),
                Gender: +values.Gender
              }
            };

            registerUser(createUser);
          }}
        >
          <FormTextInput type="text" label="Email" name="Email" />
          <FormTextInput type="password" label="Password" name="Password" />
          <FormTextInput type="password" label="Confirm password" name="ConfirmPassword" />

          <FormTextInput type="text" label="First Name" name="FirstName" />
          <FormTextInput type="text" label="Last Name" name="LastName" />
          <FormDateInput label="Birthday" name="Birthday" />
          <FormSelectOptionInput label="Gender" options={GenderOptions} name="Gender" />

          <Button className={classes.submitButton} type="submit">Submit</Button>
        </Form>
    </Card>
  );

}