-- translate cache
CREATE TABLE IF NOT EXISTS translate_cache (
    url TEXT PRIMARY KEY,
    content TEXT
);

-- summary cache
CREATE TABLE IF NOT EXISTS summary_cache (
    url TEXT PRIMARY KEY,
    content TEXT
);
