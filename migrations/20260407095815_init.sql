-- +goose Up
CREATE TABLE IF NOT EXISTS account (
  id UUID PRIMARY KEY,
  name TEXT,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS role (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  slug TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE,
  is_base BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX idx_role_only_one_base
ON role (is_base)
WHERE (is_base IS TRUE);

CREATE TABLE IF NOT EXISTS permission (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  slug TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS role_permission (
  role_id INT REFERENCES role(id) ON DELETE CASCADE,
  permission_id INT REFERENCES permission(id) ON DELETE CASCADE,
  PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS account_role (
  account_id UUID REFERENCES account(id) ON DELETE CASCADE,
  role_id INT REFERENCES role(id) ON DELETE CASCADE,
  PRIMARY KEY (account_id, role_id)
);


-- +goose Down
DROP TABLE IF EXISTS account_role;
DROP TABLE IF EXISTS role_permission;
DROP TABLE IF EXISTS permission;
DROP INDEX IF EXISTS idx_role_only_one_base;
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS account;