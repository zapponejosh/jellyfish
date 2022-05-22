import React from 'react';

import styles from './TileGrid.module.css';
import ProjectTile from '../ProjectTile/ProjectTile';

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

function TileGrid({ projects }: { projects: ProjectProps[] }): JSX.Element {
  const projectList = projects.map((p) => {
    return <ProjectTile key={p.ID} {...p} />;
  });
  return <div className={styles.gridContainer}>{projectList}</div>;
}

export default TileGrid;
