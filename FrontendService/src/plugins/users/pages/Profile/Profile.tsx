import { AxiosError } from "axios";
import { FC, useContext, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useMutation, useQuery } from "react-query";
import { useParams } from "react-router-dom";
import { useReportsResult } from "../../hooks";
import { ADMIN_ROLE, AuthContext, Button, ConfirmationModal, Container, extractErrorMessage, Modal, useNotificationService } from "../../imports";
import { useFriendRequestsService, useUserBlocksService, useUserFriendsService, useUserService } from "../../services";
import { AdditionalProfileForm } from "./AdditionalProfileForm";
import { ProfileForm } from "./ProfileForm";
import { ReportUserForm } from "./ReportUserForm";

const useStyles = createUseStyles({
  container: {
    margin: '2% 3% 0% 3%',
    '& button': {
      margin: '0.5em 0.5em 0.5em 0.5em'
    }
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginTop: '20px'
  },
  notFoundContainer: {
    margin: '2% 3% 0% 3%',
    textAlign: 'center'
  }
});

export const Profile: FC = () => {
  
  const { id } = useParams();
  const authContext = useContext(AuthContext);

  const userService = useUserService();
  const userBlocksService = useUserBlocksService(authContext.user?.ID ?? '');
  const friendRequestsService = useFriendRequestsService(authContext.user?.ID ?? '');
  const userFriendsService = useUserFriendsService(authContext.user?.ID ?? '');
  const notificationService = useNotificationService();

  const [isAddFriendOpen, setIsAddFriendOpen] = useState(false);
  const [isRemoveFriendOpen, setIsRemoveFriendOpen] = useState(false);
  const [isBlockUserOpen, setIsBlockUserOpen] = useState(false);
  const [isReportUserOpen, setIsReportUserOpen] = useState(false);

  const { result, setResult } = useReportsResult();

  
  const userData = useQuery([id, result, userService], () => userService.fetch(id || ''), { enabled: !!id });
  const userFriendData = useQuery([id, result, userFriendsService], () => userFriendsService.fetchById(id || ''), { enabled: !!id && authContext?.user?.Role !== ADMIN_ROLE });

  const user = userData?.data;
  const userFriend = userFriendData?.data;

  const blockUserMutation = useMutation(() => userBlocksService.add(user?.ID ?? ''), {
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

  const addFriendMutation = useMutation(() => friendRequestsService.add(user?.ID ?? ''), {
    onSuccess: () => {
      notificationService.success('You have successfully added user.');
      setResult({ status: 'OK', type: 'ADD_FRIEND' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'ADD_FRIEND' });
    }
  });
  const addFriend = () => addFriendMutation.mutate();

  const removeFriendMutation = useMutation(() => userFriendsService.delete(user?.ID ?? ''), {
    onSuccess: () => {
      notificationService.success('You have successfully removed user from the friend list.');
      setResult({ status: 'OK', type: 'REMOVE_FRIEND' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'REMOVE_FRIEND' });
    }
  });
  const removeFriend = () => removeFriendMutation.mutate();



  useEffect(() => {
    if (!result) return;
    
    if (result.status === 'OK' && result.type === 'REPORT_USER') {
      setIsReportUserOpen(false);
    }

    if (result.status === 'OK' && result.type === 'BLOCK_USER') {
      setIsBlockUserOpen(false);
    }

    if (result.status === 'OK' && result.type === 'ADD_FRIEND') {
      setIsAddFriendOpen(false);
    }

    if (result.status === 'OK' && result.type === 'REMOVE_FRIEND') {
      setIsRemoveFriendOpen(false);
    }
    setResult(undefined);
  }, [result]);
  
  const classes = useStyles();

  return (
    <>
      {
        user ?
          <div className={classes.container}>

            {
              (id !== authContext.user?.ID + '' && authContext.user?.Role !== ADMIN_ROLE) ?
                <Container>
                  
                  <Modal title="Report User" open={isReportUserOpen} onClose={() => setIsReportUserOpen(false)}>
                    <ReportUserForm />
                  </Modal>

                  <ConfirmationModal title="Block User" open={isBlockUserOpen} onClose={() => setIsBlockUserOpen(false)} onYes={blockUser}>
                    <p>Are you sure you want to block this user?</p>
                  </ConfirmationModal>

                  <ConfirmationModal title="Add User" open={isAddFriendOpen} onClose={() => setIsAddFriendOpen(false)} onYes={addFriend}>
                    <p>Are you sure you want to add this user?</p>
                  </ConfirmationModal>

                  <ConfirmationModal title="Remove User" open={isRemoveFriendOpen} onClose={() => setIsRemoveFriendOpen(false)} onYes={removeFriend}>
                    <p>Are you sure you want to remove this user from friend list?</p>
                  </ConfirmationModal>

                  <div className={classes.buttons}>
                    {
                      !!userFriend ?
                        <Button onClick={() => { setIsRemoveFriendOpen(true)} }>Remove User</Button> :
                        <Button onClick={() => { setIsAddFriendOpen(true)} }>Add User</Button>
                    }
                    <Button onClick={() => { setIsBlockUserOpen(true)} }>Block User</Button>
                    <Button onClick={() => { setIsReportUserOpen(true)} }>Report User</Button>
                  </div>

                </Container> : <></>
            }

            <Container>
              <img src={user.Image} style={{maxHeight: '400px'}} />

              <ProfileForm user={user} />
            </Container>

            <br />
            <br />
            
            <AdditionalProfileForm user={user} />

          </div>
          :
          <div className={classes.notFoundContainer}>
            <p>Profile can not be shown becase:</p>
            <p>- Profile does not exists.</p>
            <p>- You have blocked profile.</p>
            <p>- You are blocked by profile.</p>
          </div>
          
      }
    </>
  );
}