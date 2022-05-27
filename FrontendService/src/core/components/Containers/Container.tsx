import { FC } from "react";
import { createUseStyles } from "react-jss";

type StyleProps = {
  flexDirection: "row" | "column",
  alignItems: "stretch" | "center",
}

const useStyles = createUseStyles({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
    flexDirection: ({ flexDirection }: StyleProps) => flexDirection,
    alignItems: ({ alignItems }: StyleProps) => alignItems,
    '& > *': {
      flexGrow: 1,
    }
  }
});

type Props = {
  className?: string;
  flexDirection?: "row" | "column";
  alignItems?: "stretch" | "center";
}

export const Container: FC<Props> = ({ className, children, flexDirection = "row", alignItems = "stretch" }) => {
  const classes = useStyles({ flexDirection, alignItems });

  return (
    <div className={`${className || ""} ${classes.container}`}>
      {children}
    </div>
  )
}