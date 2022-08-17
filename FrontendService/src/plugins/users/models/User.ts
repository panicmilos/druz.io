export type Profile = {
  ID: string,
  Role: string,
  FirstName: string,
  LastName: string,
  Birthday: string,
  Gender: Gender,

  About: string,
  PhoneNumber: string,
  LivePlaces: LivePlace[],
  WorkPlaces: WorkPlace[],
  Educations: Education[],
  Intereses: Interes[],

  Image: string
}

export enum Gender {
  Male,
  Gemale,
  Other
}

export type LivePlace = {
  id: string,
  Place: string,
  LivePlaceType: LivePlaceType
}

export enum LivePlaceType {
  Currently,
  Lived,
  Birthplace
}

export type WorkPlace = {
  id: string,
  Place: string,
  From: string,
  To: string
}

export type Education = {
  id: string,
  Place: string,
  From: string,
  To: string
}

export type Interes = {
  id: string,
	Interes: string
}
