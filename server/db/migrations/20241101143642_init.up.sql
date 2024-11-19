-- create users table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- soft delete
    deleted INTEGER NOT NULL DEFAULT 0,
    deleted_at TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_username ON users (username);

-- delete trigger
CREATE TRIGGER IF NOT EXISTS update_deleted_at
    AFTER UPDATE ON users
    WHEN NEW.deleted == 1
BEGIN
    UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- update trigger for updated_at
CREATE TRIGGER IF NOT EXISTS update_updated_at
    AFTER UPDATE ON users
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- create feeds table
CREATE TABLE IF NOT EXISTS feeds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    link TEXT NOT NULL,

    -- base64 encoded favicon
    favicon TEXT,
    description TEXT NOT NULL,

    -- feed last build(updated) date
    last_build_date TEXT NOT NULL,
    category_id INTEGER,

    -- updated when user clicks on item and refresh feed
    unread_count INTEGER NOT NULL DEFAULT 0,
    -- updated when refresh feed
    total_count INTEGER NOT NULL DEFAULT 0,
    -- updated when refresh feed
    recent_update_count INTEGER NOT NULL DEFAULT 0,

    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- create index on title
CREATE INDEX IF NOT EXISTS idx_title ON feeds (title);

-- CREATE TABLE IF NOT EXISTS ITEMS
CREATE TABLE IF NOT EXISTS items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feed_id INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    -- some feeds use description as content
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    -- 0: unread, 1: read
    read INTEGER NOT NULL DEFAULT 0,
    -- item pub date
    pub_date TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- create categories table
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
