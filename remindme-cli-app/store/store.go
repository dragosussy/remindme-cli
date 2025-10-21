package store

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"remind.me/common/model"
)

const (
	// fileModeOwnerRW grants read/write to the owner only (octal 0600)
	fileModeOwnerRW = 0o600

	INSERT_REMINDER_QUERY = `
		INSERT INTO reminders (id, title, text, cron_expression, next_run_at, acknowledged)
		VALUES (?, ?, ?, ?, ?, ?)
	`
)

type Store struct {
	db *sql.DB
}

func dbFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/.remindme-db", nil
}

func NewStore() (*Store, error) {
	path, err := dbFilePath()
	if err != nil {
		return nil, err
	}

	dbFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, fileModeOwnerRW)
	if err != nil {
		return nil, err
	}
	_ = dbFile.Close()

	dsn := path + "?_journal_mode=WAL&_busy_timeout=5000"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(2)

	return &Store{db: db}, nil
}

func (s *Store) Save(r *model.Reminder) error {
	_, err := s.db.Exec(
		INSERT_REMINDER_QUERY,
		r.Id,
		r.Title,
		r.Text,
		r.CronExpression,
		r.NextRunAt,
		r.Acknowledged,
	)
	return err
}
