-- sample query 
select projects.name AS Project,
    users.name AS Admin,
    institutions.name AS Institution,
    categories.name AS Category
from projects
    INNER JOIN users ON admin_id = users.id
    INNER JOIN institutions ON projects.institution_id = institutions.id
    INNER JOIN categories ON projects.category_id = categories.id;
--DELETE institution deletes projects as well
DELETE FROM institutions
WHERE name = 'Ins2';
-- project 2 is deleted as well
--put them back with Ins3 attached to project
INSERT INTO institutions(name, website, email, owner_id)
VALUES ('Ins2', 'ins2.com', 'ins2@email.com', 2);
INSERT INTO projects(
        name,
        description,
        start_date,
        category_id,
        admin_id,
        institution_id
    )
VALUES (
        'DoinStuff2',
        'Seed 2 - A project to do some stuff!',
        '2022-02-02',
        2,
        2,
        3
    );
--Update project to use Ins2
UPDATE projects
SET institution_id = 5 -- new Ins2
WHERE name = 'DoinStuff2';
-- see that updated_at has updated according to trigger