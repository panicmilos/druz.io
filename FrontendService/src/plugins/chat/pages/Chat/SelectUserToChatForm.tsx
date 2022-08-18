import { FC, useContext } from "react";
import { createUseStyles } from "react-jss";
import { useQuery } from "react-query";
import { AuthContext, Button, Form, FormSelectOptionInput, useUserFriendsService } from "../../imports";

const useStyles = createUseStyles({
  submitButton: {
    marginTop: '1em',
  }
});

type Props = {
  onSubmit: (any: string) => any
}

export const SelectUserToChatForm: FC<Props> = ({ onSubmit }) => {

  const { user } = useContext(AuthContext);

  const userFriendService = useUserFriendsService(user?.ID ?? '');

  const { data: userFriend } = useQuery([userFriendService], () => userFriendService.fetch());
  const userFriendOptions = userFriend?.map(uf => ({ value: uf.FriendId, label: `${uf.Friend.FirstName} ${uf.Friend.LastName}` })); 

  const classes = useStyles();

  return (
    <>
      <Form
        schema={undefined}
        onSubmit={(values) => values.UserId && onSubmit(values.UserId)}
      >
        <FormSelectOptionInput label="Friend" name="UserId" options={userFriendOptions || []} />

        <Button className={classes.submitButton} type="submit">Submit</Button>
      </Form>
    </>
  )

}