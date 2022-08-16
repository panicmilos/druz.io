import axios from "axios";
import { useState } from "react";
import { USERS_SERVICE_URL } from "../imports";

export const useUserReactivationService = () => {
  const [userReactivationService] = useState(new UserReactivationService());

  return userReactivationService;
}

export class UserReactivationService {

  public ID: string;
  private baseUrl: string;

  constructor() {
    this.ID = "UserReactivationService";
    this.baseUrl = `${USERS_SERVICE_URL}/users`;
  }

  public async request(email: string): Promise<any> {
    return (await axios.post(`${this.baseUrl}/reactivation/request`, {email})).data;
  }

  public async reactivate(id: string, token: string): Promise<any> {
    return (await axios.put(`${this.baseUrl}/${id}/reactivation?token=${token}`)).data;
  }

}