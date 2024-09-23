DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS manufacturers CASCADE;
DROP TABLE IF EXISTS medications CASCADE;
DROP TABLE IF EXISTS borrows CASCADE;
DROP TABLE IF EXISTS borrow_history CASCADE;

-- Создание таблиц

-- Таблица users
-- Описание: Таблица пользователей для хранения учетных данных.
CREATE TABLE users (
user_id SERIAL PRIMARY KEY,
username VARCHAR(100),
password VARCHAR(100)
);

ALTER TABLE users ADD UNIQUE (username);
ALTER TABLE users ADD COLUMN created_at DATE DEFAULT NOW();
ALTER TABLE users ADD COLUMN updated_at DATE;
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255);

-- Таблица manufacturers
--  Описание: Таблица авторов с именем и биографией. Каждый автор может иметь
--  множество книг.
CREATE TABLE manufacturers (
manufacturer_id SERIAL PRIMARY KEY,
name_company VARCHAR(100),
leaflet TEXT
);

ALTER TABLE manufacturers ADD COLUMN address_company VARCHAR(100);
ALTER TABLE manufacturers ADD COLUMN created_at DATE NOT NULL DEFAULT NOW();
ALTER TABLE manufacturers ADD COLUMN updated_at DATE;

-- Таблица medications
-- Описание: Таблица книг, связанная с таблицей авторов через manufacturer_id.
-- Отношение "один к многим" между авторами и книгами.
CREATE TABLE medications (
medication_id SERIAL PRIMARY KEY,
name_medication VARCHAR(255),
manufacturer_id INT,
CONSTRAINT UQ_medications_name_medication UNIQUE (name_medication),
CONSTRAINT FK_medications_manufacturer_id FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(manufacturer_id)
);

ALTER TABLE medications ADD COLUMN created_at DATE NOT NULL DEFAULT NOW();
ALTER TABLE medications ADD COLUMN updated_at DATE;

-- Таблица borrow
-- Описание: Таблица для отслеживания взятых и возвращенных книг.
-- Связь "много ко многим" между пользователями и книгами.
CREATE TABLE borrows (
borrow_id SERIAL PRIMARY KEY,
user_id INT NOT NULL,
medication_id INT NOT NULL,
borrow_date DATE,
return_date DATE
);

ALTER TABLE borrows
ADD CONSTRAINT FK_borrow_user_id FOREIGN KEY (user_id) REFERENCES users(user_id);

ALTER TABLE borrows
ADD CONSTRAINT FK_borrow_medication_id FOREIGN KEY (medication_id) REFERENCES medications(medication_id);

-- Таблица borrow_history
-- Описание: Таблица для хранения истории получения и возврата книг.
-- Каждое действие (взятие или возврат книги) будет записываться в этой таблице.
CREATE TABLE borrow_history (
history_id SERIAL PRIMARY KEY,
borrow_id INT NOT NULL,
action_type VARCHAR(50) CHECK (action_type IN ('borrow', 'return')), -- Тип действия: 'borrow' для взятия, 'return' для возврата
action_date DATE NOT NULL DEFAULT CURRENT_DATE,
CONSTRAINT FK_borrow_history_borrow_id FOREIGN KEY (borrow_id) REFERENCES borrows(borrow_id)
);

-- Таблица profiles
-- Описание: Таблица профилей, связь "один к одному" с таблицей
-- пользователей. Хранит email и адрес пользователя.
CREATE TABLE profiles (
user_id INT PRIMARY KEY,
email VARCHAR(255),
address VARCHAR(255),
CONSTRAINT FK_profiles_user_id FOREIGN KEY (user_id) REFERENCES users(user_id)
);


ALTER TABLE profiles ADD COLUMN created_at DATE NOT NULL DEFAULT NOW();
ALTER TABLE profiles ADD COLUMN updated_at DATE;

-- Таблица reviews
-- Описание: Таблица для хранения отзывов на книги. Связана с таблицами users и medications.
CREATE TABLE reviews (
review_id SERIAL PRIMARY KEY,
user_id INT NOT NULL,
medication_id INT NOT NULL,
review_text TEXT,
rating DECIMAL(2, 1) CHECK (rating >= 1.0 AND rating <= 5.0), -- Рейтинг с одним десятичным знаком
review_date DATE DEFAULT CURRENT_DATE,
CONSTRAINT FK_reviews_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
CONSTRAINT FK_reviews_medication_id FOREIGN KEY (medication_id) REFERENCES medications(medication_id)
);

ALTER TABLE reviews ADD COLUMN created_at DATE NOT NULL DEFAULT NOW();
ALTER TABLE reviews ADD COLUMN updated_at DATE;

-- Добавление начальных данных

TRUNCATE TABLE medications RESTART IDENTITY;
TRUNCATE TABLE borrows RESTART IDENTITY;
TRUNCATE TABLE users RESTART IDENTITY CASCADE;
TRUNCATE TABLE profiles RESTART IDENTITY;
TRUNCATE TABLE reviews RESTART IDENTITY;
TRUNCATE TABLE manufacturers RESTART IDENTITY CASCADE;


-- INSERT INTO manufacturers (name_company, leaflet, address_company)
-- VALUES
--
--
-- INSERT INTO medications (name_medication, manufacturer_id)
-- VALUES

INSERT INTO users (username, password)
VALUES
('johndoe', 'password123'),
('janedoe', 'securepassword'),
('alice', 'alicepass'),
('bob', 'bobpassword'),
('charlie', 'charlie123'),
('david', 'davidpass'),
('eva', 'evapass'),
('frank', 'frankpass'),
('grace', 'gracepass'),
('henry', 'henrypass'),
('isabel', 'isabelpass'),
('jack', 'jackpass'),
('karen', 'karenpass'),
('leo', 'leopass'),
('mia', 'miapass');

INSERT INTO profiles (user_id, email, address)
VALUES
(1, 'johndoe@example.com', '123 Main St'),
(2, 'janedoe@example.com', NULL),          -- Пустое поле адреса
(3, NULL, '789 Elm St'),                   -- Пустое поле email
(4, 'bob@example.com', '321 Pine St'),
(5, NULL, NULL),                           -- Оба поля пустые
(6, 'david@example.com', '987 Cedar St'),
(7, 'eva@example.com', NULL),
(8, NULL, '654 Oak Ave'),
(9, 'grace@example.com', '321 Elm St'),
(10, 'henry@example.com', NULL),
(11, NULL, NULL),
(12, 'jack@example.com', '987 Oak St'),
(13, 'karen@example.com', '456 Maple Ave'),
(14, NULL, '369 Pine St'),
(15, 'mia@example.com', '159 Cedar St');

-- INSERT INTO borrows (user_id, medication_id, borrow_date, return_date)
-- VALUES
--
--
-- INSERT INTO reviews (user_id, medication_id, review_text, rating)
-- VALUES

