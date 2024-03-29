import axios from "axios";
import {
  Service,
  USERS_SERVICE_URL
} from '../imports';
import { AuthResponse } from "../models/AuthResponse";

export const AUTH_SERVICE_ID = 'AuthService';

export class AuthService extends Service {
  constructor() {
    super(AUTH_SERVICE_ID, `${USERS_SERVICE_URL}/auth`)
  }

  public async login(email: string, password: string): Promise<AuthResponse> {
    return (await axios.post(`${this.baseUrl}`, { email, password })).data;
  }

}