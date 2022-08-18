import { useEffect } from "react";
import { useNavigate, useParams, useSearchParams } from "react-router-dom";
import { extractErrorMessage, useNotificationService } from "../../imports";
import { useUserReactivationService } from "../../services";


export const Reactivation = () => {

  const nav = useNavigate();
  const { id } = useParams();
  let [searchParams] = useSearchParams();
  const token = searchParams.get("token");

  const userReactivationService = useUserReactivationService();
  const notificationService = useNotificationService();

  useEffect(() => {

    if (!token || !id) {
      return;
    }

    userReactivationService.reactivate(id, token)
      .then(_ => {
        notificationService.success('You have succesfully reactivated your profile.');
        nav('/');
      })
      .catch(error => {
        notificationService.error(extractErrorMessage(error.response?.data));
        nav('/');
      })

  }, []);

  return (
    <>
    </>
  );
}