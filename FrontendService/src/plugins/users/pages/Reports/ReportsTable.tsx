import { AxiosError } from "axios";
import { IncomingMessage } from "http";
import { FC, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { useReportsResult } from "../../hooks";
import { Button, ConfirmationModal, extractErrorMessage, Table, TableBody, TableHead, TableRow, useNotificationService } from "../../imports";
import { Report } from "../../models/Report";
import { useUserReportsService, useUserService } from "../../services";

type Props = {
  reports: Report[]
}

const useStyles = createUseStyles({
  container: {
    '& button': {
      margin: '0em 0.5em 0.5em 0.5em'
    }
  },
  nameContainer: {
    display: 'flex',
    alignItems: 'center',
    '& p': {
      marginLeft: '5px',
    },
    '& img': {
      width: '36px',
      height: '36px',
      borderRadius: '50%'
    }
  },
  
});

export const ReportsTable: FC<Props> = ({ reports }) => {

  const [isIgnoreReportOpen, setIsIgnoreReportOpen] = useState(false);
  const [isBlockUserOpen, setIsBlockUserOpen] = useState(false);
  const [selectedReport, setSelectedReport] = useState<Report|undefined>();

  const userReportsService = useUserReportsService();
  const usersService = useUserService();
  const notificationService = useNotificationService();

  const { result, setResult } = useReportsResult();

  const ignoreReportMutation = useMutation(() => userReportsService.ignore(selectedReport?.ID ?? ''), {
    onSuccess: () => {
      notificationService.success('You have successfully ignored report.');
      setResult({ status: 'OK', type: 'IGNORE_REPORT' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'IGNORE_REPORT' });
    }
  });
  const ignoreReport = () => ignoreReportMutation.mutate();

  const blockUserMutation = useMutation(() => usersService.block(selectedReport?.ReportedId + ''), {
    onSuccess: () => {
      notificationService.success('You have successfully blocked user.');
      setResult({ status: 'OK', type: 'BLOCK_USER' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'BLOCK_USER' });
    }
  });
  const blockUser = () => blockUserMutation.mutate();

  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && result.type === 'IGNORE_REPORT') {
      setIsIgnoreReportOpen(false);
    }

    if (result.status === 'OK' && result.type === 'BLOCK_USER') {
      setIsBlockUserOpen(false);
    }

  }, [result]);
  

  const ActionsButtonGroup = ({ report }: any) =>
  <>
    <Button onClick={() => { setSelectedReport(report); setIsIgnoreReportOpen(true); }}>Ignore</Button>
    <Button onClick={() => { setSelectedReport(report); setIsBlockUserOpen(true); }}>Block</Button>
  </>

  const nav = useNavigate();
  const navigateToUser = (userId: number) => {
    nav(`/users/${userId}/`)
  }

  const classes = useStyles();


  return (
    <div className={classes.container}>

      <ConfirmationModal title="Ingore Report" open={isIgnoreReportOpen} onClose={() => setIsIgnoreReportOpen(false)} onYes={ignoreReport}>
        <p>Are you sure you want to ignore this report?</p>
      </ConfirmationModal>

      <ConfirmationModal title="Block User" open={isBlockUserOpen} onClose={() => setIsBlockUserOpen(false)} onYes={blockUser}>
        <p>Are you sure you want to block this user?</p>
      </ConfirmationModal>

      <Table hasPagination={false}>
        <TableHead columns={['Reported', 'Reported By', 'Reason', 'Action']}/>
        <TableBody>
          {
            reports?.map((report: Report) => 
            <TableRow 
              key={report.ID}
              cells={[
                <div onClick={() => navigateToUser(report.ReportedId)} className={classes.nameContainer}><img src={report.Reported.Image || '/images/no-image.png'} /><p>{report.Reported.FirstName} {report.Reported.LastName}</p></div>,
                <div onClick={() => navigateToUser(report.ReportedById)} className={classes.nameContainer}><img src={report.ReportedBy.Image || '/images/no-image.png'} /><p>{report.ReportedBy.FirstName} {report.ReportedBy.LastName}</p></div>,
                report.Reason,
                <ActionsButtonGroup report={report}/>
            ]}/>
            )
          }
        </TableBody>
      </Table>
    </div>
  );
}