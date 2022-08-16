import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { DateInput } from "../../../../core";
import { Button, Container, Table, TableBody, TableHead, TableRow, TextInput } from "../../imports";
import { WorkPlace } from "../../models/User";

type Props = {
  workPlaces: WorkPlace[],
  onChange: (workPlaces: WorkPlace[]) => any
}

const useStyles = createUseStyles(() => ({
  button: {
    alignSelf: 'center',
    padding: '0.25em',
    flexGrow: 0
  },
  container: 'margin-bottom: 0.5em'
}));


export const WorkPlacesForm: FC<Props> = ({ workPlaces, onChange }) => {

  const [localWorkPlaces, setLocalWorkPlaces] = useState(workPlaces);

  const [place, setPlace] = useState('');
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const classes = useStyles();

  const handleAdd = () => {
    const existingPlaces = localWorkPlaces.map(wp => wp.Place);
    if (existingPlaces.includes(place)) {
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

    const newWorkPlaces = [...localWorkPlaces, { Place: place, From: from, To: to } as WorkPlace];
    setLocalWorkPlaces(newWorkPlaces)
    onChange(newWorkPlaces);

    resetState();
  };

  const resetState = () => {
    setPlace('');
    setFrom('');
    setTo('');
    setErrorMessage('');
  }

  const handleRemove = (place: string) => {
    const newWorkPlaces = localWorkPlaces.filter(wp => wp.Place !== place);
    setLocalWorkPlaces(newWorkPlaces);
    onChange(newWorkPlaces);
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
          localWorkPlaces?.map((workPlace, i) =>
            <TableRow
              key={i}          
              cells={[
                workPlace.Place,
                workPlace.To,
                workPlace.From,
                <Button type="button" onClick={() => handleRemove(workPlace.Place)}>Remove</Button>
              ]}
            />
          )
        }
        </TableBody>
      </Table>
    </>
  )
}