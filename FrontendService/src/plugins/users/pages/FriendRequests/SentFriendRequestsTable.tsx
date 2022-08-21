import { AxiosError } from "axios";
import { FC, useContext, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
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

export const SentFriendRequestsTable: FC<Props> = ({ friendRequests }) => {

  const [isDeleteOpen, setIsDeleteOpen] = useState(false);
  const [selectedRequest, setSelectedRequest] = useState<FriendRequest|undefined>();

  const { user } = useContext(AuthContext);

  const friendRequestsService = useFriendRequestsService(user?.ID ?? '');
  const notificationService = useNotificationService();
  
  const { result, setResult } = useFriendRequestsResult();

  const deleteRequestMutation = useMutation(() => friendRequestsService.deleteSent(selectedRequest?.FriendId + ''), {
    onSuccess: () => {
      notificationService.success('You have successfully removed friend request.');
      setResult({ status: 'OK', type: 'DELETE_FRIEND_REQUEST' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'DELETE_FRIEND_REQUEST' });
    }
  });
  const deleteRequest = () => deleteRequestMutation.mutate();

  const ActionsButtonGroup = ({ friendRequest }: any) =>
  <>
    <Button onClick={() => { setSelectedRequest(friendRequest); setIsDeleteOpen(true); }}>Delete</Button>
  </>

  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && result.type === 'DELETE_FRIEND_REQUEST') {
      setIsDeleteOpen(false);
    }

  }, [result]);

  const classes = useStyles();

  const nav = useNavigate();
  const navigateToUser = (userId: number) => {
    nav(`/users/${userId}/`)
  }

  return (
    <div className={classes.container}>

      <ConfirmationModal title="Delete Friend Request" open={isDeleteOpen} onClose={() => setIsDeleteOpen(false)} onYes={deleteRequest}>
        <p>Are you sure you want to delete this friend request?</p>
      </ConfirmationModal>

      <Table hasPagination={false}>
        <TableHead columns={['Sent to', 'Action']}/>
        <TableBody>
          {
            friendRequests?.map((friendRequest: FriendRequest) => 
            <TableRow 
              key={friendRequest.ID}
              cells={[
                <div onClick={() => navigateToUser(friendRequest.FriendId)} className={classes.nameContainer}><img src={friendRequest.Friend.Image || '/images/no-image.png'} /><p>{friendRequest.Friend.FirstName} {friendRequest.Friend.LastName}</p></div>,
                <ActionsButtonGroup friendRequest={friendRequest} />
            ]}/>
            )
          }
        </TableBody>
      </Table>
    </div>
    );
  }