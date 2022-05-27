import { FC, useContext, useState } from "react";
import { createUseStyles } from "react-jss";
import { animated, useSpring } from "react-spring";
import {
  CoreContext,
  DropdownItem,
  AuthService,
  AuthContext,
  Theme
} from './imports';
import { MenuItem } from "./MenuItem";
import DensityMediumIcon from "@mui/icons-material/DensityMedium";
import CloseIcon from "@mui/icons-material/Close";
import MapsHomeWorkIcon from "@mui/icons-material/MapsHomeWork";
import { IconButton } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { Authorized } from "../auth/Authorized";

const useStyles = createUseStyles((theme: Theme) => ({
  innerSidebarContainer: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "space-between",
    height: '100%'
  },
  sideOverlay: {
    width: "100vw",
    height: "100vh",
    position: "fixed",
  },
  sidebarContainer: {
    position: "fixed",
    top: 0,
    width: "250px",
    padding: "25px",
    boxSizing: "border-box",
    height: "100%",
    backgroundColor: theme.colors.primary,
    borderRadius: "0 15px 15px 0",
  },
  sidebar: {
    color: theme.colors.textLight,
  },
  sidebarTop: {
    height: "3em",
    marginBottom: "1em",
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
  },
  sidebarTopLogo: {
    display: "flex",
    alignItems: "center",
    "& h3": {
      marginLeft: "0.6em",
    },
  },
  sidebarOpener: {
    position: "absolute",
    top: 0,
    left: 0,
    cursor: "pointer",
    color: theme.colors.textDark,
  },
  sidebarCloser: {
    color: theme.colors.textLight,
  },
  dropdownContainer: {
    display: 'flex',
    justifyContent: 'center'
  }
}));

type Props = {
};

export const Sidebar: FC<Props> = () => {
  const [open, setOpen] = useState(false);
  const [position, setPosition] = useState(-250);
  const [prevPosition, setPrevPosition] = useState(-250);

  const classes = useStyles();

  const [{ left, inversedLeft, buttonOpacity, sidebarOpacity }] = useSpring(
    {
      to: {
        left: position,
        inversedLeft: -prevPosition,
        buttonOpacity: open ? 0 : 1,
        sidebarOpacity: open ? 1 : 0.5,
      },
      from: {
        left: prevPosition,
        inversedLeft: -position,
        buttonOpacity: open ? 1 : 0,
        sidebarOpacity: open ? 0.5 : 1,
      },
    },
    [open, position, prevPosition]
  );

  const { menuItems } = useContext(CoreContext);

  const openSidebar = () => {
    setOpen(true);
    setPrevPosition(position);
    setPosition(0);
  };

  const closeSidebar = () => {
    setOpen(false);
    setPrevPosition(position);
    setPosition(-250);
  };

  const { setUser, setAuthenticated } = useContext(AuthContext);
  const [authService] = useState(new AuthService());
  const nav = useNavigate();

  const onLogout = () => {
    authService.logout()
      .then(() => {
        setUser(undefined);
        setAuthenticated(false);
        nav('/');
        sessionStorage.removeItem('jwt-token');
      })
      .catch(console.log);
  }

  return (
    <>
      <animated.div
        style={{ left, opacity: sidebarOpacity }}
        className={`${classes.sidebarContainer} ${classes.sidebar}`}
      >
        <div className={classes.innerSidebarContainer}>
          <div>
            <div className={classes.sidebarTop}>
              <span onClick={() => nav('/')} className={classes.sidebarTopLogo}>
                <MapsHomeWorkIcon />
                <h3>druz.io</h3>
              </span>
              <IconButton onClick={closeSidebar}>
                <CloseIcon className={classes.sidebarCloser} />
              </IconButton>
            </div>
            {menuItems.map((item, i) => (
              <MenuItem
                onClick={closeSidebar}
                key={i + item.path}
                menuItem={item}
              />
            ))}
          </div>
          <Authorized>
            <div className={classes.dropdownContainer}>
              <DropdownItem title="Logout" onClick={() => onLogout()} />
            </div>
          </Authorized>
        </div>
      </animated.div>
      <animated.span
        style={{ opacity: buttonOpacity }}
        className={`${classes.sidebarOpener}`}
      >
        <IconButton onClick={openSidebar}>
          <DensityMediumIcon />
        </IconButton>
      </animated.span>
      {open && (
        <animated.div
          style={{ left: inversedLeft }}
          className={classes.sideOverlay}
          onClick={closeSidebar}
        ></animated.div>
      )}
    </>
  );
};
