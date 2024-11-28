-- create web page cache table
CREATE TABLE IF NOT EXISTS web_page_cache (
    url TEXT NOT NULL PRIMARY KEY,
    content TEXT NOT NULL
);
