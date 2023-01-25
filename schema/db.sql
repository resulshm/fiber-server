DROP TABLE IF EXISTS videos;

CREATE TABLE videos(
    id SERIAL NOT NULL UNIQUE ,
    title VARCHAR(255) NOT NULL,
    description TEXT DEFAULT '',
    likes INT DEFAULT 0,
    dislikes INT DEFAULT 0,
    views INT DEFAULT 0
);
