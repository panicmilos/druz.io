import { useState } from 'react';
import { FC } from "react";
import { createUseStyles } from 'react-jss';
import { Theme } from '../../Context';
import KeyboardArrowLeftIcon from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRightIcon from '@mui/icons-material/KeyboardArrowRight';
import { Button } from '../Button';

type Props = {
  onPageChange: (n: number) => any;
  totalPages?: number;
}

const useStyles = createUseStyles((theme: Theme) => ({
  paginationContainer: {
    display: 'flex',
    justifyContent: 'flex-end'
  },
  buttonsContainer: {
    marginTop: '12px',
    display: 'flex',
  },
  button: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    textAlign: 'center',
    width: '20px',
    height: '20px',
    marginLeft: '10px',
    padding: '15px',
    borderRadius: '50%',
    userSelect: 'none',
    '&:active': {
      padding: '15px'
    }
  },
  active: {
    backgroundColor: `${theme.colors.secondary} !important`
  }
}));

export const Pagination: FC<Props> = ({ onPageChange, totalPages = Infinity }) => {
  const classes = useStyles();

  const [currentPage, setCurrentPage] = useState(1);
  const [anchoredPage, setAnchoredPage] = useState(2);

  const handlePageChangeClick = (newPage: number) => (e: any) => {
    e.preventDefault();
    if (newPage > 0) {
      setCurrentPage(newPage);
      if (newPage === anchoredPage + 2) {
        setAnchoredPage(anchoredPage + 1);
      } else if (newPage === anchoredPage - 2) {
        setAnchoredPage(anchoredPage - 1);
      }
      onPageChange(newPage);
    }
  }

  return (
    <div className={classes.paginationContainer}>
      <ul className={classes.buttonsContainer}>
        {
          currentPage !== 1 &&
          <li onClick={handlePageChangeClick(currentPage - 1)}>
            <Button className={classes.button}><KeyboardArrowLeftIcon /></Button>
          </li>
        }
        {[anchoredPage - 1, anchoredPage, anchoredPage + 1].map((page, i) =>
          page <= totalPages &&
          <li key={i} onClick={handlePageChangeClick(page)}>
            <Button className={`${classes.button} ${currentPage === page ? classes.active : ''}`}>{page}</Button>
          </li>
        )}
        {
          currentPage < totalPages &&
          <li onClick={handlePageChangeClick(currentPage + 1)}>
            <Button className={classes.button}><KeyboardArrowRightIcon /></Button>
          </li>
        }
      </ul>
    </div>
  );
};