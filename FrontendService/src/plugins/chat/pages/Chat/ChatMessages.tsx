import { FC, useEffect, useState } from "react";
import { useChatService, useMessageService } from "../../services";
import * as Yup from 'yup';
import { Button, Form, DropdownMenu, TextAreaInput, DropdownItem, useNotificationService, extractErrorMessage } from "../../imports";
import { useMutation, useQuery } from "react-query";
import { Message } from "../../models/Message";
import { ChatMessagesForm } from "./ChatMessagesForm";
import { ChatMessagesList } from "./ChatMessagesList";
import { AxiosError } from "axios";

type Props = {
  chat: any,
  onInitial: () => any
}


export const ChatMessages: FC<Props> = ({ chat, onInitial}) => {

  const { chatId, friendId } = chat;

  const chatService = useChatService();
  const notificationService = useNotificationService();

  const [messages, setMessages] = useState<Message[]>([]);


  useQuery([chat, chatService], () => chatService.fetchById(chat.chatId || ''), {
    enabled: chat.chatId !== 'NOT_CREATED_YET',
    onSuccess: (messages: Message[]) => setMessages(messages || [])
  })

  const OnSendMessage = (message: Message) => { 
    if (chat.chatId !== 'NOT_CREATED_YET') {
      setMessages([...messages, message]);
      var element = document.getElementById("scrollable");
      element?.scrollTop && (element.scrollTop = element?.scrollHeight);
    } else {
      onInitial();
    }
  }

  const OnDeleteMessage = (message: Message) => { 
    setMessages(messages.filter(m => m.ID !== message.ID));
  }


  const deleteChatMutator = useMutation((mode: string) => chatService.delete(chatId, mode), {
    onSuccess: () => {
      notificationService.success("You have succesfully deleted chat.");
      onInitial();
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
    }
  });
  const deleteChat = (mode: string) => deleteChatMutator.mutate(mode);


  return (
    <>
    <div style={{ display: 'flex', justifyContent: 'flex-end', margin: '0 1em 1em 0'}}>
      <Button onClick={() => deleteChat('for_me')}>Delete For Me</Button>
      <Button onClick={() => deleteChat('for_both')}>Delete For Both</Button>
    </div>

    <div id="scrollable" style={{ maxHeight: '600px', minHeight: '600px', 'overflowY': 'scroll', display: 'flex', flexDirection: 'column'}}>
      <ChatMessagesList chatId={chatId} messages={messages} onDeleteCallback={OnDeleteMessage} />
    </div>

    <ChatMessagesForm friendId={friendId} onSendCallback={OnSendMessage} />

    </>
  );
}