import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
// import TileGrid from '../components/TileGrid/TileGrid';
// import Welcome from '../components/Welcome/Welcome';

type ProjectDetails = {
  ID: number;
  AdminID: number;
  Admin: string;
  Start_date: string;
  End_date: string;
  InstitutionID: number;
  Institution: string;
  Created_at: string;
  Updated_at: string;
  CategoryID: number;
  Category: string;
  Name: string;
  Description: string;
  error: string; // I need to not do this...
};

function Project(): JSX.Element {
  const [project, setProject] = useState<ProjectDetails | null>(null);
  const [error, setError] = useState<string | null>(null);
  const { projectId } = useParams();

  useEffect(() => {
    fetch(`/api/project/${projectId}`)
      .then((response) => {
        return response.json();
      })
      .then((data: ProjectDetails) => {
        console.log(data);
        if (data.error) {
          setError(data.error);
          return;
        }
        if (data.ID) {
          setProject(data);
        }
      });
  }, []);
  return (
    <main>
      {project ? (
        <>
          <h1 data-project-id={project?.ID}>{project?.Name}</h1>
          <p
            data-inst-id={project?.InstitutionID}
            data-admin-id={project?.AdminID}
          >
            {project?.Admin} | {project?.Institution}
          </p>
          <p>
            <b>Start Date:</b> {project?.Start_date}
          </p>
          {project?.End_date && (
            <p>
              <b>End Date:</b> {project.End_date}
            </p>
          )}
          <p>{project?.Description}</p>
        </>
      ) : (
        <>
          <h4>Error: {error}</h4>
        </>
      )}
    </main>
  );
}

export default Project;
