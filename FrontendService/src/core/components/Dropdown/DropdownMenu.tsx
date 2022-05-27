import { FC, useEffect, useRef, useState } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../../Context";
import { Button } from "../Button";

const useStyles = createUseStyles((theme: Theme) => ({
  dropdown: {
    position: 'relative',
    display: "inline-block",
  },
  dropdownContent: {
    display: ({ show }: { show: boolean }) => show ? "block" : "none",
    position: "absolute",
    top: 'auto',
    bottom: '100%',
    backgroundColor: "#f1f1f1",
    minWidth: "160px",
    overflow: "auto",
    boxShadow: "0px 8px 16px 0px rgba(0,0,0,0.2)",
    zIndex: 1,
    '& span': {
      cursor: 'pointer',
      color: 'black',
      padding: '12px 16px',
      textDecoration: 'none',
      display: 'block',
    },
    '& span:hover': {
      backgroundColor: '#ddd',
    },
  },
}));

type Props = {
  className?: string;
  title?: string | JSX.Element;
};

export const DropdownMenu: FC<Props> = ({ children, className = "", title }) => {
  const [show, setShow] = useState(false);
  const classes = useStyles({ show });

  const wrapperRef = useRef(null);

  useEffect(() => {
    window.onclick = function(event) {
      if (wrapperRef?.current && !(wrapperRef?.current as any)?.contains(event.target)) {
        setShow(false);
      }
    }
  }, [wrapperRef]);

  return (
    <div ref={wrapperRef} className={`${classes.dropdown} ${className}`}>
      <Button
        onClick={() => {
          setShow(!show);
        }}
      >
        {title}
      </Button>
      <div className={classes.dropdownContent}>
        {children}
      </div>
    </div>
  );
};
