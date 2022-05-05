import React, { useEffect, useState } from 'react';
import classes from './Welcome.module.css';

function Welcome(): JSX.Element {
  const [message, setMessage] = useState('');
  const [dataStr, setDataStr] = useState('');

  useEffect(() => {
    fetch('/api/hello')
      .then((response) => response.json())
      .then((data) => {
        setDataStr(data.dataTest);
        setMessage(data.message);
      });
  }, []);

  return (
    <>
      <p className={classes.message}>{message}</p>
      <p>{dataStr || 'FAIL'}</p>
    </>
  );
}

export default Welcome;
