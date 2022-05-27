import { CrudService } from "../../core";
import { USERS_SERVICE_URL } from "../../urls";

export type { ContextPlugin } from "../../core";
export type User = {
  id: string;
  fullName: string;
  email: string;
  phoneNumber: string;
  groupId: string;
};
export type Role = {
  id: string;
  name: string;
  permissions: string[];
};
export class UsersService extends CrudService<User> {
  constructor(groupId: string) {
    super("UsersService", `${USERS_SERVICE_URL}/groups/${groupId}/users`);
  }
}
