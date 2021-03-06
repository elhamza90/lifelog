package db_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elhamza90/lifelog/internal/store/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo db.Repository
var grmDb *gorm.DB

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	var err error
	grmDb, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		os.Exit(1)
	}
	grmDb.AutoMigrate(&db.Tag{}, &db.Expense{}, &db.Activity{})
	repo = db.NewRepository(grmDb)
	log.Debug("Test Setup Complete")
	os.Exit(m.Run())
}

func clearDB() {
	grmDb.Where("1 = 1").Delete(&db.Expense{})
	grmDb.Where("1 = 1").Delete(&db.Activity{})
	grmDb.Where("1 = 1").Delete(&db.Tag{})
	defer grmDb.Exec("DELETE FROM expense_tags")
	defer grmDb.Exec("DELETE FROM activity_tags")
}
