import React from 'react';
import styles from './App.module.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import Nav from './components/Nav/Nav';
import Project from './pages/Project';

function App(): JSX.Element {
  return (
    <Router>
      <div className={styles.App}>
        <header>
          <Nav />
        </header>
        <Routes>
          <Route path="/about" element={<About />} />
          <Route path="/" element={<Home />} />
          <Route path="/project/:projectId" element={<Project />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
