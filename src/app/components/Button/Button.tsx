import React, { ReactNode } from 'react';
import classes from './Button.module.css';

export interface ButtonProps extends React.ComponentPropsWithoutRef<'button'> {
  children: ReactNode;
}
function Button({ children, ...rest }: ButtonProps): JSX.Element {
  return (
    <button {...rest} className={classes.button}>
      {children}
    </button>
  );
}

export default Button;
