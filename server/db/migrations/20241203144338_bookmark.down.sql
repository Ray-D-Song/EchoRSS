-- create new items table
CREATE TABLE new_items (
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

-- copy data from old items table to new items table
INSERT INTO new_items (id, feed_id, user_id, title, link, description, content, read, pub_date, created_at)
SELECT id, feed_id, user_id, title, link, description, content, read, pub_date, created_at FROM items;

-- drop old items table
DROP TABLE items;

-- rename new items table to items
ALTER TABLE new_items RENAME TO items;
