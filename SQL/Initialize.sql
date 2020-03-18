DROP TABLE APP_CONTACTS;
DROP TABLE APP_USER;
DROP TABLE APP_PICTURES;

CREATE TABLE APP_PICTURES (
	ID INTEGER,
	UPLOADED_AT DATETIME DEFAULT CURRENT_TIMESTAMP,
	NAME VARCHAR(255),
	UPLOADED_BY VARCHAR(255)
);

CREATE TABLE APP_USER (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	CREATED_AT DATETIME DEFAULT CURRENT_TIMESTAMP,
	LAST_ACTION BIGINT DEFAULT CURRENT_TIMESTAMP,
	FIRSTNAME VARCHAR(255),
	USERNAME VARCHAR(255),
	GENDER INTEGER,
	GENDER_WANTED INTEGER,
	COUNTRY VARCHAR(255),
	AGE INTEGER,
	AGE_FROM INTEGER,
	AGE_TO INTEGER,
	PICTURE VARCHAR(255),
	LANG VARCHAR(255),
	PREMIUM BOOL,
	ADULT BOOL,
	PASSWORD VARCHAR(255),
	LATEST_TOKEN VARCHAR(255)
);

CREATE TABLE APP_CONTACTS (
	INITIATOR INTEGER,
	FRIEND INTEGER,
	IS_FRIENDLY BOOL,
	FRIEND_LEVEL INTEGER,
	PLAYABLE BOOL,
	PROGRESS INTEGER,
	PRIMARY KEY (INITIATOR, FRIEND),
	FOREIGN KEY(INITIATOR) REFERENCES APP_USER(ID),
	FOREIGN KEY(FRIEND) REFERENCES APP_USER(ID)
);