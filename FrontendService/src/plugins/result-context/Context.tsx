import { createContext, FC, useContext, useState } from "react"
import debounce from 'lodash.debounce';

export type ResultType = {
  status: 'START'|'OK'|'ERROR'
  type: string
}

export type ResultContextType = {
  result: any
  setResult: (r: ResultType|undefined) => any
}

const initialResultValue = {
  result: undefined,
  setResult: (_: ResultType|undefined) => {}
}

export const ResultContext = createContext<ResultContextType>(initialResultValue);

export const useResult = (scope: string) => {

  const { result, setResult } = useContext(ResultContext);

  const setScopedResult = (r: ResultType|undefined) => setResult({...result, [scope]: r });

  if (!Object.keys(result).includes(scope)) {
    debounce(() => setScopedResult(undefined), 100);
  }

  return { result: result[scope] as ResultType|undefined, setResult: (r: ResultType|undefined) => setScopedResult(r) };
}

export const ResultContextProvider: FC = ({ children }) => {

  const [result, setResult] = useState<any>({});

  return (
    <ResultContext.Provider value={{ result, setResult }}>
      {children}
    </ResultContext.Provider>
  );
}