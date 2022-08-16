import moment from "moment";
import { FC } from "react";
import { createUseStyles } from "react-jss";
import { Col, Container, Table, TableBody, TableHead, TableRow } from "../../imports";
import { Profile } from "../../models/User";

type Props = {
  user?: Profile
}


const useStyles = createUseStyles({
  parentContainer: {
    display: 'flex',
    width: '100%'
  },
  soloContainer: {
    width: '100%',
    flexGrow: 1
  },
  twoContainer: {
    minWidth: '30%',
    maxWidth: '100%'
  },
  firstContainer: {
    marginRight: '2%',
  },
  growFlex: {
    flexGrow: 1
  }
});


const LivePlaceTypes = ['Currently', 'Lived', 'Birthplace'];


export const AdditionalProfileForm: FC<Props> = ({ user }) => {

  const classes = useStyles();

  return (
    <>

    <div className={classes.parentContainer}>
      {
        user?.LivePlaces?.length ? 
          <div className={(user?.LivePlaces?.length && user?.WorkPlaces?.length) ? `${classes.twoContainer} ${classes.firstContainer}` : classes.soloContainer}>
            <Table hasPagination={false}>
              <TableHead columns={['Place', 'Type']}/>
              <TableBody>
              {
                user?.LivePlaces?.map((livePlace, i) =>
                  <TableRow
                    key={i}          
                    cells={[
                      livePlace.Place,
                      LivePlaceTypes[livePlace.LivePlaceType],
                    ]}
                  />
                )
              }
              </TableBody>
            </Table>
          </div> : <></>
        }
      
        {
          user?.WorkPlaces?.length ?
            <div className={(user?.LivePlaces?.length && user?.WorkPlaces?.length) ? `${classes.twoContainer} ${classes.growFlex}` : classes.soloContainer}>
              <Table hasPagination={false}>
                <TableHead columns={['Place', 'From', 'To']}/>
                <TableBody>
                {
                  user?.WorkPlaces?.map((workPlace, i) =>
                    <TableRow
                      key={i}          
                      cells={[
                        workPlace.Place,
                        moment(workPlace.To).format('yyyy-MM-DD'),
                        moment(workPlace.From).format('yyyy-MM-DD'),
                      ]}
                    />
                  )
                }
                </TableBody>
              </Table>
            </div> : <></>
        }

    </div>
    
    <br />
    <br />
    <br />

    <div className={classes.parentContainer}>
      
      {
        user?.Intereses?.length ?
        <div className={(user?.Educations?.length && user?.Intereses?.length) ? `${classes.twoContainer} ${classes.firstContainer}` : classes.soloContainer}>
          <Table hasPagination={false}>
            <TableHead columns={['Interes']}/>
            <TableBody>
            {
              user?.Intereses?.map((interes, i) =>
                <TableRow
                  key={i}          
                  cells={[
                    interes.Interes,
                  ]}
                />
              )
            }
            </TableBody>
          </Table>
        </div> : <></>
      }

{
        user?.Educations?.length ?
        <div className={(user?.Educations?.length && user?.Intereses?.length) ? `${classes.twoContainer} ${classes.growFlex}` : classes.soloContainer}>
          <Table hasPagination={false}>
            <TableHead columns={['Place', 'From', 'To']}/>
            <TableBody>
            {
              user?.Educations?.map((education, i) =>
                <TableRow
                  key={i}          
                  cells={[
                    education.Place,
                    moment(education.To).format('yyyy-MM-DD'),
                    moment(education.From).format('yyyy-MM-DD'),
                  ]}
                />
              )
            }
            </TableBody>
          </Table>
        </div> : <></>
      }

    </div>
    </>
  )
}