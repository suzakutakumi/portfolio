use portfolio;

DROP TABLE IF EXISTS account;

create table IF NOT EXISTS account(
    session int NOT NULL,
    refresh text NOT NULL,
    access text NOT NULL
);

desc account;
