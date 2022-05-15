DROP TABLE IF EXISTS users,
categories,
institutions,
projects;
CREATE TABLE users(
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    profile_image VARCHAR(255) NOT NULL,
    UNIQUE(email),
    PRIMARY KEY(id)
);
CREATE TABLE categories(
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    UNIQUE(name),
    PRIMARY KEY(id)
);
CREATE TABLE institutions(
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    website VARCHAR(255) NULL,
    email VARCHAR(255) NOT NULL,
    owner_id INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    UNIQUE(email, name),
    PRIMARY KEY(id),
    CONSTRAINT institutions_owner_id_foreign FOREIGN KEY(owner_id) REFERENCES users(id)
);
CREATE TABLE projects(
    id INT GENERATED ALWAYS AS IDENTITY,
    admin_id INTEGER NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    institution_id INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    category_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(500) NOT NULL,
    UNIQUE(name),
    PRIMARY KEY(id),
    CONSTRAINT projects_admin_id_foreign FOREIGN KEY(admin_id) REFERENCES users(id),
    CONSTRAINT projects_category_id_foreign FOREIGN KEY(category_id) REFERENCES categories(id),
    CONSTRAINT projects_institution_id_foreign FOREIGN KEY(institution_id) REFERENCES institutions(id) ON DELETE CASCADE
);
CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- Update time triggers
CREATE TRIGGER set_timestamp_users BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
CREATE TRIGGER set_timestamp_projects BEFORE
UPDATE ON projects FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
CREATE TRIGGER set_timestamp_institutions BEFORE
UPDATE ON institutions FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
-- CREATE USERS
INSERT INTO users(name, email, password, profile_image)
VALUES (
        'Josh',
        'test1@email.com',
        'password',
        'https://picsum.photos/200'
    ),
    (
        'Jimmy',
        'test2@email.com',
        'password',
        'https://picsum.photos/200'
    ),
    (
        'Tex',
        'test3@email.com',
        'password',
        'https://picsum.photos/200'
    );
-- CREATE INSTITUTIONS
INSERT INTO institutions(name, website, email, owner_id)
VALUES (
        'Ins1 - Smithsonian',
        'ins1.com',
        'ins1@email.com',
        1
    ),
    (
        'Ins2 - Babies R Us',
        'ins2.com',
        'ins2@email.com',
        2
    ),
    ('Ins3 - IHOP', 'ins3.com', 'ins3@email.com', 3);
-- CREATE CATEGORIES
INSERT INTO categories(name)
VALUES ('Cat1 - Sciene'),
    ('Cat2 - Literature'),
    ('Cat3 - Dogs');
-- CREATE PROJECTS
INSERT INTO projects(
        name,
        description,
        start_date,
        category_id,
        admin_id,
        institution_id
    )
VALUES (
        'DoinStuff1',
        'Seed 1 - A project to do some stuff!',
        '2022-01-01',
        1,
        1,
        1
    ),
    (
        'DoinStuff2',
        'Seed 2 - A project to do some stuff!',
        '2022-02-02',
        2,
        2,
        2
    ),
    (
        'DoinStuff3',
        'Seed 3 - A project to do some stuff!',
        '2022-03-03',
        3,
        3,
        3
    );