import { FC, useContext, useEffect } from "react";
import { Modify } from "../../utils/types";
import { Props as TextAreaProps, TextAreaInput } from "../Inputs/TextAreaInput"
import { FormContext } from "./Form";

type Props = Modify<TextAreaProps, {
  name: string;
  value?: string;
  onChange?: (v: string) => any;
}>

export const FormTextAreaInput: FC<Props> = (props) => {

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
      <TextAreaInput
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