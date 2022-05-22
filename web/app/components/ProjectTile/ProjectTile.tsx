import React from 'react';
import { Link } from 'react-router-dom';
import Button from '../Button/Button';
import styles from './ProjectTile.module.css';

type ProjectProps = {
  Name: string;
  ID: number;
  Admin: string;
  AdminID: number;
  Institution: string;
  InstitutionID: number;
  Category: string;
  CategoryID: number;
};

function ProjectTile(project: ProjectProps): JSX.Element {
  return (
    <div className={styles.tile}>
      <div>
        <h3>{project.Name}</h3>
        <p className={styles.cat}>
          <span>{project.Category}</span>
        </p>
        <p className={styles.inst}>{project.Institution}</p>
      </div>
      <Link to={'/project/' + project.ID}>
        <Button>Learn More</Button>
      </Link>
    </div>
  );
}

export default ProjectTile;
