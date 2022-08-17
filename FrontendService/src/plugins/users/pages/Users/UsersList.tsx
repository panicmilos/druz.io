import { FC } from "react"
import { createUseStyles } from "react-jss";
import { useNavigate } from "react-router-dom";
import { Container } from "../../imports";
import { Profile } from "../../models/User"

type Props = {
  users: Profile[]
}

const chunks = (a: any[], size: number): any[] =>
  Array.from(
    new Array(Math.ceil(a.length / size)),
    (_, i) => a.slice(i * size, i * size + size)
  );

const useStyles = createUseStyles({
  container: {
    maxWidth: '17%',
    margin: '3% 3% 3% 3%',
    border: '1px solid #000000',
    '& p': {
      textAlign: 'center'
    }
  }
});

export const UsersList:FC<Props> = ({ users }) => {

  const nav = useNavigate();

  const userChunks = chunks(users || [], 5);

  const classes = useStyles();

  return (
    <>
      {
        userChunks?.map((chunk: Profile[]) => {

          return (
            <Container>
              {
                chunk?.map((profile: Profile) => {
                  return (
                    <div onClick={() => nav(`/users/${profile.ID}`)} className={classes.container}>
                      <img src={profile.Image} width={'100%'} />
                      <p>{`${profile.FirstName} ${profile.LastName}`}</p>
                    </div>
                  )
                })
              }
            </Container>
          )

        })
      }
    </>
  );
}