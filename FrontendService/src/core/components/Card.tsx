import { FC } from "react";
import { createUseStyles } from "react-jss";
import { animated, useSpring } from "react-spring";

const useStyles = createUseStyles({
  cardContainer: {
    border: '0.1px solid rgba(0, 0, 0, .05)',
    borderRadius: '15px',
    backgroundColor: "#fff",
    padding: '1em',
    margin: '1em',
    boxShadow: 'rgba(149, 157, 165, 0.2) 0px 8px 24px'
  },
  cardTitle: {
    marginBottom: '1em',
    fontWeight: '500',
    fontSize: '1.4rem',
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center'
  }
});

type Props = {
  title?: string;
  animationDisabled?: boolean;
  animation?: {
    delay: number;
    duration: number;
  },
  IconRight?: JSX.Element;
  className?: string;
}

export const Card: FC<Props> = ({ title, children, animation, animationDisabled = false, IconRight=<></>, className }) => {
  const classes = useStyles();

  const {opacity} = useSpring({
    from: {
      opacity: 0
    },
    to: {
      opacity: 1
    },
    delay: animation?.delay ?? 200,
    config: {
      duration: animation?.duration ?? 600,
    }
  });

  const Component = !animationDisabled ? animated.div : ({ children }: any) => <div className={classes.cardContainer}>{children}</div>;

  return (
    <Component style={{opacity}} className={`${classes.cardContainer} ${className || ""}`}>
      {title && <div className={classes.cardTitle}>
        <span>{title}</span>
        {IconRight}
      </div>}
      {children}
    </Component>
  )
}