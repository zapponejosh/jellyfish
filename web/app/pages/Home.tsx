import React, { useEffect, useState } from 'react';
import TileGrid from '../components/TileGrid/TileGrid';

type Project = {
  Name: string;
  ID: number;
  Admin: string;
  AdminID: number;
  Institution: string;
  InstitutionID: number;
  Category: string;
  CategoryID: number;
};

function Home(): JSX.Element {
  const [projects, setProjects] = useState<Project[] | null>(null);

  useEffect(() => {
    fetch('/api/project')
      .then((response) => response.json())
      .then((data: Project[]) => {
        setProjects(data);
      });
  }, []);
  return (
    <main>
      <h1>Home page :dancing-panda:</h1>
      {/* <Welcome /> */}
      <section>{projects && <TileGrid projects={projects} />}</section>
    </main>
  );
}

export default Home;
