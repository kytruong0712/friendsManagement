--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user (
    id integer NOT NULL PRIMARY KEY,
    name character varying(255),
    email character varying(255) unique,
    friends TEXT [],
    subscribe TEXT [],
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.user ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: -
--

insert into public.user (name, email, friends, subscribe, created_at, updated_at) values ('Tom Nguyen', 'tom@example.com', ARRAY [ 'andrew@example.com','peter@example.com' ], ARRAY [ 'donald@example.com','peter@example.com' ], '2022-08-23 00:00:00', '2022-09-23 00:00:00');
insert into public.user (name, email, created_at, updated_at) values ('Andrew Do', 'andrew@example.com', '2022-10-21 00:00:00', '2022-11-21 00:00:00');
insert into public.user (name, email, created_at, updated_at) values ('Peter Do', 'peter@example.com', '2022-11-21 00:00:00', '2022-12-21 00:00:00');
insert into public.user (name, email, created_at, updated_at) values ('Donald Tran', 'donald@example.com', '2021-07-21 00:00:00', '2021-07-21 00:00:00');

