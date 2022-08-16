/* eslint-disable react-hooks/exhaustive-deps */
import { FC, useContext, useEffect } from "react";
import { Modify } from "../../utils/types";
import {Props as DateInputProps, DateInput} from '../Inputs/DateInput';
import { FormContext } from "./Form";

type Props = Modify<DateInputProps, {
  name: string;
  value?: string;
  onChange?: (v: string) => any;
}>

export const FormDateInput: FC<Props> = (props) => {

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
      <DateInput
        {...props}
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