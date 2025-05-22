CREATE SEQUENCE IF NOT EXISTS user_id START WITH 1;
CREATE TABLE users
(
    id INTEGER PRIMARY KEY DEFAULT nextval('user_id'),
    name CHARACTER VARYING(30),
    phone CHARACTER VARYING(12) UNIQUE,
    password CHARACTER VARYING(30)
);
-- 1 - Лилия Малышева - 8989898989 - pass
-- 2 - Алина Малинина - 9898989898 - pass
-- 3 - Таня Лобова - 89829808130 - rootcisco123
-- 4 - Настя Змановская - 89898989898 - ciscoroot321
CREATE SEQUENCE IF NOT EXISTS master_id START WITH 1;
CREATE TABLE master
(
    id INTEGER PRIMARY KEY DEFAULT nextval('master_id'),
    user_id INTEGER REFERENCES users (id)
);

CREATE SEQUENCE IF NOT EXISTS client_id START WITH 1;
CREATE TABLE client
(
    id INTEGER PRIMARY KEY DEFAULT nextval('client_id'),
    user_id INTEGER REFERENCES users (id)
);

CREATE SEQUENCE IF NOT EXISTS specialisation_id START WITH 1;
CREATE TABLE specialisation
(
    id INTEGER PRIMARY KEY DEFAULT nextval('specialisation_id'),
    name CHARACTER VARYING(30) UNIQUE
);
-- 1 - мастер по маникюру
-- 2 - парикмахер
-- 3 - дерматолог
-- 4 - визажист
-- 5 - косметолог

CREATE SEQUENCE IF NOT EXISTS service_id START WITH 1;
CREATE TABLE service
(
    id INTEGER PRIMARY KEY DEFAULT nextval('service_id'),
    name CHARACTER VARYING(30) UNIQUE,
    price DECIMAL(10, 2),
    time INTEGER,
    specialisation_id INTEGER REFERENCES specialisation (id)
);
-- 1 - наращивание - 500 - 50 - 1(мастер по маникюру)
-- 2 - простой дизайн - 300 - 30 - 1(мастер по маникюру)
-- 3 - сложный дизайн - 500 - 60 - 1(мастер по маникюру)
-- 4 - укладка - 700 - 90 - 2(парикмахер)
-- 5 - стрижка - 1000 - 80 - 2(парикмахер)
-- 6 - окрашивание - 3000 - 120 - 2(парикмахер)
-- 7 - простой макияж - 500 - 50 - 4(визажист)
-- 8 - сложный макияж - 100 - 90 - 4(визажист)
-- 9 - прием косметолога - 1000 - 15 - 5(косметолог)
-- 10 - выжигание новообразований - 1000 - 10 - 5(косметолог)
-- 11 - прием дерматолога - 1000 - 20 - 3(дерматолог)

CREATE SEQUENCE IF NOT EXISTS master_specialisation_id START WITH 1;
CREATE TABLE master_specialisation
(
    id INTEGER PRIMARY KEY DEFAULT nextval('master_specialisation_id'),
    master_id INTEGER REFERENCES master (id),
    specialisation INTEGER REFERENCES specialisation (id)
);
-- 1 - 1(Лилия Малышева) - 1 (мастер по маникюру)
-- 2 - 1(Лилия Малышева) - 4 (визажист)
-- 3 - 2(Алина Малинина) - 2 (парикмахер)
-- 4 - 2(Алина Малинина) - 4 (визажист)
-- 5 - 2(Алина Малинина) - 3 (дерматолог)
-- 6 - 2(Алина Малинина) - 5 (косметолог)

CREATE SEQUENCE IF NOT EXISTS receipt_id START WITH 1;
CREATE TABLE receipt
(
    id INTEGER PRIMARY KEY DEFAULT nextval('receipt_id'),
    client_id INTEGER REFERENCES client (id),
    date TIMESTAMP,
    debt BOOLEAN
);
-- 1 - 1(Таня Лобова) - 11:15:00:20.05.2025 - 0
-- 2 - 1(Таня Лобова) - 19:45:00-20.05.2025 - 1
-- 3 - 2(Настя Змановская) - 18:00:00-22.05.2025 - 0
CREATE TYPE appointment_status AS ENUM ('booked', 'completed', 'canceled_correctly', 'canceled_incorrectly');

CREATE SEQUENCE IF NOT EXISTS receipt_service_id START WITH 1;
CREATE TABLE receipt_service
(
    id INTEGER PRIMARY KEY DEFAULT nextval('receipt_service_id'),
    receipt_id INTEGER REFERENCES receipt (id),
    service_id INTEGER REFERENCES service (id),
    master_id INTEGER REFERENCES master (id),
    number INTEGER,
    status TEXT
);
-- 1 - 1(чек) - 5(стрижка) - 2(Алина Малинина) - 2(порядок) - completed
-- 2 - 1(чек) - 5(укладка) - 2(Алина Малинина) - 1(порядок) - completed
-- 3 - 2(чек) - 1(наращивание) - 1(Лилия Малышева) - 1(порядок) - canceled_incorrectly
-- 4 - 3(чек) - 1(наращивание) - 1(Лилия Малышева) - 1(порядок) - booked

