CREATE TABLE "user" (
                         "id" serial PRIMARY KEY,
                         "username" text NOT NULL,
                         "email" text NOT NULL UNIQUE ,
                       "password" text NOT NULL,
                         "isActivated" bool DEFAULT (false),
                         "isActivationLink" text NOT NULL DEFAULT('')

);
