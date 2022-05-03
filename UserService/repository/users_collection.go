package repository

import (
	"UserService/dto"
	"UserService/models"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UsersCollection struct {
	DB *gorm.DB
}

func (userCollection *UsersCollection) Search(params *dto.UsersSearchParams) *[]models.Profile {
	users := &[]models.Profile{}

	query := userCollection.DB.Table("profiles")
	query.Where("(profiles.disabled is NULL OR profiles.disabled = 0)")

	if len(strings.TrimSpace(params.Name)) != 0 {
		query.Where("CONCAT(LOWER(profiles.first_name), ' ', LOWER(profiles.last_name)) like ?", "%"+strings.ToLower(params.Name)+"%")
	}

	if params.Gender != nil {
		query.Where("profiles.gender = ?", params.Gender)
	}

	if len(strings.TrimSpace(params.LivePlace)) != 0 {
		query.Joins("JOIN live_places lp ON profiles.id = lp.profile_id").Where("LOWER(lp.place) like ?", "%"+strings.ToLower(params.LivePlace)+"%")
	}

	if len(strings.TrimSpace(params.WorkPlace)) != 0 {
		query.Joins("JOIN work_places wp ON profiles.id = wp.profile_id").Where("LOWER(wp.place) like ?", "%"+strings.ToLower(params.WorkPlace)+"%")
	}

	if len(strings.TrimSpace(params.Education)) != 0 {
		query.Joins("JOIN educations e ON profiles.id = e.profile_id").Where("LOWER(e.place) like ?", "%"+strings.ToLower(params.Education)+"%")
	}

	if len(strings.TrimSpace(params.Interes)) != 0 {
		query.Joins("JOIN interes i ON profiles.id = i.profile_id").Where("LOWER(i.interes) like ?", "%"+strings.ToLower(params.Interes)+"%")
	}

	query.Distinct().Find(users)
	return users
}

func (userCollection *UsersCollection) ReadAccountByEmail(email string) *models.Account {
	account := &models.Account{}

	result := userCollection.DB.Preload("Profile").Where("email = ?", email).First(account)
	if result.RowsAffected == 0 || account.Profile.Disabled == true {
		return nil
	}

	return account
}

func (userCollection *UsersCollection) ReadAccountByProfileId(id uint) *models.Account {
	profile := &models.Profile{}

	result := userCollection.DB.First(profile, id)
	if result.RowsAffected == 0 || profile.Disabled == true {
		return nil
	}

	account := &models.Account{}
	result = userCollection.DB.Preload("Profile").First(account, profile.AccountID)
	if result.RowsAffected == 0 {
		return nil
	}

	return account
}

func (userCollection *UsersCollection) ReadById(id uint) *models.Profile {
	profile := &models.Profile{}

	result := userCollection.DB.Preload("LivePlaces").Preload("WorkPlaces").Preload("Education").Preload("Intereses").First(profile, id)
	if result.RowsAffected == 0 || profile.Disabled == true {
		return nil
	}

	return profile
}

func (userCollection *UsersCollection) ReadDeactivatedByEmail(email string) *models.Profile {
	account := &models.Account{}
	result := userCollection.DB.Preload("Profile").Where("email = ?", email).First(account)
	if result.RowsAffected == 0 || account.Profile.Disabled == false {
		return nil
	}

	return &account.Profile
}

func (userCollection *UsersCollection) ReadDeactivatedById(id uint) *models.Profile {
	profile := &models.Profile{}
	result := userCollection.DB.Preload("LivePlaces").Preload("WorkPlaces").Preload("Education").Preload("Intereses").First(profile, id)
	if result.RowsAffected == 0 || profile.Disabled == false {
		return nil
	}

	return profile
}

func (userCollection *UsersCollection) Create(user *models.Account) *models.Profile {
	userCollection.DB.Create(user)

	return &user.Profile
}

func (userCollection *UsersCollection) UpdateAccount(account *models.Account) *models.Account {
	userCollection.DB.Save(account)

	return account
}

func (userCollection *UsersCollection) UpdateProfile(profile *models.Profile) *models.Profile {
	userCollection.DB.Save(profile)

	return profile
}

func (userCollection *UsersCollection) Delete(id uint) *models.Profile {
	account := userCollection.ReadAccountByProfileId(id)

	userCollection.DB.Select(clause.Associations).Delete(account)

	return &account.Profile
}
