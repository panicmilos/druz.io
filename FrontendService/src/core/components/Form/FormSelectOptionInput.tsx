/* eslint-disable react-hooks/exhaustive-deps */
import { FC, useContext, useEffect } from "react";
import { Modify } from "../../utils/types";
import { Props as SelectOptionProps, SelectOptionInput } from "../Inputs/SelectOptionInput";
import { FormContext } from "./Form";

type Props = Modify<SelectOptionProps, {
  name: string;
  label?: string;
  value?: string;
  onChange?: (v: string) => any;
}>

export const FormSelectOptionInput: FC<Props> = (props) => {

  const { value, setValue, errors, showErrors } = useContext(FormContext);

  const setFieldValue = (v: string) => setValue((prevValue: any) => ({...prevValue, [props.name]: v}));

  useEffect(() => {
    if(!value[props.name]) {
      setFieldValue("");
    }
  }, []);

  const fieldValue = value[props.name] ?? "";
  const errorMessage = errors[props.name];

  return (
    <>
      <SelectOptionInput
        {...props}
        label={props.label ?? props.name}
        value={fieldValue}
        onChange={v => {
          props.onChange && props.onChange(v);
          setFieldValue(v);
        }}
      />
      {showErrors && errorMessage && <p>{errorMessage}</p>}
    </>
  );
}