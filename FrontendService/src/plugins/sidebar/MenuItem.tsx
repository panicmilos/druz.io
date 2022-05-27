import { FC} from "react";
import { createUseStyles } from "react-jss";
import { Link } from "react-router-dom";
import { MenuItem as MenuItemType, Theme } from "./imports";

type Props = {
  menuItem: MenuItemType;
  onClick?: () => any;
}

const useStyles = createUseStyles((theme: Theme) => ({
  menuItemContainer: {
    borderRadius: '15px',
    padding: '.6em',
    textDecoration: 'none',
    color: theme.colors.textLight,
    display: 'flex',
    marginBottom: '0.5em',
    alignItems: 'center',
    '&:hover': {
      backgroundColor: 'rgba(0, 0, 0, 0.04)',
    },
    transition: '0.5s'
  },
  menuItemLabel: {
    marginLeft: '.7em'
  }
}));

export const MenuItem: FC<Props> = ({ menuItem, onClick = () => {} }) => {
  const classes = useStyles();

  const shouldShow = !menuItem.shouldShow || menuItem.shouldShow();

  const Icon = menuItem.icon ? menuItem.icon : <></>;
  return shouldShow ? (
    <Link
      to={menuItem.path}
      className={`${classes.menuItemContainer}`}
      onClick={onClick}
    >
      {Icon}
      <p className={classes.menuItemLabel}>{menuItem.label}</p>
    </Link>
  ) : <></>;
}