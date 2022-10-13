CREATE TABLE users (
	wallet_id SERIAL PRIMARY KEY,
	name VARCHAR(50) not null,
	email VARCHAR(20) unique not null,
	password VARCHAR(256) not null,
	balance integer null,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


ALTER SEQUENCE users_wallet_id_seq RESTART WITH 100000 INCREMENT BY 1;

INSERT INTO users (name, email, password)
VALUES
('test1', 'test1@shopee.com', '$2a$10$SWvAZsaXq4bAlqQCeF6JjeBZqc6vU7OUS29ELmlbIIV07tfQjlkLq'),
('test2', 'test2@shopee.com', '$2a$10$gLBtAvZvh8ti8ovQW/cWLeqYPQ1iLJ7tLqC2t1XXGK81pvvyJc68S'),
('test3', 'test3@shopee.com', '$2a$10$HYrj0GFcbXUJMCLbTnIBxe3XIWGAQ6umuuqajxdF0QR06ydajUtdu'),
('test4', 'test4@shopee.com', '$2a$10$T68bat3.VnPEfy4/7BZ6fu1IeP3RgBHV8gu67iSBwmoB3FkKJjdAu'),
('test5', 'test5@shopee.com', '$2a$10$ZOM4HekO53evaqBPoejH8.n/Ve0PScJdLRIsvoVvNlJMMECGFbS5G');


create table transactions (
	transaction_id SERIAL primary key,
	wallet_id integer not null, 
	trans_type VARCHAR(20) not null,
	amount integer not null,
	source_id integer null,
	target_id integer null,
	fund_id integer not null,
	description VARCHAR(35) null,
	created_at TIMESTAMP not null default current_timestamp
)

insert into transactions (wallet_id, trans_type, amount, source_id, target_id, fund_id, description, created_at)
values 
(100003,'TOPUP',100000,100003,0,1,'Top up from bank transfer','2022-10-13 09:27:50.906'),
(100003,'TOPUP',100000,100003,0,2,'Top up from credit card','2022-10-13 09:27:56.230'),
(100003,'TOPUP',100000,100003,0,3,'Top up from cash','2022-10-13 09:28:01.896'),
(100003,'TRANSFER',100000,0,100004,0,'payload','2022-10-13 09:28:40.833'),
(100004,'RECEIVED TRANSFER',100000,	100003,	0,0,'payload','2022-10-13 09:28:40.835'),
(100007,'TOPUP',100000,	100007,0,3,	'Top up from cash',	'2022-10-13 09:30:43.093'),
(100007,'TOPUP',100000, 100007,	0,1,'Top up from bank transfer','2022-10-13 09:31:06.654'),
(100007,'TRANSFER',	50000,0,100006,0,'payload',	'2022-10-13 09:31:37.006'),
(100006,'RECEIVED TRANSFER',50000,100007,0,	0,'payload','2022-10-13 09:31:37.008'),
(100003,'TOPUP',100000,100003,0,1,'Top up from bank transfer','2021-10-13 09:27:50.906'),
(100003,'TOPUP',100000,100003,0,2,'Top up from credit card','2021-11-13 09:27:56.230'),
(100003,'TOPUP',100000,100003,0,3,'Top up from cash','2021-11-13 09:28:01.896'),
(100003,'TRANSFER',100000,0,100004,0,'payload','2021-12-13 09:28:40.833'),
(100004,'RECEIVED TRANSFER',100000,	100003,	0,0,'payload','2021-12-13 09:28:40.835'),
(100007,'TOPUP',100000,	100007,0,3,	'Top up from cash',	'2021-12-13 09:30:43.093'),
(100007,'TOPUP',100000, 100007,	0,1,'Top up from bank transfer','2021-12-13 09:31:06.654'),
(100007,'TRANSFER',	50000,0,100006,0,'payload',	'2021-12-15 09:31:37.006'),
(100006,'RECEIVED TRANSFER',50000,100007,0,	0,'payload','2021-12-15 09:31:37.008'),
(100004,'TOPUP',100000,100003,0,1,'Top up from bank transfer','2022-09-13 09:27:50.906'),
(100004,'TOPUP',100000,100003,0,2,'Top up from credit card','2022-09-13 09:27:56.230');

create table funds (
	fund_id SERIAL primary key,
	fund_name varchar(20) not null
)

INSERT INTO funds (fund_name)
VALUES 
('bank transfer'),
('credit card'),
('cash');
