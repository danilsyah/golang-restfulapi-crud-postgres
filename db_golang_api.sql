--
-- PostgreSQL database dump
--

-- Dumped from database version 11.0
-- Dumped by pg_dump version 11.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: buku; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.buku (
    id integer NOT NULL,
    judul_buku character varying NOT NULL,
    penulis character varying,
    tgl_publikasi date
);


ALTER TABLE public.buku OWNER TO postgres;

--
-- Name: buku_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.buku_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buku_id_seq OWNER TO postgres;

--
-- Name: buku_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.buku_id_seq OWNED BY public.buku.id;


--
-- Name: buku id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buku ALTER COLUMN id SET DEFAULT nextval('public.buku_id_seq'::regclass);


--
-- Data for Name: buku; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.buku (id, judul_buku, penulis, tgl_publikasi) FROM stdin;
1	Matematika Diskrit	Selviana	2009-06-20
2	Pemrograman Golang	Danil	2024-02-19
4	Konsep OOP	Haykal	2024-02-19
5	Buku Basic Laravel	Danil	2024-02-20
6	Buku Basic Golang	Danil	2020-01-01
\.


--
-- Name: buku_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.buku_id_seq', 6, true);


--
-- Name: buku pk_buku; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buku
    ADD CONSTRAINT pk_buku PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

