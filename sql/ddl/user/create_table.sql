CREATE TABLE users (
    id serial,
    name varchar(20),
    PRIMARY KEY (id)
);

-- サンプルレコード作成
INSERT INTO
    users (name)
VALUES
('yamato');

INSERT INTO
    users (name)
VALUES
('ottotto');

INSERT INTO
    users (name)
VALUES
('kizuku');