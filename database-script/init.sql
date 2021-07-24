CREATE TABLE public.generator (
    id character varying(60) NOT NULL,
    name character varying(60) NOT NULL,
    email text NOT NULL,
    password text NOT NULL,

    generated_link text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    expirate_at timestamptz NOT NULL,

    CONSTRAINT generator_pk PRIMARY KEY (id),
    CONSTRAINT generator_unique_1 UNIQUE (email),
    CONSTRAINT generator_unique_2 UNIQUE (generated_link)
);

CREATE TABLE public.contributor (
    email text NOT NULL,
    generated_link text NOT NULL,
    contribution int NOT NULL,

    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,

    CONSTRAINT contributor_unique UNIQUE (email, generated_link),
    CONSTRAINT contributor_fk FOREIGN KEY (generated_link) REFERENCES generator(generated_link)
);