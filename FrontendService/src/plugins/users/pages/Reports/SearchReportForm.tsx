import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Button, Form, FormTextInput } from "../../imports";

type Props = {
  onSearch: (params: any) => void
}

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '0.5em'
  }
});


export const SearchReportForm: FC<Props> = ({ onSearch }) => {

  const classes = useStyles();

  return (
    <Form
      schema={undefined}
      onSubmit={onSearch}
    >
      <FormTextInput label="Reported" name="Reported" />
      <FormTextInput label="Reported By" name="ReportedBy" />
      <FormTextInput label="Reason" name="Reason" />

      <Button className={classes.submitButton} type="submit">Submit</Button>

    </Form>
  );
}