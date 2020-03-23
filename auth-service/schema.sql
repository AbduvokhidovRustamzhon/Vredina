Create table if not exists users
(
    id       Bigserial primary key,
    name text not null ,
    login    Text UNIQUE not null,
    password text        not null
);

Create table if not exists mitings
(
    id       Bigserial primary key,
    status   bool,
    timeInHour int,
    timeInMinutes int,
    timeOutHour int,
    timeOutMinutes int
);
