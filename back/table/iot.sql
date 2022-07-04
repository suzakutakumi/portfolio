use portfolio;

DROP TABLE IF EXISTS iot;

create table IF NOT EXISTS iot(
    id int NOT NULL,
    nickName text,
    kind int,
    userId text NOT NULL,
    PRIMARY KEY(id)
);

desc iot;