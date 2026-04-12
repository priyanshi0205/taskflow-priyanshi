-- Seed credentials:
-- email: seed@example.com
-- password: password

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name text NOT NULL,
  email text UNIQUE NOT NULL,
  password text NOT NULL,
  created_at timestamp DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS projects (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name text NOT NULL,
  description text,
  owner_id uuid REFERENCES users(id) ON DELETE CASCADE,
  created_at timestamp DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS tasks (
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
);

ALTER TABLE tasks
ADD COLUMN IF NOT EXISTS created_by_id uuid REFERENCES users(id) ON DELETE CASCADE;

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint WHERE conname = 'tasks_status_check'
  ) THEN
    ALTER TABLE tasks
    ADD CONSTRAINT tasks_status_check CHECK (status IN ('todo', 'in_progress', 'done'));
  END IF;
END$$;

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint WHERE conname = 'tasks_priority_check'
  ) THEN
    ALTER TABLE tasks
    ADD CONSTRAINT tasks_priority_check CHECK (priority IN ('low', 'medium', 'high'));
  END IF;
END$$;

INSERT INTO users (id, name, email, password, created_at)
VALUES (
  '11111111-1111-1111-1111-111111111111',
  'Test User',
  'test@example.com',
  '$2b$10$wH8V2dY0V6F9n5HcH3JcG.6QwWl1vF5mZ2q8K9GvYp2sJ6WbYcZ9K',
  NOW()
)
ON CONFLICT (email) DO UPDATE
SET name = EXCLUDED.name;

INSERT INTO projects (id, name, description, owner_id, created_at)
VALUES (
  '22222222-2222-2222-2222-222222222222',
  'Greening Pilot Project',
  'Seeded project for local API/UI testing',
  '11111111-1111-1111-1111-111111111111',
  NOW()
)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name,
    description = EXCLUDED.description,
    owner_id = EXCLUDED.owner_id;

INSERT INTO tasks (
  id,
  title,
  description,
  status,
  priority,
  project_id,
  created_by_id,
  assignee_id,
  due_date,
  created_at,
  updated_at
)
VALUES
(
  '33333333-3333-3333-3333-333333333331',
  'Collect site baseline data',
  'Capture existing vegetation and soil metrics.',
  'todo',
  'high',
  '22222222-2222-2222-2222-222222222222',
  '11111111-1111-1111-1111-111111111111',
  '11111111-1111-1111-1111-111111111111',
  CURRENT_DATE + INTERVAL '7 days',
  NOW(),
  NOW()
),
(
  '33333333-3333-3333-3333-333333333332',
  'Prepare stakeholder workshop',
  'Coordinate local authorities and community groups.',
  'in_progress',
  'medium',
  '22222222-2222-2222-2222-222222222222',
  '11111111-1111-1111-1111-111111111111',
  NULL,
  CURRENT_DATE + INTERVAL '14 days',
  NOW(),
  NOW()
),
(
  '33333333-3333-3333-3333-333333333333',
  'Finalize plantation timeline',
  'Publish final timeline with milestones and owners.',
  'done',
  'low',
  '22222222-2222-2222-2222-222222222222',
  '11111111-1111-1111-1111-111111111111',
  '11111111-1111-1111-1111-111111111111',
  CURRENT_DATE + INTERVAL '21 days',
  NOW(),
  NOW()
)
ON CONFLICT (id) DO UPDATE
SET title = EXCLUDED.title,
    description = EXCLUDED.description,
    status = EXCLUDED.status,
    priority = EXCLUDED.priority,
    project_id = EXCLUDED.project_id,
    created_by_id = EXCLUDED.created_by_id,
    assignee_id = EXCLUDED.assignee_id,
    due_date = EXCLUDED.due_date,
    updated_at = NOW();
