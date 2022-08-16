import axios from "axios";
import { useState } from "react";
import { USERS_SERVICE_URL } from "../imports";

export const usePasswordRecoveryService = () => {
  const [passwordRecoveryService] = useState(new PasswordRecoveryService());

  return passwordRecoveryService;
}

export class PasswordRecoveryService {

  public ID: string;
  private baseUrl: string;

  constructor() {
    this.ID = "UserReactivationService";
    this.baseUrl = `${USERS_SERVICE_URL}/users`;
  }

  public async request(email: string): Promise<any> {
    return (await axios.post(`${this.baseUrl}/password/recover/request`, {email})).data;
  }

  public async recover(id: string, token: string, newPassword: string): Promise<any> {
    return (await axios.put(`${this.baseUrl}/${id}/password/recover?token=${token}`, {newPassword})).data;
  }


}