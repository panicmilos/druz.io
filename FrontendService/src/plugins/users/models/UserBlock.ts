import { Profile } from "./User"

export type UserBlock = {
  ID: string,

  BlockedById: number,

  BlockedId: number,
  Blocked: Profile
}