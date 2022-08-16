import { FC, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useReportsResult } from "../../hooks";
import { Button, Container, Modal } from "../../imports";
import { ReportUserForm } from "./ReportUserForm";

const useStyles = createUseStyles({
  container: {
    '& button': {
      margin: '0em 0.5em 0.5em 0.5em'
    }
  },
  buttonReportUser: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginTop: '-50px'
  }
});

export const Profile: FC = () => {

  const [isReportUserOpen, setIsReportUserOpen] = useState(false);
  
  const classes = useStyles();

  const { result, setResult } = useReportsResult();

  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && result.type === 'REPORT_USER') {
      setIsReportUserOpen(false);
    }

    setResult(undefined);
  }, [result]);

  return (
    <>

    <div className={classes.container}>
        <Container>
          
        <Modal title="Report User" open={isReportUserOpen} onClose={() => setIsReportUserOpen(false)}>
          <ReportUserForm />
        </Modal>

        <div className={classes.buttonReportUser}>
          <Button onClick={() => { setIsReportUserOpen(true)} }>Report User</Button>
        </div> 

        </Container>
      </div>      
    </>
  );
}