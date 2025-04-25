
CREATE TABLE IF NOT EXISTS genders (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS nationalities (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT,
    age INTEGER,
    gender_id INTEGER REFERENCES genders(id) ON DELETE SET NULL,
    nationality_id INTEGER REFERENCES nationalities(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
