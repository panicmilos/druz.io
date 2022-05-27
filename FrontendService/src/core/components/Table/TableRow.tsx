import { FC } from "react";

type CellContent = string | JSX.Element;

type Props = {
  cells: CellContent[]
}

export const TableRow: FC<Props> = ({ cells }) => {
  return (
    <tr>
      {cells.map((cell: CellContent, i: number) => <td key={i}>{cell}</td>)}
    </tr>
  );
};