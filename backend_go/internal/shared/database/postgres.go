package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgres(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}

func PrepareSchema(db *gorm.DB, models ...any) error {
	_ = models

	statements := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
		`CREATE TABLE IF NOT EXISTS users (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name text NOT NULL,
			email text UNIQUE NOT NULL,
			password text NOT NULL,
			created_at timestamp DEFAULT current_timestamp
		);`,
		`CREATE TABLE IF NOT EXISTS projects (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name text NOT NULL,
			description text,
			owner_id uuid REFERENCES users(id) ON DELETE CASCADE,
			created_at timestamp DEFAULT current_timestamp
		);`,
		`CREATE TABLE IF NOT EXISTS tasks (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			title text NOT NULL,
			description text,
			status text DEFAULT 'todo',
			priority text DEFAULT 'medium',
			project_id uuid REFERENCES projects(id) ON DELETE CASCADE,
			created_by_id uuid REFERENCES users(id) ON DELETE CASCADE,
			assignee_id uuid REFERENCES users(id) ON DELETE SET NULL,
			due_date date,
			created_at timestamp DEFAULT current_timestamp,
			updated_at timestamp DEFAULT current_timestamp
		);`,
		`ALTER TABLE tasks
		ADD COLUMN IF NOT EXISTS created_by_id uuid REFERENCES users(id) ON DELETE CASCADE;`,
		`DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'tasks_status_check'
			) THEN
				ALTER TABLE tasks
				ADD CONSTRAINT tasks_status_check CHECK (status IN ('todo', 'in_progress', 'done'));
			END IF;
		END$$;`,
		`DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'tasks_priority_check'
			) THEN
				ALTER TABLE tasks
				ADD CONSTRAINT tasks_priority_check CHECK (priority IN ('low', 'medium', 'high'));
			END IF;
		END$$;`,
	}

	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			return err
		}
	}

	return nil
}
