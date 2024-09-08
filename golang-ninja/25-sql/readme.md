postgres
docker run -d --name ninja-db -e POSTGRES_PASSWORD=goLANGn1nja -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5433:5432 postgres

docker exec -it ninja-db bash

psql -U postgres

create table users (
id serial not null unique,
name varchar(255) not null,
email varchar(255) not null,
password varchar(255) not null,
registered_at timestamp not null default now()
);

insert into users (name, email, password) values ('vasya', 'vasya@mail.ru', '12345');

select * from users;