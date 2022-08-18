import axios from "axios";
import { useEffect, useState } from "react";
import { CHAT_SERVICE_URL } from "../imports";
import { Message } from "../models/Message";

export const useMessageService = (userId: string) => {

  const [messageService, setMessageService] = useState(new MessageService(userId));

  useEffect(() => {
    setMessageService(new MessageService(userId));
  }, [userId]);

  return messageService;
}

export class MessageService {

  public ID: string;
  private baseUrl: string;

  constructor(userId: string) {
    this.ID = `MessageService`;
    this.baseUrl = `${CHAT_SERVICE_URL}/users/${userId}/message`
  }

  public async message(message: string): Promise<Message> {
    return (await axios.post(this.baseUrl, { message, type: 'Text'})).data;
  }
}