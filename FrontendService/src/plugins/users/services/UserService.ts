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

  public async search(params: any): Promise<Profile[]> {
    return (await axios.get( `${this.baseUrl}/search`, {params})).data;
  }

  public async changeImage(userId: string, image: any): Promise<Profile> {
    return (await axios.put( `${this.baseUrl}/${userId}/image`, { image })).data;
  }

  public async changePassword(userId: string, request: any): Promise<Profile> {
    return (await axios.put( `${this.baseUrl}/${userId}/password`, request)).data;
  }

  public async disable(userId: string): Promise<Profile> {
    return (await axios.put( `${this.baseUrl}/${userId}/disable`)).data;
  }

  public async block(userId: string): Promise<Profile> {
    return (await axios.delete( `${this.baseUrl}/${userId}/block`)).data;
  }

  public async report(userId: string, reason: string): Promise<any> {
    return (await axios.post( `${this.baseUrl}/${userId}/report`, {reason})).data;
  }


}