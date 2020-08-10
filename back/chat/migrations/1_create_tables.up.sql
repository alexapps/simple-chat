CREATE TABLE if NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    date_of_creation INTEGER,
    participants INTEGER[]
);