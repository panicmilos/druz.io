import { AxiosError } from "axios";
import { FC, useContext, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useMutation } from "react-query";
import { useFriendRequestsResult } from "../../hooks";
import { AuthContext, Button, ConfirmationModal, extractErrorMessage, Table, TableBody, TableHead, TableRow, useNotificationService } from "../../imports";
import { FriendRequest } from "../../models/FriendRequest";
import { useFriendRequestsService } from "../../services";

type Props = {
  friendRequests: FriendRequest[]
}

const useStyles = createUseStyles({
  container: {
    '& button': {
      margin: '0em 0.5em 0.5em 0.5em'
    }
  }
});


export const ReceivedFriendRequestsTable: FC<Props> = ({ friendRequests }) => {

  const [isAcceptRequestOpen, setIsAcceptRequestOpen] = useState(false);
  const [isDeclineRequestOpen, setIsDeclineRequestOpen] = useState(false);
  const [selectedRequest, setSelectedRequest] = useState<FriendRequest|undefined>();

  const { user } = useContext(AuthContext);

  const friendRequestsService = useFriendRequestsService(user?.ID ?? '');
  const notificationService = useNotificationService();
  
  const { result, setResult } = useFriendRequestsResult();

  const acceptUserMutation = useMutation(() => friendRequestsService.accept(selectedRequest?.UserId + ''), {
    onSuccess: () => {
      notificationService.success('You have successfully accepted user.');
      setResult({ status: 'OK', type: 'ACCEPT_USER' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'ACCEPT_USER' });
    }
  });
  const acceptUser = () => acceptUserMutation.mutate();

  const declineUserMutation = useMutation(() => friendRequestsService.decline(selectedRequest?.UserId + ''), {
    onSuccess: () => {
      notificationService.success('You have successfully declined user.');
      setResult({ status: 'OK', type: 'DECLINE_USER' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'DECLINE_USER' });
    }
  });
  const declineUser = () => declineUserMutation.mutate();
  
  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && result.type === 'ACCEPT_USER') {
      setIsAcceptRequestOpen(false);
    }

    if (result.status === 'OK' && result.type === 'DECLINE_USER') {
      setIsDeclineRequestOpen(false);
    }

  }, [result]);

  const ActionsButtonGroup = ({ friendRequest }: any) =>
    <>
      <Button onClick={() => { setSelectedRequest(friendRequest); setIsAcceptRequestOpen(true); }}>Accept</Button>
      <Button onClick={() => { setSelectedRequest(friendRequest); setIsDeclineRequestOpen(true); }}>Decline</Button>
    </>

  const classes = useStyles();

  return (
    <div className={classes.container}>

      <ConfirmationModal title="Accept User" open={isAcceptRequestOpen} onClose={() => setIsAcceptRequestOpen(false)} onYes={acceptUser}>
        <p>Are you sure you want to accept this user?</p>
      </ConfirmationModal>

      <ConfirmationModal title="Decline User" open={isDeclineRequestOpen} onClose={() => setIsDeclineRequestOpen(false)} onYes={declineUser}>
        <p>Are you sure you want to decline this user?</p>
      </ConfirmationModal>

      <Table hasPagination={false}>
        <TableHead columns={['Received from', 'Action']}/>
        <TableBody>
          {
            friendRequests?.map((friendRequest: FriendRequest) => 
            <TableRow 
              key={friendRequest.ID}
              cells={[
                `${friendRequest.User.FirstName} ${friendRequest.User.LastName}`,
                <ActionsButtonGroup friendRequest={friendRequest} />
            ]}/>
            )
          }
        </TableBody>
      </Table>
    </div>
    );
  }