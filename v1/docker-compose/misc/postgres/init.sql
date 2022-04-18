DROP TABLE IF EXISTS currencies;
DROP TABLE IF EXISTS requests;

CREATE TABLE currencies (
	id serial4 NOT NULL,
	code varchar(3) NULL,
	value numeric(15, 6) NULL,
	last_update_at int8 NULL,
	created_at int8 NULL,
	updated_at int8 NULL,
	CONSTRAINT currencies_pk PRIMARY KEY (id)
);

CREATE TABLE requests (
	id serial4 NOT NULL,
	"date" timestamp NULL,
	duration varchar(20) NULL,
	status varchar(6) NULL,
	created_at int8 NULL,
	updated_at int8 NULL,
	CONSTRAINT requests_pk PRIMARY KEY (id)
);