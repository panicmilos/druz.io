import { CrudService } from "../../core";
import { USERS_SERVICE_URL } from "../../urls";

export type { ContextPlugin } from "../../core";

export type Profile = {
  id: string;
  fullName: string;
  email: string;
  phoneNumber: string;
  role: string;
};

export class UsersService extends CrudService<Profile> {
  constructor() {
    super("UsersService", `${USERS_SERVICE_URL}/users`);
  }
}
