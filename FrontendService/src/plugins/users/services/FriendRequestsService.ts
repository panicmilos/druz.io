import axios from "axios";
import { useEffect, useState } from "react";
import { USER_RELATIONS_SERVICE_URL } from "../imports";
import { FriendRequest } from "../models/FriendRequest";


export const useFriendRequestsService = (userId: string) => {

  const [friendRequestsService, setFriendRequestsService] = useState(new FriendRequestsService(+userId));

  useEffect(() => {
    setFriendRequestsService(new FriendRequestsService(+userId));
  }, [userId]);
  
  return friendRequestsService;
}

export class FriendRequestsService {

  public ID: string;
  private baseUrl: string;

  constructor(userId: number) {
    this.ID = "FriendRequestsService";
    this.baseUrl = `${USER_RELATIONS_SERVICE_URL}/users/${userId}/friends/requests`;
  }

  public async fetchSent(): Promise<FriendRequest[]> {
    return (await axios.get(`${this.baseUrl}/sent`)).data;
  }

  public async fetchReceived(): Promise<FriendRequest[]> {
    return (await axios.get(`${this.baseUrl}/received`)).data;
  }


  public async add(friendId: string): Promise<FriendRequest> {
    return (await axios.post(`${this.baseUrl}`, {FriendId: +friendId})).data;
  }

  public async accept(friendId: string): Promise<FriendRequest> {
    return (await axios.post(`${this.baseUrl}/accept`, {FriendId: +friendId})).data;
  }

  public async decline(friendId: string): Promise<FriendRequest> {
    return (await axios.delete(`${this.baseUrl}/decline`, { data: { FriendId: +friendId } })).data;
  }

}