import axios from "axios";
import { useEffect, useState } from "react";
import { USER_RELATIONS_SERVICE_URL } from "../imports";
import { UserBlock } from "../models/UserBlock";

export const useUserBlocksService = (userId: string) => {

  const [userBlocksService, setUserBlocksService] = useState(new UserBlocksService(+userId));

  useEffect(() => {
    setUserBlocksService(new UserBlocksService(+userId));
  }, [userId]);
  
  return userBlocksService;
}

export class UserBlocksService {

  public ID: string;
  private baseUrl: string;

  constructor(userId: number) {
    this.ID = "UserReportsService";
    this.baseUrl = `${USER_RELATIONS_SERVICE_URL}/users/${userId}/block-list`;
  }

  public async fetch(): Promise<UserBlock[]> {
    return (await axios.get(`${this.baseUrl}`)).data;
  }

  public async add(blockedId: string): Promise<UserBlock> {
    return (await axios.post(`${this.baseUrl}`, { blockedId: +blockedId })).data;
  }

  public async delete(blockedId: string): Promise<UserBlock> {
    return (await axios.delete(`${this.baseUrl}`, { data: { blockedId: +blockedId } })).data;
  }

}