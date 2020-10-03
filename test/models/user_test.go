package models

import (
	"testing"

	"../../src/models/user"
	"../../src/utils/dbloader"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func beforeTest(saved *user.User) {
	db := dbloader.InitDB()
	db.Unscoped().Delete(saved)
	db.Close()
}

func TestCreateUser(t *testing.T) {
	saved, _ := user.New("hogefuga").Save()
	found, err := user.Find("hogefuga")
	if err != nil {
		t.Errorf("user not found")
	}
	if found.ID != saved.ID {
		t.Errorf("ID should be " + string(found.ID) + " same to " + string(saved.ID))
	}
	beforeTest(saved)
}

func TestIncrementLeaveCount(t *testing.T) {
	saved, _ := user.New("takashiii").Save()
	saved2, _ := user.New("tanaka").Save()

	found1, err1 := user.Find("takashiii")
	found2, err2 := user.Find("tanaka")

	if err1 != nil || err2 != nil {
		t.Errorf("not found error")
	}

	for i := 0; i < 3; i++ {
		if err := found1.AddLeaveTime(); err != nil {
			t.Errorf("increment error")
		}
	}
	if err := found2.AddLeaveTime(); err != nil {
		t.Errorf("increment error")
	}

	if val, _ := found1.GetLeaveCount(); val != 3 {
		t.Errorf("LeaveCount should be 3")
	}
	if val, _ := found2.GetLeaveCount(); val != 1 {
		t.Errorf("Leave count should be 1")
	}
	beforeTest(saved)
	beforeTest(saved2)
}

func TestUpdateAllBug(t *testing.T) {
	a, _ := user.New("takahashi").Save()
	b, _ := user.New("yoshida").Save()

	found1, _ := user.Find("takahashi")
	found1.AddLeaveTime()
	found2, _ := user.Find("yoshida")
	notFound, _ := user.Find("notFound")

	notFound.AddLeaveTime()

	if num, _ := found1.GetLeaveCount(); num != 1 {
		t.Errorf("update all bug. dangerous!!!")
	}
	if num, _ := found2.GetLeaveCount(); num != 0 {
		t.Errorf("update all bug. dangerous!!!")
	}
	beforeTest(a)
	beforeTest(b)
}
