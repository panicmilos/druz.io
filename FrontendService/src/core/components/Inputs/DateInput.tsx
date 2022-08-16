import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../../Context";

const useStyles = createUseStyles((theme: Theme) => ({
  inputContainer: {
    color: theme.colors.textDark,
    paddingLeft: "0.7em",
    paddingRight: "0.7em",
    '& label': {
      display: 'flex',
      fontSize: '15px',
      padding: '0.3em',
    },
    "& input[type=date], & input[type=datetime-local]": {
      "-webkit-border-radius": "20px",
      "-moz-border-radius": "20px",
      color: theme.colors.textDark,
      borderRadius: "20px",
      border: `1px solid ${theme.colors.secondary}`,
      width: "100%",
      height: "30px",
      paddingLeft: "10px",
    },
    "& input[type=date]:focus, & input[type=datetime-local]:focus": {
      outline: "none",
      border: `1px solid ${theme.colors.primary}`,
    },
  },
}));

export type Props = {
  value: string;
  onChange: (v: string) => any;
  label?: string;
  type?: string;
  className?: string;
  placeholder?: string;
  disabled?: boolean;
};

export const DateInput: FC<Props> = ({
  label,
  type = "date",
  className,
  placeholder,
  value,
  onChange,
  disabled = false,
}) => {
  const classes = useStyles();

  const handleChange = (e: any) => {
    onChange(e.target.value);
  };

  return (
    <div className={classes.inputContainer}>
      {label && <label>{label}</label>}
      <input
        type={type}
        className={`${className || ""}`}
        placeholder={placeholder || label}
        value={value != null ? value : ""}
        disabled={disabled}
        onChange={handleChange}
      />
    </div>
  );
};