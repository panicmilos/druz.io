import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { Button, Container, Table, TableBody, TableHead, TableRow, TextInput } from "../../imports";
import { Interes } from "../../models/User";

type Props = {
  intereses: Interes[],
  onChange: (intereses: Interes[]) => any
}

const useStyles = createUseStyles(() => ({
  button: {
    alignSelf: 'center',
    padding: '0.25em',
    flexGrow: 0
  },
  container: 'margin-bottom: 0.5em'
}));


export const InteresesForm: FC<Props> = ({ intereses, onChange }) => {

  const [localIntereses, setLocalIntereses] = useState(intereses);

  const [interes, setInteres] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const classes = useStyles();

  const handleAdd = () => {
    const existingIntereses = localIntereses.map(i => i.Interes);
    if (existingIntereses.includes(interes)) {
      setErrorMessage('Interes already exists.');
      return;
    }

    if (!interes) {
      setErrorMessage('Place must be provided.');
      return;
    }

    const newIntereses = [...localIntereses, { Interes: interes } as Interes];
    setLocalIntereses(newIntereses)
    onChange(newIntereses);

    resetState();
  };

  const resetState = () => {
    setInteres('');
    setErrorMessage('');
  }

  const handleRemove = (interes: string) => {
    const newIntereses = localIntereses.filter(i => i.Interes !== interes);
    setLocalIntereses(newIntereses);
    onChange(newIntereses);
  }
  
  return (
    <>
      <Container className={classes.container}>

        <TextInput label="Interes" value={interes} onChange={setInteres} />
        <Button className={classes.button} onClick={handleAdd}>Add</Button>
      </ Container>
      {errorMessage}

      <Table hasPagination={false}>
        <TableHead columns={['Interes', 'Action']}/>
        <TableBody>
        {
          localIntereses?.map((interes, i) =>
            <TableRow
              key={i}          
              cells={[
                interes.Interes,
                <Button type="button" onClick={() => handleRemove(interes.Interes)}>Remove</Button>
              ]}
            />
          )
        }
        </TableBody>
      </Table>
    </>
  )
}