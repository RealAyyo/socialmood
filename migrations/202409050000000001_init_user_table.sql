-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         email VARCHAR(255) UNIQUE NOT NULL,
                         password VARCHAR(255) NOT NULL,
                         first_name VARCHAR(255) NOT NULL,
                         last_name VARCHAR(255) NOT NULL,
                         birth TIMESTAMP WITH TIME ZONE NOT NULL,
                         gender VARCHAR(5) NOT NULL,
                         interests VARCHAR(1000) NOT NULL,
                         city VARCHAR(255) NOT NULL,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
