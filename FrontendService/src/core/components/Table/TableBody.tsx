import { FC } from "react";

type Props = {

}

export const TableBody: FC<Props> = ({ children }) => {
  return (
    <tbody>
      {children}
    </tbody>
  );
};
