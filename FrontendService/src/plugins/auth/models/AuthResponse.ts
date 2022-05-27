import { User } from "./User";

export type AuthResponse = {
  user: User;
  token: Token;
}

export type Token = {
  value: string
}