
--Создание табличи юзеров
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users
(
    user_id            SERIAL PRIMARY KEY                             NOT NULL,
    user_name          VARCHAR                                        NOT NULL,
    user_hash_password VARCHAR                                        NOT NULL,
    user_role          VARCHAR CHECK (user_role IN ('admin', 'user')) NOT NULL
);

--Создание табличи заказов
DROP TABLE IF EXISTS orders;
CREATE TABLE IF NOT EXISTS orders
(
    order_id      SERIAL PRIMARY KEY                                               NOT NULL,
    delivery_type VARCHAR CHECK ( delivery_type IN ('by_courier', 'self_pickup') ) NOT NULL,
    delivery_time TIMESTAMP                                                        NOT NULL,
    order_time    TIMESTAMP                                                        NOT NULL,
    total_price   FLOAT                                                            NOT NULL,
    canceled      BOOLEAN                                                          NOT NULL
);
