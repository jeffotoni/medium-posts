--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.9
-- Dumped by pg_dump version 9.6.9

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: login; Type: TABLE; Schema: public; Owner: s3wf
--

CREATE TABLE public.login (
    id integer NOT NULL,
    nome character varying(300) NOT NULL,
    email character varying(300) NOT NULL
);


ALTER TABLE public.login OWNER TO s3wf;

--
-- Name: login_id_seq; Type: SEQUENCE; Schema: public; Owner: s3wf
--

CREATE SEQUENCE public.login_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.login_id_seq OWNER TO s3wf;

--
-- Name: login_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: s3wf
--

ALTER SEQUENCE public.login_id_seq OWNED BY public.login.id;


--
-- Name: login id; Type: DEFAULT; Schema: public; Owner: s3wf
--

ALTER TABLE ONLY public.login ALTER COLUMN id SET DEFAULT nextval('public.login_id_seq'::regclass);


--
-- Name: login login_email_key; Type: CONSTRAINT; Schema: public; Owner: s3wf
--

ALTER TABLE ONLY public.login
    ADD CONSTRAINT login_email_key UNIQUE (email);


--
-- Name: login login_pkey; Type: CONSTRAINT; Schema: public; Owner: s3wf
--

ALTER TABLE ONLY public.login
    ADD CONSTRAINT login_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

