import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Button, Form, FormSelectOptionInput, FormTextInput } from "../../imports";


type Props = {
  onSearch: (params: any) => void
}

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});

const GenderOptions = [
  { label: 'Male', value: 0 },
  { label: 'Female', value: 1 },
  { label: 'Other', value: 2 }
]

export const SearchUserForm: FC<Props> = ({ onSearch }) => {

  const classes = useStyles();

  return (
    <Form
      schema={undefined}
      onSubmit={(values) => {
        onSearch({
          ...values,
          ...(values.Gender ? { Gender: +values.Gender } : {})
        });
      }}
    >
      <FormTextInput label="Name" name="Name" />
      <FormSelectOptionInput label="Gender" name="Gender" options={GenderOptions} />
      <FormTextInput label="Live Place" name="LivePlace" />
      <FormTextInput label="Work Place" name="WorkPlace" />
      <FormTextInput label="Education" name="Education" />
      <FormTextInput label="Interes" name="Interes" />

      <Button className={classes.submitButton} type="submit">Submit</Button>

    </Form>
  );
}