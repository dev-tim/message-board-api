CREATE TABLE IF NOT EXISTS messages(
    id varchar not null primary key,
    name varchar not null,
    email varchar not null unique,
    text varchar not null,
    external_creation_time timestamptz not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
)
