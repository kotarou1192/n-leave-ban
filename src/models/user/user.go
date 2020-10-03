package user

import (
	"database/sql"
	"errors"

	"../../utils/dbloader"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	LeaveCount sql.NullInt64  `gorm:"not null"`
	DiscordID  sql.NullString `gorm:"unique;not null"`
}

func (user *User) GetLeaveCount() (int, error) {
	if user.LeaveCount.Valid {
		return int(user.LeaveCount.Int64), nil
	}
	return -1, errors.New("user is nil")
}

func (user *User) Set(leaveCount int64, discordID string) {
	user.LeaveCount = sql.NullInt64{Int64: leaveCount, Valid: true}
	user.DiscordID = sql.NullString{String: discordID, Valid: true}
}

func (user *User) SetNull() {
	user.LeaveCount = sql.NullInt64{Int64: 0, Valid: false}
	user.DiscordID = sql.NullString{String: "", Valid: false}
}

func New(discordID string) *User {
	db := dbloader.InitDB()
	defer db.Close()
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		panic(err.Error())
	}
	user := User{}
	user.Set(0, discordID)
	return &user
}

func (user *User) Save() (*User, error) {
	db := dbloader.InitDB()
	defer db.Close()
	if !db.HasTable(user) {
		db.Debug().CreateTable(user)
	}
	if err := db.Debug().Create(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) AddLeaveTime() error {
	db := dbloader.InitDB()
	defer db.Close()
	user.tableChecker(db)
	user.leaveCountUp()
	if !user.DiscordID.Valid {
		return errors.New("user not found")
	}
	db.Debug().Model(user).Update("LeaveCount", user.LeaveCount.Int64)
	return nil
}

// 見つからなかったらnullUserを返す
func Find(discordID string) (*User, error) {
	db := dbloader.InitDB()
	defer db.Close()
	user := User{}
	user.SetNull()
	db.Debug().Where("discord_id = ?", discordID).First(&user)
	if !user.DiscordID.Valid {
		return &user, errors.New("user not found")
	}
	return &user, nil
}

func (user *User) leaveCountUp() {
	user.LeaveCount.Int64 = user.LeaveCount.Int64 + 1
}

func (user *User) tableChecker(db *gorm.DB) {
	if !db.Debug().HasTable(user) {
		panic(errors.New("Record is not saved"))
	}
}
