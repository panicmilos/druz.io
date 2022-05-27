import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "../../Context";

type Props = {
  hasPagination?: boolean;
  children: any[]
}

const borderRadius = '5px';

const useStyles = createUseStyles((theme: Theme) => ({
  table: {
    borderCollapse: 'collapse',
    width: '100%',
    '& tr:nth-child(even)': {
      backgroundColor: '#f2f2f2',
      borderRadius: `${borderRadius} 0 0 ${borderRadius}`
    },
    '& tr:hover': {
      backgroundColor: '#ddd',
    },
    '& th': {
      paddingTop: '12px',
      paddingLeft: '20px',
      paddingBottom: '12px',
      textAlign: 'left',
      backgroundColor: theme.colors.primary,
      color: theme.colors.textLight,
    },
    '& th:first-child': {
      borderRadius: `${borderRadius} 0 0 ${borderRadius}`
    },
    '& th:last-child': {
      borderRadius: `0 ${borderRadius} ${borderRadius} 0`
    },
    '& td:first-child': {
      borderRadius: `${borderRadius} 0 0 ${borderRadius}`
    },
    '& td:last-child': {
      borderRadius: `0 ${borderRadius} ${borderRadius} 0`
    },
    '& td': {
      paddingLeft: '20px',
      paddingTop: '10px',
      paddingBottom: '10px',
    }
  }
}))

export const Table: FC<Props> = ({ children, hasPagination=true }) => {
  const classes = useStyles();

  return (
    <>
      <table className={classes.table}>
        {hasPagination ? children.slice(0, children.length - 1) : children}
      </table>
      {hasPagination && children[children.length - 1]}
    </>
  );
};