CREATE TABLE public.kana
(
    id integer NOT NULL DEFAULT nextval('kana_id_seq'::regclass),
    representation text COLLATE pg_catalog."default" NOT NULL,
    romaji text COLLATE pg_catalog."default" NOT NULL,
    classification integer,
    initial text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT kana_pkey PRIMARY KEY (id),
    CONSTRAINT kana_unique_constr UNIQUE (representation)
)