-- create users table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    refresh_token TEXT,
    banned INTEGER NOT NULL DEFAULT 0,
    banned_at TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_username ON users (username);

-- create feeds table
CREATE TABLE IF NOT EXISTS feeds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    -- base64 encoded favicon
    favicon TEXT NOT NULL,
    description TEXT NOT NULL,
    last_build_date TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- create index on title
CREATE INDEX IF NOT EXISTS idx_title ON feeds (title);

-- CREATE TABLE IF NOT EXISTS ITEMS
CREATE TABLE IF NOT EXISTS items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feed_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    description TEXT NOT NULL,
    pub_date TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
