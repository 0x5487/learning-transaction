package main

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (repo *EventRepo) Event(ctx context.Context, eventID int64, forUpdate bool, tx ...*gorm.DB) (Event, error) {
	//logger := log.FromContext(ctx)
	db := repo.db
	if tx != nil {
		db = tx[0]
	}

	event := Event{}

	if forUpdate {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	err := db.First(&event, eventID).Error
	if err != nil {
		return event, err
	}

	// sql := db.Statement.SQL
	// logger.Debugf("sql: %s", sql.String())

	return event, nil
}

func (repo *EventRepo) UpdateTitle(ctx context.Context, eventID int64, title string, tx ...*gorm.DB) error {
	db := repo.db
	if tx != nil {
		db = tx[0]
	}

	err := db.Model(Event{}).
		Where("id = ?", eventID).
		UpdateColumn("title", title).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found la")
		}
		return err
	}

	return nil
}
