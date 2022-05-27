import { FC } from "react";

type Props = {
  title: string;
  onClick: () => any;
}

export const DropdownItem: FC<Props> = ({ title, onClick }) => {
  return (
    <span onClick={() => onClick()}>{title}</span>
  );
}
