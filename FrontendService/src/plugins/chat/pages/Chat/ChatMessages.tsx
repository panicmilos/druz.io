import { FC, useContext, useEffect, useState } from "react";
import { useChatService } from "../../services";
import { Button, useNotificationService, extractErrorMessage, SocketContext, useDebounce, TextInput } from "../../imports";
import { useMutation, useQuery } from "react-query";
import { Message } from "../../models/Message";
import { ChatMessagesForm } from "./ChatMessagesForm";
import { ChatMessagesList } from "./ChatMessagesList";
import { AxiosError } from "axios";

type Props = {
  chat: any,
  onInitial: () => any
}

var globalMessages: Message[] = [];

const srollToBottom = (timeout = 20) => {
  setTimeout(() => {
    var element = document.getElementById("scrollable");
    element?.scrollTop && (element.scrollTop = element?.scrollHeight);
  }, timeout);
}

export const ChatMessages: FC<Props> = ({ chat, onInitial }) => {

  const { chatId, friendId } = chat;

  const chatService = useChatService();
  const notificationService = useNotificationService();

  const [messages, setMessages] = useState<Message[]>([]);
  globalMessages = messages;

  const [searchText, setSearchText] = useState('');
  const debouncedSearchText = useDebounce(searchText, 300);

  useQuery([chat, chatService, debouncedSearchText], () => chatService.fetchById(chat.chatId || '', debouncedSearchText), {
    enabled: chat.chatId !== 'NOT_CREATED_YET',
    onSuccess: (messages: Message[]) => { setMessages(messages?.sort((m1: Message, m2: Message) => m1.CreatedAt.localeCompare(m2.CreatedAt)) || []); }
  })


  const { client } = useContext(SocketContext);
  useEffect(() => {
    if (!client) return;

    client?.on('messages_chat', function(data: any) {
      const message = JSON.parse(data.text).Message as Message;
      if (!message.ID.includes(`/${chatId}/`)) return;
      setMessages([...globalMessages, message]);
      srollToBottom();
    });

    client?.on('messages_delete', function(data: any) {
      const message = JSON.parse(data.text).Message as Message;
      if (!message.ID.includes(`/${chatId}/`)) return;

      setMessages([...globalMessages.filter(m => m.ID !== message.ID)]);
      srollToBottom();
    });

    client?.on('chat_delete', function(data: any) {
      const chat = JSON.parse(data.text);

      if (chat.ChatId !== chatId) return;

      setMessages([]);
      srollToBottom();
    });

    return () => { client.removeAllListeners('messages_chat'); client.removeAllListeners('messages_delete'); }
  }, [client]);
  



  const OnSendMessage = (message: Message) => { 
    if (chat.chatId !== 'NOT_CREATED_YET') {
      setMessages([...messages, message]);
      srollToBottom();
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

    <div style={{ display: 'flex', justifyContent: 'flex-end', margin: '0 1em 1em 0'}}>
      <TextInput value={searchText} onChange={setSearchText} placeholder="Keyword..." />
    </div>

    <div id="scrollable" style={{ maxHeight: '600px', minHeight: '600px', 'overflowY': 'scroll', display: 'flex', flexDirection: 'column'}}>
      <ChatMessagesList chatId={chatId} messages={messages} onDeleteCallback={OnDeleteMessage} />
    </div>

    <ChatMessagesForm friendId={friendId} onSendCallback={OnSendMessage} />

    </>
  );
}