import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../../Context";

type Option = { label: string, value: string };

export type Props = {
  values: string[];
  options: Option[];
  onChange: (v: string[]) => any;
  label?: string;
  className?: string;
  disabled?: boolean;
}

const useStyles = createUseStyles((theme: Theme) => ({
  inputContainer: {
    display: 'grid',
    gridTemplateColumns: 'auto auto auto',
  },
  inputElement: {
    padding: '20px',
    textAlign: 'center'
  }
}));

export const MultipleCheckboxInput: FC<Props> = ({
  label,
  className,
  values,
  onChange,
  options,
  disabled = false
}) => {
  const classes = useStyles();

  const toggleValue = (option: string) => {
    if (values.includes(option)) {
      onChange(values.filter(o => o !== option));
    } else {
      onChange([...values, option]);
    }
  }

  return (
    <div>
      {label && <label>{label}</label>}
      <div className={classes.inputContainer}>
        {
          options?.map((o: Option, i: number) =>
            <>
              <div key={i} onClick={() => toggleValue(o.value)} className={classes.inputElement}>
                {o.label && <label>{o.label}</label>}
                <input
                  type="checkbox"
                  className={`${className || ""}`}
                  onChange={() => { }}
                  checked={values.includes(o.value)}
                  disabled={disabled}
                />
              </div>
            </>
          )
        }
      </div>
    </div>
  );
};