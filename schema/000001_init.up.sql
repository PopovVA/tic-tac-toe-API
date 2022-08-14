CREATE TABLE players
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE games
(
    id serial not null unique,
    winner int references players (id) on delete cascade null
);

CREATE TABLE moves
(
    id serial not null unique,
    playerId int references players (id) on delete cascade not null,
    gameId int REFERENCES games (id) on delete CASCADE not NULL,
    rowNumber int not null,
    columnNumber int not null,
    movedAt int not null
);