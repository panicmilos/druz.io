
import type { Profile } from './User'

export type Report = {

  ID: string,
  
  ReportedId: number,
  Reported: Profile,

  ReportedById: number,
  ReportedBy: Profile,

  Reason: string
}