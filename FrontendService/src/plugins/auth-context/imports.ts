import { CrudService } from "../../core";
import { USERS_SERVICE_URL } from "../../urls";

export type { ContextPlugin } from "../../core";

export type Profile = {
  ID: string,
  Role: string,
  FirstName: string,
  LastName: string,
  Birthday: string,
  Gender: any,

  About: string,
  PhoneNumber: string,
  LivePlaces: any[],
  WorkPlaces: any[],
  Educations: any[],
  Intereses: any[],

  Image: string
};

export class UsersService extends CrudService<Profile> {
  constructor() {
    super("UsersService", `${USERS_SERVICE_URL}/users`);
  }
}
