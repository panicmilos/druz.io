import { FC } from "react"
import { createUseStyles } from "react-jss";
import { Button } from "./Button";
import { Modal } from "./Modal"

const useStyles = createUseStyles({
  buttonsGroup: {
    display: 'flex',
    justifyContent: 'flex-end',
    '& button': {
      marginLeft: '0.5em'
    }
  }
});
type Props = {
  title: string;
  onYes: () => any;
  open: boolean;
  onClose?: () => any;
}

export const ConfirmationModal: FC<Props> = ({title, onYes, open, onClose, children}) => {
  const classes = useStyles();

  return (
    <>
    <Modal title={title} open={open} onClose={onClose}>
      {children}
      <div className={classes.buttonsGroup}>
        <Button onClick={onYes}>Yes</Button>
        <Button onClick={onClose}>No</Button>
      </div>
    </Modal>
    </>
  )
}