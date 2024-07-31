-- +goose Up
-- +goose StatementBegin
CREATE TYPE biome AS ENUM (
  'water',
  'forest',
  'grassland',
  'dessert',
  'mountain'
);

CREATE TABLE dryad_user (
  id serial PRIMARY KEY
);

CREATE TABLE dryad_map (
  id serial PRIMARY KEY,
  dryad_user integer references dryad_user(id)
);

CREATE TABLE map_region (
  id serial PRIMARY KEY,
  x integer NOT NULL,
  y integer NOT NULL,
  dryad_map integer references dryad_map(id),
  CONSTRAINT unique_region_x_y UNIQUE(x, y)
);

CREATE TABLE map_point (
  id serial PRIMARY KEY,
  x integer NOT NULL,
  y integer NOT NULL,
  biome_type biome NOT NULL, 
  region integer references map_region(id) NOT NULL,
  CONSTRAINT unique_point_x_y UNIQUE(x, y)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE map_point;
DROP TABLE map_region;
DROP TYPE biome;
DROP TABLE dryad_map;
DROP TABLE dryad_user;
-- +goose StatementEnd
