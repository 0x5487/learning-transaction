package main

import (
	"context"
	"os"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jasonsoft/learning-transaction/internal/pkg/config"
	internalDatabase "github.com/jasonsoft/learning-transaction/internal/pkg/database"
	"github.com/jasonsoft/log/v2"

	"gorm.io/gorm"
)

var (
	_db *gorm.DB

	_cfg config.Configuration

	_eventRepo *EventRepo
)

func TestMain(m *testing.M) {
	var err error
	_cfg = config.New("app.yml")

	_cfg.InitLogger("tx")
	defer log.Flush()

	// initial database
	_db, err = _cfg.InitDatabase("starter")
	if err != nil {
		panic(err)
	}

	_eventRepo = NewEventRepository(_db)

	exitVal := m.Run()

	os.Exit(exitVal)

}

func TestRepeatableReadTX(t *testing.T) {
	ctx := context.Background()
	logger := log.FromContext(ctx)

	// clear database data
	err := internalDatabase.RunSQLScripts(_db, _cfg.Path("test", "database", "starter_db"))
	if err != nil {
		panic(err)
	}

	go func() {
		//sess1 := _db.Begin()
		err2 := _db.Transaction(func(tx *gorm.DB) error {
			var err1 error

			_, err1 = _eventRepo.Event(ctx, 1, true, tx)
			if err != nil {
				return err1
			}
			logger.Debug("event1 read")

			time.Sleep(2 * time.Second)

			err1 = _eventRepo.UpdateTitle(ctx, 1, "event1", tx)
			if err != nil {
				return err1
			}
			logger.Debug("event1 update")

			// return nil will commit the whole transaction
			return nil
		})

		if err2 != nil {
			panic(err2)
		}
		//logger.Debug("event1 done")
	}()

	go func() {
		//sess2 := _db.Begin()
		err4 := _db.Transaction(func(tx *gorm.DB) error {
			var err3 error
			//logger.Debug("event2")
			time.Sleep(1 * time.Second)

			_, err3 = _eventRepo.Event(ctx, 1, false, tx)
			if err != nil {
				return err
			}
			logger.Debug("event2 read")

			err3 = _eventRepo.UpdateTitle(ctx, 1, "event2", tx)
			if err3 != nil {
				return err3
			}
			logger.Debug("event2 update")

			return nil
		})

		if err4 != nil {
			panic(err4)
		}
		//sess2.Commit()
		//logger.Debug("event2 done")
	}()

	time.Sleep(5 * time.Second)

}
