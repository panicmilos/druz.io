package repository

import "gorm.io/gorm"

type Repository struct {
	DB         *gorm.DB
	Users      *UsersCollection
	LivePlaces *LivePlacesCollection
	WorkPlaces *WorkPlacesCollection
	Educations *EducationsCollection
	Intereses  *InteresesCollection

	PasswordRecoveriesCollection *PasswordRecoveriesCollection
}
