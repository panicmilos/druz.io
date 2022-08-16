import axios from "axios";
import { useState } from "react";
import { CrudService, USERS_SERVICE_URL } from "../imports";
import { Profile } from "../models/User";


export const useUserService = () => {
  const [userService] = useState(new UserService());

  return userService;
}

export class UserService extends CrudService<Profile> {
  
  constructor() {
    super("UsersService", `${USERS_SERVICE_URL}/users`);
  }

  public async changePassword(userId: string, request: any): Promise<Profile> {
    console.log(`${this.baseUrl}/${userId}/password`);
    return (await axios.put( `${this.baseUrl}/${userId}/password`, request)).data;
  }


}