CREATE TABLE IF NOT EXISTS users (
    id varchar(40) NOT NULL PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    password varchar(100) NOT NULL,
    is_active boolean NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE default NOW(),
	created_by varchar(50),
    updated_at TIMESTAMP WITH TIME ZONE default NOW(), 
	updated_by varchar(50),
    is_deleted boolean NOT NULL DEFAULT false,
    UNIQUE("username","email")
);
create unique index unique_email on public."users" (email) where is_deleted=false;
create unique index unique_username on public."users" (username) where is_deleted=false;
