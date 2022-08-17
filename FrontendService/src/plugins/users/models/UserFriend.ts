import { Profile } from "./User"

export type UserFriend = {

  ID: string,

  UserId: number,

  FriendId: number,
  Friend: Profile
}