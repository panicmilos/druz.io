import moment from "moment";
import { FC } from "react";
import { Form, FormDateInput, FormSelectOptionInput, FormTextAreaInput, FormTextInput } from "../../imports";
import { Profile } from "../../models/User";

type Props = {
  user?: Profile
}

const GenderOptions = [
  { label: 'Male', value: 0 },
  { label: 'Female', value: 1 },
  { label: 'Other', value: 2 }
]

export const ProfileForm: FC<Props> = ({ user }) => {

  return (
    <>
      <Form
        schema={undefined}
        initialValue={{
          ...user,
          Gender: user?.Gender + '',
          Birthday: moment(user?.Birthday).format('yyyy-MM-DD')
        }}
      >
        <FormTextInput type="text" label="First Name" name="FirstName" disabled={true} />
        <FormTextInput type="text" label="Last Name" name="LastName" disabled={true} />
        <FormDateInput label="Birthday" name="Birthday" disabled={true} />
        <FormSelectOptionInput label="Gender" options={GenderOptions} name="Gender" disabled={true} />

        <FormTextAreaInput label="About" name="About" disabled={true} />
        <FormTextInput type="text" label="Phone Number" name="PhoneNumber" disabled={true} />
      </Form>

    </>
  )
}