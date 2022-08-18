import { FC, useState } from "react"
import { useMutation } from "react-query";
import { Button, Form, TextAreaInput } from "../../imports";
import { Message } from "../../models/Message";
import { useMessageService } from "../../services";

type Props = {
  friendId: string,
  onSendCallback: (message: Message) => any
}

export const ChatMessagesForm: FC<Props> = ({ friendId, onSendCallback }) => {

  const messageService = useMessageService(friendId);
  const [message, setMessage] = useState('');

  const sendMessageMutation = useMutation((message: string) => messageService.message(message), {
    onSuccess: onSendCallback
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