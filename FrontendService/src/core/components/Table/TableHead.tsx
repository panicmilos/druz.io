import { FC } from "react";

type Props = {
  columns: string[]
}

export const TableHead: FC<Props> = ({ columns }) => {
  return (
    <thead>
      <tr>
        {columns.map((col: string, i: number) => <th key={i} scope="col">{col}</th>)}
      </tr>
    </thead>
  );
};