-- if not done already
-- CREATE USER dba ENCRYPTED PASSWORD '' CREATEDB CREATEROLE SUPERUSER LOGIN INHERIT; -- admin user
-- CREATE USER dbrw ENCRYPTED PASSWORD '' LOGIN INHERIT; -- read write user
-- CREATE USER dbr ENCRYPTED PASSWORD '' LOGIN INHERIT; -- read only user

-- DROP ROLE IF EXISTS admin_role;
-- DROP ROLE IF EXISTS read_write_role;
-- DROP ROLE IF EXISTS read_only_role;

-- CREATE ROLE admin_role CREATEDB CREATEROLE NOINHERIT;
-- CREATE ROLE read_write_role NOINHERIT;
-- CREATE ROLE read_only_role NOINHERIT;

-- GRANT admin_role TO dba;
-- GRANT read_write_role TO dbrw;
-- GRANT read_only_role TO dbr;

-- DROP DATABASE IF EXISTS star_wars_clash;
-- TODO: one per env (drift_prod / drift_test / drift_dev)
CREATE TABLESPACE star_wars_clash_tablespace OWNER dba LOCATION '/data/sql/tablespaces/star_wars_clash';
CREATE DATABASE star_wars_clash_db TABLESPACE star_wars_clash_tablespace OWNER dba ENCODING 'UTF8';
-- revoke default privileges
REVOKE CONNECT ON DATABASE star_wars_clash_db FROM PUBLIC;
-- elevate dba to all privileges
-- ALTER DATABASE star_wars_clash_db owner to dba;
GRANT ALL PRIVILEGES ON DATABASE star_wars_clash_db TO dba;
-- grant connect and access to schema and relevant objects
GRANT CONNECT ON DATABASE star_wars_clash_db TO dbrw, dbr;

-- READY TO TAKE IN THE GENERATED schema.sql file (output of generate_datamodel.py, in ./generated/sql/)
