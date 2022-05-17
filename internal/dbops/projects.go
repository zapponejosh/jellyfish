package dbops

import (
	"context"

	"github.com/zapponejosh/jellyfish/internal/models"
)

// func (d DB) GetProject(ctx context.Context, id int) (*models.Project, error) {
// 	var project models.Project
// 	row := d.db.QueryRow(ctx, "SELECT * FROM projects WHERE id = $1", id)
// 	err := row.Scan(&project.)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &project, nil
// }

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
