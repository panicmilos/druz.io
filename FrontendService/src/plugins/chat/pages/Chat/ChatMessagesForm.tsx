import { AxiosError } from "axios";
import { FC, useState } from "react"
import { useMutation } from "react-query";
import { Button, extractErrorMessage, Form, TextAreaInput, useNotificationService } from "../../imports";
import { Message } from "../../models/Message";
import { useMessageService } from "../../services";

type Props = {
  friendId: string,
  onSendCallback: (message: Message) => any
}

export const ChatMessagesForm: FC<Props> = ({ friendId, onSendCallback }) => {

  const messageService = useMessageService(friendId);
  const notificationService = useNotificationService();
  const [message, setMessage] = useState('');

  const sendMessageMutation = useMutation((message: string) => messageService.message(message), {
    onSuccess: onSendCallback,
    onError: (error: AxiosError) => notificationService.error(extractErrorMessage(error?.response?.data))
  });

  const onSubmit = (message: string) => {
    if (!message) return;

    sendMessage(message);
    setMessage('');
  };
  
  const sendMessage = (message: string) => sendMessageMutation.mutate(message);
  
  return (
    <Form
      schema={undefined}
      onSubmit={() => onSubmit(message)}
    >
      <TextAreaInput label="Message" value={message} onChange={setMessage} />

      <Button type="submit">Send</Button>
    </Form>
  );

}