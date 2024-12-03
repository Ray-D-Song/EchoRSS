-- Create user_setting table
CREATE TABLE IF NOT EXISTS user_setting (
    user_id TEXT PRIMARY KEY,
    OPENAI_API_KEY TEXT DEFAULT '',
    API_ENDPOINT TEXT DEFAULT 'https://api.openai.com/v1/chat/completions',
    TARGET_LANGUAGE TEXT DEFAULT 'Chinese'
);
