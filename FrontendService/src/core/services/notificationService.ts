import { useState } from "react";
import { toast } from "react-toastify";

export const useNotificationService = () => {

  const [notificationService] = useState(new NotificationService());

  return notificationService;
}

export class NotificationService {

  public success(text: string) {
    toast.success(text);
  }

  public info(text: string) {
    toast.info(text, { autoClose: 500, position: 'bottom-right' });
  }

  public error(text: string) {
    toast.error(text);
  }
}