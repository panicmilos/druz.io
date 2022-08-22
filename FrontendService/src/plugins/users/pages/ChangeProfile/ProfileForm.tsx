import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { Button, Form, FormDateInput, FormSelectOptionInput, FormTextInput, PHONE_NUMBER_REGEX, ALPHANUMERIC_REGEX, FormTextAreaInput, Card, useNotificationService, extractErrorMessage } from "../../imports";
import { Education, LivePlace, Profile, WorkPlace } from "../../models/User";
import * as Yup from 'yup';
import moment from "moment";
import { LivePlacesForm } from "./LivePlacesForm";
import { WorkPlacesForm } from "./WorkPlacesForm";
import { EducationsForm } from "./EducationsForm";
import { InteresesForm } from "./Intereses";
import { useUserService } from "../../services";
import { useMutation } from "react-query";
import { AxiosError } from "axios";

type Props = {
  user?: Profile;
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

function IsMoreThenYears(numOfYears: number, date: Date) {
  const secondDate = new Date();
  secondDate.setFullYear(secondDate.getFullYear() - numOfYears);

  return secondDate < date;
}

export const ProfileForm: FC<Props> = ({ user }) => {

  const userService = useUserService();
  const notificationService = useNotificationService();
  
  const [livePlaces, setLivePlaces] = useState(user?.LivePlaces ?? []);
  const [workPlaces, setWorkPlaces] = useState(user?.WorkPlaces ?? []);
  const [educations, setEducations] = useState(user?.Educations ?? []);
  const [intereses, setIntereses] = useState(user?.Intereses ?? []);

  const classes = useStyles();

  const schema = Yup.object().shape({
    FirstName: Yup.string()
      .required(() => ({ FirstName: "First name must be provided." })) 
      .matches(ALPHANUMERIC_REGEX, () => ({FirstName: "Must be a valid first name."})),
    LastName: Yup.string()
      .required(() => ({ LastName: "Last name mustbe provided." })) 
      .matches(ALPHANUMERIC_REGEX, () => ({LastName: "Must be a valid last name."})),
    Birthday: Yup.string()
      .required(() => ({ Birthday: "Birthday must be provided." })) 
      .test('be>13', () => ({Birthday: "Must be at least 13 year old."}), (v: any) => !IsMoreThenYears(13, new Date(v))),
    Gender: Yup.string()
      .required(() => ({ Gender: "Gender must be provided." })),
    About: Yup.string()
      .required(() => ({ About: "About must be provided." })) ,
    PhoneNumber: Yup.string()
      .required(() => ({ PhoneNumber: "Phone number must be provided." })) 
      .matches(PHONE_NUMBER_REGEX, () => ({PhoneNumber: "Must be a valid phone number."})),
  });

  const updateProfileMutator = useMutation((profile: Profile) => userService.update(profile?.ID, profile), {
    onSuccess: (_) => {
      notificationService.success("You have successfully updated your profile.");
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const updateProfile = (profile: Profile) => updateProfileMutator.mutate(profile);

  return (

    <Form
      initialValue={{ ...user, Gender: user?.Gender + '', Birthday: moment(user?.Birthday).format('yyyy-MM-DD')}}
      schema={schema}
      onSubmit={(values: any) => {
        const updateUser = {
          ID: user?.ID,
          ...values,
          Birthday: moment(values.Birthday).format(),
          Gender: +values.Gender,
          LivePlaces: livePlaces || [],
          WorkPlaces: workPlaces?.map((wp: WorkPlace) => ({ ...wp, From: moment(wp.From).format(), To: moment(wp.To).format() })) ?? [],
          Educations: educations?.map((e: Education) => ({ ...e, From: moment(e.From).format(), To: moment(e.To).format() })) ?? [],
          Intereses: intereses || []
        }

        updateProfile(updateUser);
      }}
    >
      <FormTextInput type="text" label="First Name" name="FirstName" />
      <FormTextInput type="text" label="Last Name" name="LastName" />
      <FormDateInput label="Birthday" name="Birthday" />
      <FormSelectOptionInput label="Gender" options={GenderOptions} name="Gender" />

      <FormTextAreaInput label="About" name="About" />
      <FormTextInput type="text" label="Phone Number" name="PhoneNumber" />

      <Card title={"Live Places"} >
        <LivePlacesForm livePlaces={livePlaces} onChange={setLivePlaces} />
      </Card>

      <Card title={"Work Places"} >
        <WorkPlacesForm workPlaces={workPlaces} onChange={setWorkPlaces} />
      </Card>

      <Card title={"Educations"} >
        <EducationsForm educations={educations} onChange={setEducations} />
      </Card>

      <Card title={"Intereses"} >
        <InteresesForm intereses={intereses} onChange={setIntereses} />
      </Card>

      <Button className={classes.submitButton} type="submit">Submit</Button>
    </Form>

  );
}