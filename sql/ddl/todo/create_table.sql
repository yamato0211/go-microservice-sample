CREATE TABLE todos (
    id serial,
    user_id int,
    title varchar(30),
    done boolean DEFAULT FALSE,
    PRIMARY KEY (id)
);
