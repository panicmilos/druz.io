import axios from "axios";
import { useState } from "react";
import { USERS_SERVICE_URL } from "../imports";
import { Report } from "../models/Report";

export const useUserReportsService = () => {

  const [userReportsService] = useState(new UserReportsService());

  return userReportsService;
}

export class UserReportsService {

  public ID: string;
  private baseUrl: string;

  constructor() {
    this.ID = "UserReportsService";
    this.baseUrl = `${USERS_SERVICE_URL}/reports`;
  }

  public async search(params: any): Promise<Report[]> {
    return (await axios.get(`${this.baseUrl}/search`, { params })).data;
  }

  public async ignore(id: string): Promise<Report> {
    return (await axios.delete(`${this.baseUrl}/${id}/ignore`)).data;
  }

}