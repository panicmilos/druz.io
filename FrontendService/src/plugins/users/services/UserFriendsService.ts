import axios from "axios";
import { useEffect, useState } from "react";
import { USER_RELATIONS_SERVICE_URL } from "../imports";
import { UserFriend } from "../models/UserFriend";


export const useUserFriendsService = (userId: string) => {

  const [userFriendsService, setUserFriendsService] = useState(new UserFriendsService(+userId));

  useEffect(() => {
    setUserFriendsService(new UserFriendsService(+userId));
  }, [userId]);
  
  return userFriendsService;
}

export class UserFriendsService {

  public ID: string;
  private baseUrl: string;

  constructor(userId: number) {
    this.ID = "UserFriendsService";
    this.baseUrl = `${USER_RELATIONS_SERVICE_URL}/users/${userId}/friends`;
  }

  public async fetch(): Promise<UserFriend[]> {
    return (await axios.get(`${this.baseUrl}`)).data;
  }

  public async fetchById(id: string): Promise<UserFriend> {
    return (await axios.get(`${this.baseUrl}/${id}`)).data;
  }

  public async delete(friendId: string): Promise<UserFriend> {
    return (await axios.delete(`${this.baseUrl}`, { data: { FriendId: +friendId } })).data;
  }

}