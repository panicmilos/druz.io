import { Profile } from "./User"

export type FriendRequest = {

  ID: string,

  UserId: number,
  User: Profile,

  FriendId: number,
  Friend: Profile
}