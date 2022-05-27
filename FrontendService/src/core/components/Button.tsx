import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../Context";

const useStyles = createUseStyles((theme: Theme) => ({
  button: {
    cursor: "pointer",
    padding: "1em",
    color: theme.colors.textLight,
    backgroundColor: theme.colors.primary,
    borderRadius: "15px",
    border: 0,
    fontWeight: 600,
    fontSize: "0.875 rem",
    "&:hover": {
      opacity: 0.8,
    },
    "&:active": {
      paddingLeft: "0.97em",
      paddingRight: "0.97em",
    },
  },
}));

type Props = {
  type?: 'button'|'submit';
  className?: string;
  onClick?: () => any;
};

export const Button: FC<Props> = ({
  children,
  type = "button",
  className = "",
  onClick = () => {},
}) => {
  const classes = useStyles();

  return (
    <button
      type={type}
      className={`${classes.button} ${className}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
};
