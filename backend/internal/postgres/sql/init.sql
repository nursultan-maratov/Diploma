CREATE TABLE users
(
    id         INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name VARCHAR(100),
    last_name  VARCHAR(100),
    email      VARCHAR(100) UNIQUE,
    password   VARCHAR(255),
    phone      VARCHAR(20),
    address    TEXT,
    status     VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE companies
(
    id           INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    brand        VARCHAR(100),
    bank_account VARCHAR(100),
    users_id     INT REFERENCES users (id),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP
);

CREATE TABLE products
(
    id             INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name           VARCHAR(100),
    description    TEXT,
    stock_quantity int,
    price          DECIMAL(10, 2),
    company_id     INT REFERENCES companies (id),
    status         VARCHAR(50),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at     TIMESTAMP
);

CREATE TABLE orders
(
    id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id     INT REFERENCES users (id),
    product_id  INT REFERENCES products (id),
    quantity    INT,
    total_price DECIMAL(10, 2),
    order_date  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status      VARCHAR(50),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
);


BEGIN;

INSERT INTO users (first_name, last_name, email, password, phone, address, status)
VALUES ('Drake', 'Unknown', 'drake@example.com', 'password', '123-456-7890', 'Address 1', 'buyer'),
       ('Taylor', 'Swift', 'taylor.swift@example.com', '123456', '123-456-7891', 'Address 2', 'buyer'),
       ('Bad', 'Bunny', 'bad.bunny@example.com', '123456789', '123-456-7892', 'Address 3', 'buyer'),
       ('The', 'Weeknd', 'the.weeknd@example.com', 'guest', '123-456-7893', 'Address 4', 'buyer'),
       ('Justin', 'Bieber', 'justin.bieber@example.com', 'qwerty', '123-456-7894', 'Address 5', 'buyer'),
       ('Ed', 'Sheeran', 'ed.sheeran@example.com', '12345678', '123-456-7895', 'Address 6', 'buyer'),
       ('Eminem', 'Unknown', 'eminem@example.com', '111111', '123-456-7896', 'Address 7', 'buyer'),
       ('Ariana', 'Grande', 'ariana.grande@example.com', '1234511', '123-456-7897', 'Address 8', 'seller'),
       ('Travis', 'Scott', 'travis.scott@example.com', '696969', '123-456-7898', 'Address 9', 'seller'),
       ('Kanye', 'West', 'kanye.west@example.com', '000000', '123-456-7899', 'Address 10', 'seller'),
       ('Post', 'Malone', 'post.malone@example.com', '24022022', '123-456-7800', 'Address 11', 'buyer'),
       ('Rihanna', 'Unknown', 'rihanna@example.com', 'putin', '123-456-7801', 'Address 12', 'seller'),
       ('J', 'Balvin', 'j.balvin@example.com', 'abc123', '123-456-7802', 'Address 13', 'buyer'),
       ('BTS', 'Group', 'bts@example.com', 'access', '123-456-7803', 'Address 14', 'buyer'),
       ('Future', 'Unknown', 'future@example.com', 'baseball', '123-456-7804', 'Address 15', 'buyer'),
       ('Billie', 'Eilish', 'billie.eilish@example.com', 'batman', '123-456-7805', 'Address 16', 'seller'),
       ('Kendrick', 'Lamar', 'kendrick.lamar@example.com', 'dragon', '123-456-7806', 'Address 17', 'buyer'),
       ('Juice', 'WRLD', 'juice.wrld@example.com', 'football', '123-456-7807', 'Address 18', 'buyer'),
       ('Ozuna', 'Unknown', 'ozuna@example.com', 'letmein', '123-456-7808', 'Address 19', 'buyer'),
       ('Nicki', 'Minaj', 'nicki.minaj@example.com', 'master', '123-456-7809', 'Address 20', 'seller'),
       ('Bruno', 'Mars', 'bruno.mars@example.com', 'michael', '123-456-7810', 'Address 21', 'buyer'),
       ('Coldplay', 'Group', 'coldplay@example.com', 'monkey', '123-456-7811', 'Address 22', 'buyer'),
       ('Imagine', 'Dragons', 'imagine.dragons@example.com', 'mustang', '123-456-7812', 'Address 23', 'seller'),
       ('XXX', 'TENTACION', 'xxxtentacion@example.com', 'shadow', '123-456-7813', 'Address 24', 'buyer'),
       ('Chris', 'Brown', 'chris.brown@example.com', 'superman', '123-456-7814', 'Address 25', 'buyer'),
       ('Dua', 'Lipa', 'dua.lipa@example.com', 'trustno1', '123-456-7815', 'Address 26', 'buyer'),
       ('21', 'Savage', '21.savage@example.com', 'ukfreedom', '123-456-7816', 'Address 27', 'buyer'),
       ('Lil', 'Wayne', 'lil.wayne@example.com', 'king1qaz2wsx', '123-456-7817', 'Address 28', 'seller'),
       ('Khalid', 'Unknown', 'khalid@example.com', 'qazxsw', '123-456-7818', 'Address 29', 'buyer'),
       ('David', 'Guetta', 'david.guetta@example.com', 'qaz2wsx', '123-456-7819', 'Address 30', 'buyer'),
       ('Daddy', 'Yankee', 'daddy.yankee@example.com', 'wsx3edc', '123-456-7820', 'Address 31', 'seller'),
       ('Anuel', 'AA', 'anuel.aa@example.com', 'q2w3e4r', '123-456-7821', 'Address 32', 'seller'),
       ('Lana', 'Del Rey', 'lana.del.rey@example.com', 'zaq12wsx', '123-456-7822', 'Address 33', 'buyer'),
       ('Maroon', '5', 'maroon5@example.com', 'col123456', '123-456-7823', 'Address 34', 'buyer'),
       ('Rauw', 'Alejandro', 'rauw.alejandro@example.com', '123123', '123-456-7824', 'Address 35', 'buyer'),
       ('Lil', 'Baby', 'lil.baby@example.com', 'password', '123-456-7825', 'Address 36', 'seller'),
       ('Beyoncé', 'Unknown', 'beyonce@example.com', '123456', '123-456-7826', 'Address 37', 'seller'),
       ('Maluma', 'Unknown', 'maluma@example.com', 'maluma', '123-456-7890', 'Address 1', 'buyer'),
       ('Calvin', 'Harris', 'calvin.harris@example.com', 'calvin', '123-456-7891', 'Address 2', 'buyer'),
       ('KAROL', 'G', 'karol.g@example.com', 'karol', '123-456-7892', 'Address 3', 'buyer'),
       ('Lil', 'Uzi Vert', 'lil.uzi.vert@example.com', 'lil', '123-456-7893', 'Address 4', 'buyer'),
       ('J.', 'Cole', 'j.cole@example.com', 'cole', '123-456-7894', 'Address 5', 'buyer'),
       ('SZA', 'Unknown', 'sza@example.com', 'sza', '123-456-7895', 'Address 6', 'buyer'),
       ('Shawn', 'Mendes', 'shawn.mendes@example.com', 'shawn', '123-456-7896', 'Address 7', 'buyer'),
       ('Sia', 'Unknown', 'sia@example.com', 'sia', '123-456-7897', 'Address 8', 'buyer'),
       ('Young', 'Thug', 'young.thug@example.com', 'young', '123-456-7898', 'Address 9', 'buyer'),
       ('Myke', 'Towers', 'myke.towers@example.com', 'myke', '123-456-7899', 'Address 10', 'buyer'),
       ('Sam', 'Smith', 'sam.smith@example.com', 'sam', '123-456-7800', 'Address 11', 'buyer'),
       ('Feid', 'Unknown', 'feid@example.com', 'feid', '123-456-7801', 'Address 12', 'buyer'),
       ('Queen', 'Group', 'queen@example.com', 'queen', '123-456-7802', 'Address 13', 'buyer'),
       ('Farruko', 'Unknown', 'farruko@example.com', 'farruko', '123-456-7803', 'Address 14', 'buyer'),
       ('Lady', 'Gaga', 'lady.gaga@example.com', 'lady', '123-456-7804', 'Address 15', 'buyer'),
       ('Doja', 'Cat', 'doja.cat@example.com', 'doja', '123-456-7805', 'Address 16', 'buyer'),
       ('Harry', 'Styles', 'harry.styles@example.com', 'harry', '123-456-7806', 'Address 17', 'buyer'),
       ('One', 'Direction', 'one.direction@example.com', 'one', '123-456-7807', 'Address 18', 'buyer'),
       ('Adele', 'Unknown', 'adele@example.com', 'adele', '123-456-7808', 'Address 19', 'buyer'),
       ('Shakira', 'Unknown', 'shakira@example.com', 'shakira', '123-456-7809', 'Address 20', 'buyer'),
       ('Selena', 'Gomez', 'selena.gomez@example.com', 'selena', '123-456-7810', 'Address 21', 'buyer'),
       ('Metro', 'Boomin', 'metro.boomin@example.com', 'metro', '123-456-7811', 'Address 22', 'buyer'),
       ('Ty', 'Dolla $ign', 'ty.dolla.sign@example.com', 'ty', '123-456-7812', 'Address 23', 'buyer'),
       ('Halsey', 'Unknown', 'halsey@example.com', 'halsey', '123-456-7813', 'Address 24', 'buyer'),
       ('Linkin', 'Park', 'linkin.park@example.com', 'linkin', '123-456-7814', 'Address 25', 'buyer'),
       ('Nicky', 'Jam', 'nicky.jam@example.com', 'nicky', '123-456-7815', 'Address 26', 'buyer'),
       ('Katy', 'Perry', 'katy.perry@example.com', 'katy', '123-456-7816', 'Address 27', 'buyer'),
       ('The', 'Beatles', 'the.beatles@example.com', 'the', '123-456-7817', 'Address 28', 'buyer'),
       ('Arijit', 'Singh', 'arijit.singh@example.com', 'arijit', '123-456-7818', 'Address 29', 'buyer'),
       ('Gunna', 'Unknown', 'gunna@example.com', 'gunna', '123-456-7819', 'Address 30', 'buyer'),
       ('Arctic', 'Monkeys', 'arctic.monkeys@example.com', 'arctic', '123-456-7820', 'Address 31', 'buyer'),
       ('The', 'Chainsmokers', 'the.chainsmokers@example.com', 'the', '123-456-7821', 'Address 32', 'buyer'),
       ('Marshmello', 'Unknown', 'marshmello@example.com', 'marshmello', '123-456-7822', 'Address 33', 'buyer'),
       ('Wiz', 'Khalifa', 'wiz.khalifa@example.com', 'wiz', '123-456-7823', 'Address 34', 'buyer'),
       ('Olivia', 'Rodrigo', 'olivia.rodrigo@example.com', 'olivia', '123-456-7824', 'Address 35', 'buyer'),
       ('Miley', 'Cyrus', 'miley.cyrus@example.com', 'miley', '123-456-7825', 'Address 36', 'buyer'),
       ('Cardi', 'B', 'cardi.b@example.com', 'cardi', '123-456-7826', 'Address 37', 'buyer'),
       ('JAY', 'Z', 'jay.z@example.com', 'jay', '123-456-7827', 'Address 38', 'buyer'),
       ('Morgan', 'Wallen', 'morgan.wallen@example.com', 'morgan', '123-456-7890', 'Address 1', 'buyer'),
       ('Peso', 'Pluma', 'peso.pluma@example.com', 'peso', '123-456-7891', 'Address 2', 'buyer'),
       ('A$AP', 'Rocky', 'asap.rocky@example.com', 'asap', '123-456-7892', 'Address 3', 'buyer'),
       ('Camila', 'Cabello', 'camila.cabello@example.com', 'camila', '123-456-7893', 'Address 4', 'buyer'),
       ('Pitbull', 'Unknown', 'pitbull@example.com', 'pitbull', '123-456-7894', 'Address 5', 'buyer'),
       ('Frank', 'Ocean', 'frank.ocean@example.com', 'frank', '123-456-7895', 'Address 6', 'buyer'),
       ('DaBaby', 'Unknown', 'dababy@example.com', 'dababy', '123-456-7896', 'Address 7', 'buyer'),
       ('Twenty', 'One Pilots', 'twenty.one.pilots@example.com', 'twenty', '123-456-7897', 'Address 8', 'buyer'),
       ('Quavo', 'Unknown', 'quavo@example.com', 'quavo', '123-456-7898', 'Address 9', 'buyer'),
       ('Avicii', 'Unknown', 'avicii@example.com', 'avicii', '123-456-7899', 'Address 10', 'buyer'),
       ('Kygo', 'Unknown', 'kygo@example.com', 'kygo', '123-456-7800', 'Address 11', 'buyer'),
       ('Tyler', 'The Creator', 'tyler.the.creator@example.com', 'tyler', '123-456-7801', 'Address 12', 'buyer'),
       ('Sech', 'Unknown', 'sech@example.com', 'sech', '123-456-7802', 'Address 13', 'buyer'),
       ('One', 'Republic', 'one.republic@example.com', 'one', '123-456-7803', 'Address 14', 'buyer'),
       ('Red Hot', 'Chili Peppers', 'red.hot.chili.peppers@example.com', 'redhot', '123-456-7804', 'Address 15',
        'buyer'),
       ('Snoop', 'Dogg', 'snoop.dogg@example.com', 'snoop', '123-456-7805', 'Address 16', 'buyer'),
       ('Trippie', 'Redd', 'trippie.redd@example.com', 'trippie', '123-456-7806', 'Address 17', 'buyer'),
       ('Tyga', 'Unknown', 'tyga@example.com', 'tyga', '123-456-7807', 'Address 18', 'buyer'),
       ('Junior', 'H', 'junior.h@example.com', 'junior', '123-456-7808', 'Address 19', 'buyer'),
       ('Michael', 'Jackson', 'michael.jackson@example.com', 'michael', '123-456-7809', 'Address 20', 'buyer'),
       ('$uicideboy$', 'Unknown', 'suicideboys@example.com', 'suicide', '123-456-7810', 'Address 21', 'buyer'),
       ('Arcángel', 'Unknown', 'arcangel@example.com', 'arcangel', '123-456-7811', 'Address 22', 'buyer'),
       ('Playboi', 'Carti', 'playboi.carti@example.com', 'playboi', '123-456-7812', 'Address 23', 'buyer'),
       ('Pritam', 'Unknown', 'pritam@example.com', 'pritam', '123-456-7813', 'Address 24', 'buyer'),
       ('Jason', 'Derulo', 'jason.derulo@example.com', 'jason', '123-456-7814', 'Address 25', 'buyer');

COMMIT;

BEGIN;

INSERT INTO companies (brand, bank_account, users_id)
VALUES ('ИП Ariana Grande', '111222333', (SELECT id FROM users WHERE first_name = 'Ariana' AND last_name = 'Grande')),
       ('ИП Travis Scott', '111222334', (SELECT id FROM users WHERE first_name = 'Travis' AND last_name = 'Scott')),
       ('ИП Kanye West', '111222335', (SELECT id FROM users WHERE first_name = 'Kanye' AND last_name = 'West')),
       ('ИП Rihanna Unknown', '111222336', (SELECT id FROM users WHERE first_name = 'Rihanna' AND last_name IS NULL)),
       ('ИП Billie Eilish', '111222337', (SELECT id FROM users WHERE first_name = 'Billie' AND last_name = 'Eilish')),
       ('ИП Nicki Minaj', '111222338', (SELECT id FROM users WHERE first_name = 'Nicki' AND last_name = 'Minaj')),
       ('ИП Imagine Dragons', '111222339',
        (SELECT id FROM users WHERE first_name = 'Imagine' AND last_name = 'Dragons')),
       ('ИП Lil Wayne', '111222340', (SELECT id FROM users WHERE first_name = 'Lil' AND last_name = 'Wayne')),
       ('ИП Daddy Yankee', '111222341', (SELECT id FROM users WHERE first_name = 'Daddy' AND last_name = 'Yankee')),
       ('ИП Anuel AA', '111222342', (SELECT id FROM users WHERE first_name = 'Anuel' AND last_name = 'AA')),
       ('ИП Lil Baby', '111222343', (SELECT id FROM users WHERE first_name = 'Lil' AND last_name = 'Baby')),
       ('ИП Beyoncé Unknown', '111222344', (SELECT id FROM users WHERE first_name = 'Beyoncé' AND last_name IS NULL)),
       ('ИП Tyler The Creator', '111222345',
        (SELECT id FROM users WHERE first_name = 'Tyler' AND last_name = 'The Creator')),
       ('ИП Snoop Dogg', '111222346', (SELECT id FROM users WHERE first_name = 'Snoop' AND last_name = 'Dogg')),
       ('ИП Trippie Redd', '111222347', (SELECT id FROM users WHERE first_name = 'Trippie' AND last_name = 'Redd')),
       ('ИП Playboi Carti', '111222348', (SELECT id FROM users WHERE first_name = 'Playboi' AND last_name = 'Carti')),
       ('ИП Arcángel Unknown', '111222349', (SELECT id FROM users WHERE first_name = 'Arcángel' AND last_name IS NULL)),
       ('ИП Jason Derulo', '111222350', (SELECT id FROM users WHERE first_name = 'Jason' AND last_name = 'Derulo'));

COMMIT;

BEGIN;

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Смартфон', 'Современный смартфон с множеством функций', 50, 599.99,
        (SELECT id FROM companies WHERE brand = 'ИП Ariana Grande'), 'available'),
       ('Наушники', 'Беспроводные наушники высокого качества', 100, 199.99,
        (SELECT id FROM companies WHERE brand = 'ИП Ariana Grande'), 'available'),
       ('Ноутбук', 'Легкий и мощный ноутбук', 30, 899.99, (SELECT id FROM companies WHERE brand = 'ИП Ariana Grande'),
        'available'),
       ('Часы', 'Умные часы с различными функциями для спорта', 20, 149.99,
        (SELECT id FROM companies WHERE brand = 'ИП Ariana Grande'), 'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Куртка', 'Зимняя теплая куртка', 40, 199.99, (SELECT id FROM companies WHERE brand = 'ИП Travis Scott'),
        'available'),
       ('Кроссовки', 'Удобные кроссовки для повседневной носки', 60, 89.99,
        (SELECT id FROM companies WHERE brand = 'ИП Travis Scott'), 'available'),
       ('Футболка', 'Стильная футболка из качественного материала', 100, 29.99,
        (SELECT id FROM companies WHERE brand = 'ИП Travis Scott'), 'available'),
       ('Брюки', 'Повседневные брюки из джинсовой ткани', 50, 49.99,
        (SELECT id FROM companies WHERE brand = 'ИП Travis Scott'), 'available'),
       ('Шапка', 'Теплая зимняя шапка', 80, 19.99, (SELECT id FROM companies WHERE brand = 'ИП Travis Scott'),
        'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Спортивные штаны', 'Удобные спортивные штаны для тренировок', 40, 59.99,
        (SELECT id FROM companies WHERE brand = 'ИП Kanye West'), 'available'),
       ('Футболка', 'Футболка для занятий спортом', 80, 29.99, (SELECT id FROM companies WHERE brand = 'ИП Kanye West'),
        'available'),
       ('Кроссовки', 'Легкие и удобные кроссовки', 60, 89.99, (SELECT id FROM companies WHERE brand = 'ИП Kanye West'),
        'available'),
       ('Шорты', 'Спортивные шорты для бега', 30, 39.99, (SELECT id FROM companies WHERE brand = 'ИП Kanye West'),
        'available'),
       ('Спортивная куртка', 'Куртка для занятий на открытом воздухе', 25, 99.99,
        (SELECT id FROM companies WHERE brand = 'ИП Kanye West'), 'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Косметический набор', 'Набор косметики для ежедневного использования', 50, 79.99,
        (SELECT id FROM companies WHERE brand = 'ИП Rihanna Unknown'), 'available'),
       ('Тушь для ресниц', 'Долговечная тушь для ресниц', 100, 19.99,
        (SELECT id FROM companies WHERE brand = 'ИП Rihanna Unknown'), 'available'),
       ('Помада', 'Губная помада ярких цветов', 80, 14.99,
        (SELECT id FROM companies WHERE brand = 'ИП Rihanna Unknown'), 'available'),
       ('Тональный крем', 'Крем для идеального тона кожи', 40, 24.99,
        (SELECT id FROM companies WHERE brand = 'ИП Rihanna Unknown'), 'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Электрогитара', 'Электрогитара с отличным звучанием', 10, 499.99,
        (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish'), 'available'),
       ('Акустическая гитара', 'Классическая акустическая гитара', 15, 299.99,
        (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish'), 'available'),
       ('Микрофон', 'Профессиональный микрофон для студийных записей', 20, 199.99,
        (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish'), 'available'),
       ('Колонки', 'Колонки для профессионального звучания', 25, 149.99,
        (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish'), 'available'),
       ('Наушники', 'Профессиональные наушники для студийной работы', 50, 99.99,
        (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish'), 'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Сумка', 'Модная женская сумка', 50, 199.99, (SELECT id FROM companies WHERE brand = 'ИП Nicki Minaj'),
        'available'),
       ('Туфли', 'Женские туфли высокого качества', 30, 99.99,
        (SELECT id FROM companies WHERE brand = 'ИП Nicki Minaj'), 'available'),
       ('Платье', 'Стильное платье для вечернего выхода', 20, 149.99,
        (SELECT id FROM companies WHERE brand = 'ИП Nicki Minaj'), 'available'),
       ('Куртка', 'Модная женская куртка', 15, 299.99, (SELECT id FROM companies WHERE brand = 'ИП Nicki Minaj'),
        'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Гитара', 'Качественная акустическая гитара', 25, 349.99,
        (SELECT id FROM companies WHERE brand = 'ИП Imagine Dragons'), 'available'),
       ('Барабан', 'Большой барабан для профессиональной игры', 10, 499.99,
        (SELECT id FROM companies WHERE brand = 'ИП Imagine Dragons'), 'available'),
       ('Микрофон', 'Студийный микрофон для записи', 20, 249.99,
        (SELECT id FROM companies WHERE brand = 'ИП Imagine Dragons'), 'available'),
       ('Колонки', 'Профессиональные колонки для живых выступлений', 15, 299.99,
        (SELECT id FROM companies WHERE brand = 'ИП Imagine Dragons'), 'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Кепка', 'Модная кепка для повседневной носки', 100, 29.99,
        (SELECT id FROM companies WHERE brand = 'ИП Lil Wayne'), 'available'),
       ('Толстовка', 'Толстовка с логотипом', 50, 69.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Wayne'),
        'available'),
       ('Кроссовки', 'Легкие кроссовки', 40, 89.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Wayne'),
        'available'),
       ('Штаны', 'Спортивные штаны для тренировок', 30, 59.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Wayne'),
        'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Парфюм', 'Парфюм высокого качества', 50, 149.99, (SELECT id FROM companies WHERE brand = 'ИП Daddy Yankee'),
        'available'),
       ('Футболка', 'Стильная футболка', 80, 39.99, (SELECT id FROM companies WHERE brand = 'ИП Daddy Yankee'),
        'available'),
       ('Кепка', 'Кепка для повседневного использования', 100, 24.99,
        (SELECT id FROM companies WHERE brand = 'ИП Daddy Yankee'), 'available'),
       ('Часы', 'Модные часы', 30, 249.99, (SELECT id FROM companies WHERE brand = 'ИП Daddy Yankee'), 'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Кроссовки', 'Кроссовки для бега', 60, 99.99, (SELECT id FROM companies WHERE brand = 'ИП Anuel AA'),
        'available'),
       ('Штаны', 'Спортивные штаны', 40, 59.99, (SELECT id FROM companies WHERE brand = 'ИП Anuel AA'), 'available'),
       ('Куртка', 'Легкая куртка для тренировок', 30, 149.99, (SELECT id FROM companies WHERE brand = 'ИП Anuel AA'),
        'available'),
       ('Шорты', 'Шорты для тренировок', 80, 39.99, (SELECT id FROM companies WHERE brand = 'ИП Anuel AA'),
        'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Кроссовки', 'Спортивные кроссовки', 60, 89.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Baby'),
        'available'),
       ('Шапка', 'Зимняя шапка', 80, 19.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Baby'), 'available'),
       ('Футболка', 'Футболка с логотипом', 100, 29.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Baby'),
        'available'),
       ('Куртка', 'Куртка для занятий спортом', 50, 149.99, (SELECT id FROM companies WHERE brand = 'ИП Lil Baby'),
        'available');

INSERT INTO products (name, description, stock_quantity, price, company_id, status)
VALUES ('Платье', 'Элегантное вечернее платье', 20, 299.99,
        (SELECT id FROM companies WHERE brand = 'ИП Beyoncé Unknown'), 'available'),
       ('Туфли', 'Модные туфли', 30, 149.99, (SELECT id FROM companies WHERE brand = 'ИП Beyoncé Unknown'),
        'available'),
       ('Сумка', 'Стильная женская сумка', 40, 199.99, (SELECT id FROM companies WHERE brand = 'ИП Beyoncé Unknown'),
        'available'),
       ('Шарф', 'Модный шарф', 60, 49.99, (SELECT id FROM companies WHERE brand = 'ИП Beyoncé Unknown'), 'available');


COMMIT;

BEGIN;

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Drake'), (SELECT id
                                                            FROM products
                                                            WHERE name = 'Смартфон'
                                                              AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Ariana Grande')),
        1, 599.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Drake'), (SELECT id
                                                            FROM products
                                                            WHERE name = 'Ноутбук'
                                                              AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Ariana Grande')),
        1, 899.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Taylor' AND last_name = 'Swift'), (SELECT id
                                                                                     FROM products
                                                                                     WHERE name = 'Куртка'
                                                                                       AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Travis Scott')),
        2, 399.98, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Bad' AND last_name = 'Bunny'), (SELECT id
                                                                                  FROM products
                                                                                  WHERE name = 'Кроссовки'
                                                                                    AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Travis Scott')),
        1, 89.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Bad' AND last_name = 'Bunny'), (SELECT id
                                                                                  FROM products
                                                                                  WHERE name = 'Шапка'
                                                                                    AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Travis Scott')),
        1, 19.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'The' AND last_name = 'Weeknd'), (SELECT id
                                                                                   FROM products
                                                                                   WHERE name = 'Футболка'
                                                                                     AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Travis Scott')),
        3, 89.97, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Justin' AND last_name = 'Bieber'), (SELECT id
                                                                                      FROM products
                                                                                      WHERE name = 'Кроссовки'
                                                                                        AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Kanye West')),
        1, 89.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Justin' AND last_name = 'Bieber'), (SELECT id
                                                                                      FROM products
                                                                                      WHERE name = 'Футболка'
                                                                                        AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Kanye West')),
        2, 59.98, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Ed' AND last_name = 'Sheeran'), (SELECT id
                                                                                   FROM products
                                                                                   WHERE name = 'Кроссовки'
                                                                                     AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Kanye West')),
        1, 89.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Ed' AND last_name = 'Sheeran'), (SELECT id
                                                                                   FROM products
                                                                                   WHERE name = 'Спортивная куртка'
                                                                                     AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Kanye West')),
        1, 99.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Eminem'), (SELECT id
                                                             FROM products
                                                             WHERE name = 'Косметический набор'
                                                               AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Rihanna Unknown')),
        1, 79.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Eminem'), (SELECT id
                                                             FROM products
                                                             WHERE name = 'Помада'
                                                               AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Rihanna Unknown')),
        1, 14.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Ariana' AND last_name = 'Grande'), (SELECT id
                                                                                      FROM products
                                                                                      WHERE name = 'Электрогитара'
                                                                                        AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish')),
        1, 499.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Travis' AND last_name = 'Scott'), (SELECT id
                                                                                     FROM products
                                                                                     WHERE name = 'Наушники'
                                                                                       AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish')),
        2, 199.98, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Nicki' AND last_name = 'Minaj'), (SELECT id
                                                                                    FROM products
                                                                                    WHERE name = 'Сумка'
                                                                                      AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Nicki Minaj')),
        1, 199.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Nicki' AND last_name = 'Minaj'), (SELECT id
                                                                                    FROM products
                                                                                    WHERE name = 'Платье'
                                                                                      AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Nicki Minaj')),
        1, 149.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Billie' AND last_name = 'Eilish'), (SELECT id
                                                                                      FROM products
                                                                                      WHERE name = 'Микрофон'
                                                                                        AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Billie Eilish')),
        1, 199.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Lil' AND last_name = 'Wayne'), (SELECT id
                                                                                  FROM products
                                                                                  WHERE name = 'Толстовка'
                                                                                    AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Lil Wayne')),
        1, 69.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Lil' AND last_name = 'Wayne'), (SELECT id
                                                                                  FROM products
                                                                                  WHERE name = 'Штаны'
                                                                                    AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Lil Wayne')),
        1, 59.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Daddy' AND last_name = 'Yankee'), (SELECT id
                                                                                     FROM products
                                                                                     WHERE name = 'Парфюм'
                                                                                       AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Daddy Yankee')),
        2, 299.98, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Anuel' AND last_name = 'AA'), (SELECT id
                                                                                 FROM products
                                                                                 WHERE name = 'Кроссовки'
                                                                                   AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Anuel AA')),
        1, 99.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Anuel' AND last_name = 'AA'), (SELECT id
                                                                                 FROM products
                                                                                 WHERE name = 'Шорты'
                                                                                   AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Anuel AA')),
        1, 39.99, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Lil' AND last_name = 'Baby'), (SELECT id
                                                                                 FROM products
                                                                                 WHERE name = 'Кроссовки'
                                                                                   AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Lil Baby')),
        1, 89.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Lil' AND last_name = 'Baby'), (SELECT id
                                                                                 FROM products
                                                                                 WHERE name = 'Футболка'
                                                                                   AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Lil Baby')),
        2, 59.98, 'confirmed');

INSERT INTO orders (user_id, product_id, quantity, total_price, status)
VALUES ((SELECT id FROM users WHERE first_name = 'Beyoncé'), (SELECT id
                                                              FROM products
                                                              WHERE name = 'Платье'
                                                                AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Beyoncé Unknown')),
        1, 299.99, 'confirmed'),
       ((SELECT id FROM users WHERE first_name = 'Beyoncé'), (SELECT id
                                                              FROM products
                                                              WHERE name = 'Туфли'
                                                                AND company_id = (SELECT id FROM companies WHERE brand = 'ИП Beyoncé Unknown')),
        1, 149.99, 'confirmed');


COMMIT;

