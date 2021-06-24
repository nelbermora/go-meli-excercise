
CREATE TABLE "users" (
    "id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "password"	TEXT NOT NULL,
    "username"	TEXT NOT NULL
);

CREATE TABLE "rol" (
   "id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
   "description"	TEXT,
   "name"	TEXT NOT NULL
);

CREATE TABLE "user_rol" (
    "id_usuario"	INTEGER NOT NULL,
    "id_rol"	INTEGER NOT NULL,
    PRIMARY KEY("id_usuario","id_rol"),
    FOREIGN KEY(id_usuario) REFERENCES users(id),
    FOREIGN KEY(id_rol) REFERENCES rol(id)
);

CREATE TABLE "products" (
    "id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "description"	TEXT,
    "expiration_rate"	INTEGER,
    "freezing_rate"	INTEGER,
    "height"	INTEGER NOT NULL,
    "lenght"	INTEGER NOT NULL,
    "netweight"	INTEGER NOT NULL,
    "product_code"	TEXT NOT NULL,
    "recommended_freezing_temperature"	INTEGER,
    "width"	INTEGER NOT NULL,
    "id_product_type"	INTEGER,
    "id_seller"	INTEGER
);