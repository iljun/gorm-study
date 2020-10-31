package lifeCycle

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name 		string
	Age 		uint
	BirthDay 	time.Time
	Confirmed	bool
	Invalid 	bool
}

// Create Hook

// begin transaction
// BeforeSave
// BeforeCreate
// save before associations
// insert into database
// save after associations
// AfterCreate
// AfterSave
// commit or rollback transaction

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID > 0 {
		err = errors.New("can't save invalid data")
	}

	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", "admin")
	}

	return
}

// Update Hook

// begin transaction
// BeforeSave
// BeforeUpdate
// save before associations
// update database
// save after associations
// AfterUpdate
// AfterSave
// commit or rollback transaction

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !u.Invalid {
		err = errors.New("invalid user")
	}
	return
}

// Updating data in same transaction
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	if u.Confirmed {
		tx.Model(&User{}).Where("user_id = ?", u.ID).Update("confirmed", true)
	}
	return
}

// Delete Hook

// begin transaction
// BeforeDelete
// delete from database
// AfterDelete
// commit or rollback transaction

// Updating data in same transaction
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	if u.Confirmed {
		tx.Model(&User{}).Where("user_id = ?", u.ID).Update("invalid", false)
	}
	return
}

// Query Object Hook

// load data from database
// Preloading (eager loading)
// AfterFind
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Name == "" {
		u.Name = "default Name"
	}
	return
}
