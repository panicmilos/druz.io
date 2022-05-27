import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../../Context";

export type Props = {
  value: boolean;
  onChange: (v: boolean) => any;
  label?: string;
  className?: string;
  disabled?: boolean;
}

const useStyles = createUseStyles((theme: Theme) => ({
  inputContainer: {
  },
}));

export const CheckboxInput: FC<Props> = ({
  label,
  className,
  value,
  onChange,
  disabled = false
}) => {
  const classes = useStyles();

  const handleChange = () => {
    onChange(!value)
  }

  return (
    <div className={classes.inputContainer}>
      <div onClick={handleChange}>
        {label && <label>{label}</label>}
        <input
          type="checkbox"
          className={`${className || ""}`}
          onChange={() => {}}
          checked={value}
          disabled={disabled}
        />
      </div>
    </div>   
  );
}