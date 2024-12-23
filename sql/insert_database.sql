--Вставка данных в табличку заказов 
INSERT INTO orders (delivery_type, delivery_time, order_time, total_price, canceled)
VALUES ('by_courier', '2023-10-01 10:00:00', '2023-10-01 09:30:00', 150.75, FALSE),
       ('self_pickup', '2023-10-01 11:00:00', '2023-10-01 09:45:00', 80.50, FALSE),
       ('by_courier', '2023-10-01 12:00:00', '2023-10-01 10:15:00', 120.25, TRUE),
       ('self_pickup', '2023-10-01 13:00:00', '2023-10-01 10:45:00', 95.00, FALSE),
       ('by_courier', '2023-10-01 14:00:00', '2023-10-01 11:30:00', 200.00, FALSE),
       ('self_pickup', '2023-10-01 15:00:00', '2023-10-01 12:00:00', 60.00, TRUE),
       ('by_courier', '2023-10-01 16:00:00', '2023-10-01 13:00:00', 300.90, FALSE),
       ('self_pickup', '2023-10-01 17:00:00', '2023-10-01 14:30:00', 140.25, FALSE),
       ('by_courier', '2023-10-01 18:00:00', '2023-10-01 15:00:00', 375.50, TRUE),
       ('self_pickup', '2023-10-01 19:00:00', '2023-10-01 16:30:00', 22.00, FALSE);

--Вставка данных в табличку юзеров
INSERT INTO users (user_name, user_hash_password, user_role)
VALUES ('user1', 'user1', 'admin'),
       ('JohnDoe', 'hashed_password_1', 'user'),
       ('JaneSmith', 'hashed_password_2', 'admin'),
       ('AliceJohnson', 'hashed_password_3', 'user'),
       ('BobBrown', 'hashed_password_4', 'user'),
       ('CharlieDavis', 'hashed_password_5', 'admin'),
       ('DianaEvans', 'hashed_password_6', 'user'),
       ('EthanFoster', 'hashed_password_7', 'user'),
       ('FionaGreen', 'hashed_password_8', 'admin'),
       ('GeorgeHill', 'hashed_password_9', 'user'),
       ('HannahIvy', 'hashed_password_10', 'user');