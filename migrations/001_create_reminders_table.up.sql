CREATE TABLE IF NOT EXISTS reminders (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    text TEXT,
    cron_expression TEXT,
    next_run_at TIMESTAMP,
    acknowledged BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_reminders_next_run_at ON reminders(next_run_at);