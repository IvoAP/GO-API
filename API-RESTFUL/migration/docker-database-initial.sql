create table users(
    id serial primary key,
    username varchar not null,
    first_name varchar  not null,
    last_name varchar  not null,
    email varchar  not null,
    senha varchar  not null,
    phone varchar  not null,
    user_status integer
);

    

INSERT INTO users(username,first_name,last_name,email,senha,phone,user_status) VALUES
(
    'pimentinha',
    'Ivo',
    'Pimenta',
    'ivo@gmail.com',
    '123456',
    '85999863898',
    '1'
 ),
 (
    'Xolo',
    'Miguel',
    'Dias',
    'xolo@gmail.com',
    '245678',
    '85999864598',
    '1'
 )
