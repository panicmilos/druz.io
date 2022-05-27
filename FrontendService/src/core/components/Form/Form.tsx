/* eslint-disable react-hooks/exhaustive-deps */
import { createContext, FC, useEffect, useState } from "react";
import * as Yup from 'yup';

type FormContextValue = {
  value: any;
  setValue: (v: any) => any;
  errors: any;
  showErrors: boolean;
}

export const FormContext = createContext<FormContextValue>({
  errors: {},
  showErrors: false,
  value: {},
  setValue: () => {}
});

type Props = {
  initialValue?: any;
  schema: Yup.AnyObjectSchema | undefined;
  onSubmit?: (value: any) => any;
}

export const Form: FC<Props> = ({ children, initialValue = {}, schema, onSubmit = () => {} }) => {
  const [value, setValue] = useState(initialValue);

  const [errors, setErrors] = useState({});
  const [hasSubmitted, setHasSubmitted] = useState(false);


  const validate = (v: any, submitted: boolean) => {
    if(!submitted) return;
    setErrors({});
    return schema?.validate(v, {abortEarly: false})
      .then(() => {
        return true;
      })
      .catch(({errors}) => {
        errors.forEach((err: any) => {
          setErrors(prevErrors => ({...prevErrors, ...err}));
        });
        return false;
      }) ?? true;
  };

  useEffect(() => {
    validate(value, hasSubmitted);
  }, [value, hasSubmitted]);

  return (
    <FormContext.Provider value={{ value, setValue, errors, showErrors: hasSubmitted }}>
      <form onSubmit={async (e) => {
        e.preventDefault();
        setHasSubmitted(true);
        const valid = await validate(value, true);
        if(valid) onSubmit(value);
      }}>
        {children}
      </form>
    </FormContext.Provider>
  )
}