import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../../Context";

type UseStylesProps = {
  xs: number;
  sm: number;
  md: number;
  lg: number;
}

type Props = {
  all?: number;
  xs?: number;
  sm?: number;
  md?: number;
  lg?: number;
  className?: string;
}

export const Col: FC<Props> = ({ all=12, xs, sm, md, lg, className, children }) => {

  xs = xs ?? all;
  sm = sm ?? xs;
  md = md ?? sm;
  lg = lg ?? md;

  const useStyles = createUseStyles((theme: Theme) => ({
    container: {
      flexBasis: ({ xs }: UseStylesProps) => `${100*xs/12}%`,
    },
    [theme.breakpoints.sm]: {
      container: {
        flexBasis: ({ sm }: UseStylesProps) => `${100*sm/12}%`,
      }
    },
    [theme.breakpoints.md]: {
      container: {
        flexBasis: ({ md }: UseStylesProps) => `${100*md/12}%`,
      }
    },
    [theme.breakpoints.lg]: {
      container: {
        flexBasis: ({ lg }: UseStylesProps) => `${100*lg/12}%`,
      }
    },
  }));

  const classes = useStyles({ xs, sm, md, lg });

  return (
    <div className={`${className || ""} ${classes.container}`}>
      {children}
    </div>
  )
}