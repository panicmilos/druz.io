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
    "& textarea": {
      width: '100%',
      "-webkit-border-radius": "20px",
      "-moz-border-radius": "20px",
      color: theme.colors.textDark,
      borderRadius: "20px",
      border: `1px solid ${theme.colors.secondary}`,
      paddingLeft: "10px",
    },
    "& textarea:focus": {
      outline: "none",
      border: `1px solid ${theme.colors.primary}`,
    },
  },
}));

export type Props = {
  value: string;
  onChange: (v: string) => any;
  label?: string;
  rows?: number;
  cols?: number;
  className?: string;
  placeholder?: string;
  disabled?: boolean;
};

export const TextAreaInput: FC<Props> = ({
  label,
  className,
  placeholder,
  value,
  onChange,
  disabled = false,
  rows = 4,
  cols = 50
}) => {

  const classes = useStyles();

  const handleChange = (e: any) => {
    onChange(e.target.value);
  };

  return (
    <div className={classes.inputContainer}>
      {label && <label>{label}</label>}
      <textarea
        rows={rows}
        cols={cols}
        className={`${className || ""}`}
        placeholder={placeholder || label}
        value={value != null ? value : ""}
        disabled={disabled}
        onChange={handleChange}
      />
    </div>
  );

};