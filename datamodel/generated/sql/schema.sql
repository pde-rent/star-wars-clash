-- GENERATED - DO NOT MODIFY

-- inspired from pg_dump: PostgreSQL database version 11.9 (Ubuntu 11.9-1.pgdg18.04+1)
-- Customized

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', 'public', false);
--SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;
SET default_tablespace = '';
SET default_with_oids = false;
SET TIMEZONE='GMT';

-- DROP EXISTING SCHEMA

DROP SCHEMA IF EXISTS public CASCADE;
DROP SCHEMA IF EXISTS audit CASCADE;

CREATE SCHEMA public;
CREATE SCHEMA audit;

-- schema permissions
GRANT USAGE ON SCHEMA public TO dbrw, dbr;
GRANT ALL PRIVILEGES ON SCHEMA public, audit TO dba;

-- public objects
-- GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO dbrw;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dbrw, dba;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO dbrw, dba;
GRANT ALL PRIVILEGES ON ALL PROCEDURES IN SCHEMA public TO dbrw, dba;
GRANT ALL PRIVILEGES ON ALL ROUTINES IN SCHEMA public TO dbrw, dba;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO dbrw, dba;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO dbr;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO dbr;
GRANT EXECUTE ON ALL PROCEDURES IN SCHEMA public TO dbr;
GRANT EXECUTE ON ALL ROUTINES IN SCHEMA public TO dbr;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO dbr;

-- audit objects
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA audit TO dba;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA audit TO dba;
GRANT ALL PRIVILEGES ON ALL PROCEDURES IN SCHEMA audit TO dba;
GRANT ALL PRIVILEGES ON ALL ROUTINES IN SCHEMA audit TO dba;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA audit TO dba;


-- GENERATED - DO NOT MODIFY
CREATE TYPE public.affiliation AS ENUM (
    'JEDI_ORDER',
    'SITH_ORDER',
    'GALACTIC_REPUBLIC',
    'REBEL_ALLIANCE',
    'GALACTIC_EMPIRE',
    'RESISTANCE',
    'NEW_REPUBLIC',
    'FIRST_ORDER',
    'CONFEDERACY_OF_INDEPENDENT_SYSTEMS',
    'TRADE_FEDERATION',
    'BOUNTY_HUNTERS',
    'MERCENARIES',
    'CRIMINALS',
    'UNAFFILIATED'
);
ALTER TYPE public.affiliation OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TYPE public.climate AS ENUM (
    'TEMPERATE',
    'HOT',
    'ARID',
    'MOIST',
    'DRY',
    'TROPICAL',
    'ARCTIC',
    'SUBARCTIC',
    'UNKNOWN'
);
ALTER TYPE public.climate OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TYPE public.color AS ENUM (
    'BLACK',
    'BLOND',
    'BLUE',
    'BLUE_GRAY',
    'BROWN',
    'DARK',
    'FAIR',
    'GOLD',
    'GREEN',
    'GREY',
    'HAZEL',
    'LIGHT',
    'METAL',
    'NONE',
    'ORANGE',
    'PALE',
    'PINK',
    'RED',
    'SILVER',
    'TAN',
    'WHITE',
    'YELLOW'
);
ALTER TYPE public.color OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TYPE public.gender AS ENUM (
    'MALE',
    'FEMALE',
    'HERMAPHRODITE',
    'NA'
);
ALTER TYPE public.gender OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TYPE public.species_type AS ENUM (
    'AMPHIBIAN',
    'ARTIFICIAL',
    'GASTROPOD',
    'INSECTOID',
    'MAMMAL',
    'MAMMALS',
    'OTHER',
    'REPTILE',
    'REPTILIAN',
    'UNKOWN'
);
ALTER TYPE public.species_type OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.character_absolute_scores (
    "id"                       bigint                     generated by default as identity not null,
    "ref"                      bigint                     not null,
    "likes"                    integer                    not null,
    "dislikes"                 integer                    not null,
    PRIMARY KEY ("id")
);
ALTER TABLE public.character_absolute_scores OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.character_cross_scores (
    "id"                       bigint                     generated by default as identity not null,
    "ref"                      bigint                     not null,
    "cmp"                      bigint                     not null,
    "ref_likes"                integer                    not null,
    "cmp_likes"                integer                    not null,
    PRIMARY KEY ("id")
);
ALTER TABLE public.character_cross_scores OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.planets (
    "id"                       bigint                     generated by default as identity not null,
    "name"                     varchar                    unique not null,
    "days_to_rotation"         integer                    not null,
    "days_to_orbit"            integer                    not null,
    "diameter_km"              integer                    not null,
    "climate"                  public.climate             not null,
    "standard_gravity"         real                       not null,
    "terrain"                  varchar                    not null,
    "surface_water"            real                       not null,
    "population"               bigint                     not null,
    PRIMARY KEY ("id")
);
ALTER TABLE public.planets OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.planet_absolute_scores (
    "id"                       bigint                     generated by default as identity not null,
    "ref"                      bigint                     not null,
    "likes"                    integer                    not null,
    "dislikes"                 integer                    not null,
    PRIMARY KEY ("id")
);
ALTER TABLE public.planet_absolute_scores OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.planet_cross_scores (
    "id"                       bigint                     generated by default as identity not null,
    "ref"                      bigint                     not null,
    "cmp"                      bigint                     not null,
    "ref_likes"                integer                    not null,
    "cmp_likes"                integer                    not null,
    PRIMARY KEY ("id")
);
ALTER TABLE public.planet_cross_scores OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.species (
    "id"                       bigint                     generated by default as identity not null,
    "name"                     varchar                    unique not null,
    "species_type"             public.species_type        not null,
    "average_height"           integer                    not null,
    "average_lifespan"         integer                    not null,
    "language"                 varchar                    not null,
    "homeworld"                bigint                     references public.planets,
    PRIMARY KEY ("id")
);
ALTER TABLE public.species OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.character_skin_colors (
    "id" bigint references public.species,
    "color" public.color,
    primary key("id", "color")
);
ALTER TABLE public.character_skin_colors OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.character_hair_colors (
    "id" bigint references public.species,
    "color" public.color,
    primary key("id", "color")
);
ALTER TABLE public.character_hair_colors OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.character_eye_colors (
    "id" bigint references public.species,
    "color" public.color,
    primary key("id", "color")
);
ALTER TABLE public.character_eye_colors OWNER TO dba;

-- GENERATED - DO NOT MODIFY
CREATE TABLE public.characters (
    "id"                       bigint                     generated by default as identity not null,
    "name"                     varchar                    unique not null,
    "height"                   integer                    not null,
    "mass"                     integer                    not null,
    "hair_color"               public.color               not null,
    "skin_color"               public.color               not null,
    "eye_color"                public.color               not null,
    "birth_year"               real                       not null,
    "gender"                   public.gender              not null,
    "homeworld"                bigint                     references public.planets,
    "species"                  bigint                     references public.species,
    "main_affiliation"         public.affiliation         not null,
    PRIMARY KEY ("id")
);
ALTER TABLE public.characters OWNER TO dba;
