import { FC, useContext } from "react";
import { Profile } from "../../models/User";
import { useUserService } from "../../services";
import { AuthContext, Button, extractErrorMessage, useNotificationService } from "../../imports";
import { useMutation } from "react-query";
import { AxiosError } from "axios";
import { createUseStyles } from "react-jss";
import { useNavigate } from "react-router-dom";

const useStyles = createUseStyles({
  submitButton: {
    float: 'right'
  }
});

type Props = {
  user?: Profile;
}

export const DeactivateProfile: FC<Props> = ({ user }) => {

  const { setUser, setAuthenticated } = useContext(AuthContext);
  const nav = useNavigate();

  const usersService = useUserService();
  const notificationService = useNotificationService();

  const deactivateProfileMutation = useMutation(() => usersService.disable(user?.ID ?? ''), {
    onSuccess: () => {
      notificationService.success('You have successfully disabled.');
      setUser(undefined);
      setAuthenticated(false);
      nav('/');
      localStorage.removeItem('jwt-token');
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const deactivateProfile = () => deactivateProfileMutation.mutate();
  
  const classes = useStyles();

  return (
    <>
      Are you sure that you want to disable your profile?
      <Button className={classes.submitButton} onClick={deactivateProfile} type="button">Submit</Button>
    </>
  );
}