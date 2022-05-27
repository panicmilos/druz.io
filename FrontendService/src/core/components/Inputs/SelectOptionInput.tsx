import { FC } from 'react';
import { createUseStyles } from "react-jss";
import { Theme } from '../../Context';

type Option = {label: string, value: string|number};

const useStyles = createUseStyles((theme: Theme) => ({
  inputContainer: {
    color: theme.colors.textDark,
    '& label': {
      display: 'flex',
      fontSize: '15px',
      padding: '0.3em',
    },
    "& select": {
      "-webkit-border-radius": "20px",
      "-moz-border-radius": "20px",
      color: theme.colors.textDark,
      borderRadius: "20px",
      border: `1px solid ${theme.colors.secondary}`,
      width: "250px",
      height: "30px",
      paddingLeft: "10px",
    },
    "& select:focus": {
      outline: "none",
      border: `1px solid ${theme.colors.primary}`,
    },
  },
}));

export type Props = {
  label: string;
  value: string;
  onChange: (v: string) => any;
  options: Option[];
  disabled?: boolean;
  hasDefaultOption?: boolean;
}

export const SelectOptionInput: FC<Props> = ({ label, value, onChange, options, disabled = false, hasDefaultOption = true }) => {
  const classes = useStyles();

  const handleChange = (e: any) => {
    onChange(e.target.value);
  }

  const formKey = (i: number, option: Option) => {
    return `${i}_${option.value}_${option.label}`;
  }

  return (
    <div className={classes.inputContainer}>
      {label && <label>{label}</label>}
      <select 
        value={(value !== null && value !== undefined) ? value : ""}
        disabled={disabled}
        onChange={handleChange}
      >
        {hasDefaultOption && <option value={''}>{label}</option>}
        {options && options.map((option, i) => <option key={formKey(i, option)} value={option.value}>{option.label}</option>)}
      </select>
    </div>
  );
}