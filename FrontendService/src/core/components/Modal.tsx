import { IconButton } from "@mui/material";
import { FC, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { animated, useSpring } from "react-spring";
import { Card } from "./Card";
import CloseIcon from '@mui/icons-material/Close';

const useStyles = createUseStyles({
  overlayContainer: {
    backgroundColor: 'rgba(149, 157, 165, 0.2)',
    width: '100vw',
    height: '100vh',
    position: 'absolute',
    boxSizing: 'border-box',
    top: 0,
    left: 0,
    padding: "5vw",
    cursor: 'pointer'
  },
  contentContainer: {
    marginLeft: 'auto',
    marginRight: 'auto',
    cursor: 'initial',
    width: '60%'
  }
});

type Props = {
  title?: string;
  open: boolean;
  onClose?: () => any;
}

export const Modal: FC<Props> = ({ open, onClose, title, children }) => {
  const classes = useStyles();

  const [isOpen, setOpen] = useState(open);

  const [animationState, setAnimationState] = useState({
    opacity: 0,
    prevOpacity: 0,
    offset: 0,
    prevOffset: 0,
  });

  const [props] = useSpring({
    from: {
      offset: animationState.prevOffset,
      opacity: animationState.prevOpacity
    },
    to: {
      offset: animationState.offset,
      opacity: animationState.opacity
    },
    onRest: () => {
      setOpen(open);
    },
    config: {
      duration: 500,
    }
  }, [animationState]);

  useEffect(() => {
    if(open) setOpen(true);
    setAnimationState(open ? {
      opacity: 1,
      prevOpacity: 0,
      offset: 0,
      prevOffset: 30,
    } :{
      opacity: 0,
      prevOpacity: 1,
      offset: 30,
      prevOffset: 0,
    });
  }, [open]);

  const handleClickOverlay = () => {
    onClose && onClose();
  }

  return !isOpen ? <></> : (
    <animated.div style={{ opacity: props.opacity }} className={classes.overlayContainer} onClick={handleClickOverlay}>
      <animated.div style={{ marginTop: props.offset }} className={classes.contentContainer} onClick={e => e.stopPropagation()}>
        <Card
          title={title}
          animationDisabled={true}
          IconRight={
          <IconButton onClick={handleClickOverlay}>
            <CloseIcon />
          </IconButton>}
        >
          {children}
        </Card>
      </animated.div>
    </animated.div>
  )
}