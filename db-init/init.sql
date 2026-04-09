-- This script will run automatically when the Postgres container starts
CREATE TABLE IF NOT EXISTS daily_metrics (
    -- id is a unique serial number for every row
    id SERIAL PRIMARY KEY,
    
    -- entry_date must be unique so you can't log the same day twice
    entry_date DATE UNIQUE NOT NULL,

    -- Your 6 specific parameters with range checks
    sleep_quality     INTEGER NOT NULL CHECK (sleep_quality BETWEEN -3 AND 3),
    physical_energy   INTEGER NOT NULL CHECK (physical_energy BETWEEN -3 AND 3),
    focus             INTEGER NOT NULL CHECK (focus BETWEEN -3 AND 3),
    motivation        INTEGER NOT NULL CHECK (motivation BETWEEN -3 AND 3),
    past_view         INTEGER NOT NULL CHECK (past_view BETWEEN -3 AND 3),
    social_activity   INTEGER NOT NULL CHECK (social_activity BETWEEN -3 AND 3),

    -- This helps you see exactly when the entry was physically created
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

