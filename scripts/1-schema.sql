CREATE TABLE dummies(
   id serial PRIMARY KEY,
   some_field VARCHAR (50) NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   deleted_at TIMESTAMP
);
