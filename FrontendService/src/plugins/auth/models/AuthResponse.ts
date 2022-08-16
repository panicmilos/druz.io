import { Profile } from "./User";

export type AuthResponse = {
  Profile: Profile;
  Jwt: string;
}