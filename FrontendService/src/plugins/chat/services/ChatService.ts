import axios from "axios";
import { useState } from "react";
import { CHAT_SERVICE_URL } from "../imports";
import { Chat } from "../models/Chat";
import { Message } from "../models/Message";
import { Status } from "../models/Statuts";

export const useChatService = () => {

  const [chatService] = useState(new ChatService());

  return chatService;
}

export class ChatService {

  public ID: string;
  private baseUrl: string;

  constructor() {
    this.ID = `ChatService`;
    this.baseUrl = `${CHAT_SERVICE_URL}/users/chats`;
  }

  public async fetch(): Promise<Chat[]> {
    return (await axios.get(this.baseUrl)).data;
  }

  public async fetchStatuses(): Promise<Status[]> {
    return (await axios.get(`${this.baseUrl}/statuses`)).data;
  }

  public async fetchById(chat: string, keywoard: string): Promise<Message[]> {
    return (await axios.get(`${this.baseUrl}/${chat}`, { params: { keywoard }})).data;
  }

  public async deleteMessage(chat: string, messageId: string, mode: string = 'for_me'): Promise<Message> {
    return (await axios.delete(`${this.baseUrl}/${chat}/${messageId}`, { params: { mode }})).data;
  }

  public async delete(chat: string, mode: string = 'for_me'): Promise<Message[]> {
    return (await axios.delete(`${this.baseUrl}/${chat}`, { params: { mode }})).data;
  }

}