import React from 'react';
import { Link } from 'react-router-dom';
import Button from '../Button/Button';
import styles from './Nav.module.css';

function Nav(): JSX.Element {
  return (
    <nav className={styles.nav}>
      <h3>Jellyfish</h3>
      <div className={styles.menu}>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/about">About</Link>
          </li>
        </ul>
        <Link to="/">
          <Button>Create project</Button>
        </Link>
      </div>
    </nav>
  );
}

export default Nav;
