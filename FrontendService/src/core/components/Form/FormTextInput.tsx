/* eslint-disable react-hooks/exhaustive-deps */
import { FC, useContext, useEffect } from "react";
import { Modify } from "../../utils/types";
import {Props as TextInputProps, TextInput} from '../Inputs/TextInput';
import { FormContext } from "./Form";

type Props = Modify<TextInputProps, {
  name: string;
  value?: string;
  onChange?: (v: string) => any;
}>

export const FormTextInput: FC<Props> = (props) => {

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
      <TextInput
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