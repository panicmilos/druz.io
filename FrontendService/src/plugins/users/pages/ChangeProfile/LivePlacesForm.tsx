import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { Button, Container, SelectOptionInput, Table, TableBody, TableHead, TableRow, TextInput } from "../../imports";
import { LivePlace } from "../../models/User";

type Props = {
  livePlaces: LivePlace[],
  onChange: (livePlaces: LivePlace[]) => any
}

const useStyles = createUseStyles(() => ({
  button: {
    alignSelf: 'center',
    padding: '0.25em',
    flexGrow: 0
  },
  container: 'margin-bottom: 0.5em'
}));

const LivePlaceOptions = [
  { label: 'Currently', value: 0 },
  { label: 'Lived', value: 1 },
  { label: 'Birthplace', value: 2 }
]

const LivePlaceTypes = ['Currently', 'Lived', 'Birthplace'];

export const LivePlacesForm: FC<Props> = ({ livePlaces, onChange }) => {

  const [localLivePlaces, setLocalLivePlaces] = useState(livePlaces);

  const [place, setPlace] = useState('');
  const [type, setType] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const classes = useStyles();

  const handleAdd = () => {
    const existingPlaces = localLivePlaces.map(lp => lp.Place);
    if (existingPlaces.includes(place)) {
      setErrorMessage('Place already exists.');
      return;
    }

    if (!place) {
      setErrorMessage('Place must be provided.');
      return;
    }

    if (!type) {
      setErrorMessage('Type must be provided.');
      return;
    }

    const newLivePlaces = [...localLivePlaces, { Place: place, LivePlaceType: +type } as LivePlace];
    setLocalLivePlaces(newLivePlaces)
    onChange(newLivePlaces);

    resetState();
  };

  const resetState = () => {
    setPlace('');
    setType('');
    setErrorMessage('');
  }

  const handleRemove = (place: string) => {
    const newLivePlaces = localLivePlaces.filter(lp => lp.Place !== place);
    setLocalLivePlaces(newLivePlaces)
    onChange(newLivePlaces);
  }
  
  return (
    <>
      <Container className={classes.container}>

        <TextInput label="Place" value={place} onChange={setPlace} />
        <SelectOptionInput label="Type" value={type} onChange={setType} options={LivePlaceOptions} />
        <Button className={classes.button} onClick={handleAdd}>Add</Button>
      </ Container>
      {errorMessage}

      <Table hasPagination={false}>
        <TableHead columns={['Place', 'Type', 'Action']}/>
        <TableBody>
        {
          localLivePlaces?.map((livePlace, i) =>
            <TableRow
              key={i}          
              cells={[
                livePlace.Place,
                LivePlaceTypes[livePlace.LivePlaceType],
                <Button type="button" onClick={() => handleRemove(livePlace.Place)}>Remove</Button>
              ]}
            />
          )
        }
        </TableBody>
      </Table>
    </>
  )
}