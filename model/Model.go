package model

import (
	"gorm.io/gorm"
	"time"
)

// gorm defined default Model
type Model struct {
	ID uint				`gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt time.Time	`gorm:"index"`
}

type User struct {
	Name1 string `gorm:"<-:create"` // allow read and create
	Name2 string `gorm:"<-:update"` // allow read and update
	Name3 string `gorm:"<-"`        // allow read and write (create and update)
	Name4 string `gorm:"<-:false"`  // allow read, disable write permission
	Name5 string `gorm:"->"`        // readonly (disable write permission unless it configured )
	Name6 string `gorm:"->;<-:create"` // allow read and create
	Name7 string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	Name8 string `gorm:"-"`  // ignore this field when write and read
}

type UserTime struct {
	CreatedAt time.Time	// set to currentTime if it is zero on creating
	UpdateAt int 		// Set to current unix seconds on updaing or if it is zero on creating
	Updated int64 	`gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	Updated2 int64 	`gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	Created int64 	`gorm:"autoCreateTime"`		// Use unix seconds as creating time
}

// embedded struct

// User1 eqauls User2
type User1 struct {
	gorm.Model
	Name string
}

type User2 struct {
	ID uint						`gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt	`gorm:"index"`
	Name string
}


type Author struct {
	Name string
	Email string
}

// Blog1 equals Blog2
type Blog1 struct {
	ID	int
	Author Author	`gorm:"embedded"`
	Upvotes	int32
}

type Blog2 struct {
	ID	int64
	Name string
	Email string
	Upvotes	int32
}

// Blog3 eqauls Blog4
type Blog3 struct {
	ID	int64
	Author	Author	`gorm:"embedded;embeddedPrefix:author_"`
	Upvotes	int32
}

type Blog4 struct {
	ID	int64
	AuthorName	string
	AuthorEmail	string
	Upvotes		int32
}
