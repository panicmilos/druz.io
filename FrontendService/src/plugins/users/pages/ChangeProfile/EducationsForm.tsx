import moment from "moment";
import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { DateInput } from "../../../../core";
import { Button, Container, Table, TableBody, TableHead, TableRow, TextInput } from "../../imports";
import { Education } from "../../models/User";

type Props = {
  educations: Education[],
  onChange: (educations: Education[]) => any
}

const useStyles = createUseStyles(() => ({
  button: {
    alignSelf: 'center',
    padding: '0.25em',
    flexGrow: 0
  },
  container: 'margin-bottom: 0.5em'
}));


export const EducationsForm: FC<Props> = ({ educations, onChange }) => {

  const [localEducations, setLocalEducations] = useState(educations);

  const [place, setPlace] = useState('');
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const classes = useStyles();

  const handleAdd = () => {
    const existingEducation = localEducations.map(e => e.Place);
    if (existingEducation.includes(place)) {
      setErrorMessage('Place already exists.');
      return;
    }

    if (!place) {
      setErrorMessage('Place must be provided.');
      return;
    }

    if (!from) {
      setErrorMessage('From must be provided.');
      return;
    }

    if (!to) {
      setErrorMessage('To must be provided.');
      return;
    }

    if (from > to) {
      setErrorMessage('From must be before To.');
      return;
    }

    const newEducations = [...localEducations, { Place: place, From: from, To: to } as Education];
    setLocalEducations(newEducations)
    onChange(newEducations);

    resetState();
  };

  const resetState = () => {
    setPlace('');
    setFrom('');
    setTo('');
    setErrorMessage('');
  }

  const handleRemove = (place: string) => {
    const newEducations = localEducations.filter(e => e.Place !== place);
    setLocalEducations(newEducations);
    onChange(newEducations);
  }
  
  return (
    <>
      <Container className={classes.container}>

        <TextInput label="Place" value={place} onChange={setPlace} />
        <DateInput label="From" value={from} onChange={setFrom} />
        <DateInput label="To" value={to} onChange={setTo} />
        <Button className={classes.button} onClick={handleAdd}>Add</Button>
      </ Container>
      {errorMessage}

      <Table hasPagination={false}>
        <TableHead columns={['Place', 'From', 'To', 'Action']}/>
        <TableBody>
        {
          localEducations?.map((education, i) =>
            <TableRow
              key={i}          
              cells={[
                education.Place,
                moment(education.To).format('yyyy-MM-DD'),
                moment(education.From).format('yyyy-MM-DD'),
                <Button type="button" onClick={() => handleRemove(education.Place)}>Remove</Button>
              ]}
            />
          )
        }
        </TableBody>
      </Table>
    </>
  )
}