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

  public error(text: string) {
    toast.error(text);
  }
}