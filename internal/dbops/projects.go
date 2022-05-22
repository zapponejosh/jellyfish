package dbops

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/zapponejosh/jellyfish/internal/models"
)

func checkDate(end pgtype.Date) interface{} {
	if end.Status != 2 {
		return nil
	}
	return end
}
func (d DB) CreateProject(ctx context.Context, project *models.Project) (int, error) {
	lastInsertId := 0

	row := d.db.QueryRow(ctx, `INSERT INTO projects(
		name,
		description,
		start_date,
		end_date,
		category_id,
		admin_id,
		institution_id
)
VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
) RETURNING id`, project.Name, project.Description, project.Start_date, checkDate(project.End_date), project.Category, project.Admin, project.Institution) // using queryRow to get ID since postgres doesn't support lastInsertId
	err := row.Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (d DB) GetProject(ctx context.Context, id int) (*models.GetProject, error) {
	var project models.GetProject
	row := d.db.QueryRow(ctx, `select 
	projects.*,
	users.name AS admin,
	institutions.name AS institution,
	categories.name AS category
from projects
	INNER JOIN users ON admin_id = users.id
	INNER JOIN institutions ON projects.institution_id = institutions.id
	INNER JOIN categories ON projects.category_id = categories.id
	WHERE projects.id = $1;`, id)
	err := row.Scan(&project.ID, &project.AdminID, &project.Start_date, &project.End_date, &project.InstitutionID, &project.Created_at, &project.Updated_at, &project.CategoryID, &project.Name, &project.Description, &project.Admin, &project.Institution, &project.Category)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (d DB) ListProjects(ctx context.Context) ([]*models.ProjectPreview, error) {
	var projects []*models.ProjectPreview
	rows, err := d.db.Query(ctx, `select 
	projects.name AS project,
	projects.id as project_id,
	users.name AS admin,
	users.id AS admin_id,
	institutions.name AS institution,
	institutions.id AS institution_id,
	categories.name AS category,
	categories.id AS category_id
from projects
	INNER JOIN users ON admin_id = users.id
	INNER JOIN institutions ON projects.institution_id = institutions.id
	INNER JOIN categories ON projects.category_id = categories.id;`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p models.ProjectPreview
		err = rows.Scan(&p.Name, &p.ID, &p.Admin, &p.AdminID, &p.Institution, &p.InstitutionID, &p.Category, &p.CategoryID)
		if err != nil {
			return nil, err
			// should I break instead?
		}
		projects = append(projects, &p)
	}

	return projects, nil
}
