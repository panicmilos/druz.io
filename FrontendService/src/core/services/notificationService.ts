import { toast } from "react-toastify";

export class NotificationService {

  public success(text: string) {
    toast.success(text);
  }

  public error(text: string) {
    toast.error(text);
  }
}