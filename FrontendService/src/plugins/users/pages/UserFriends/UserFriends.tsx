import { useContext } from "react";
import { useQuery } from "react-query";
import { AuthContext, Card } from "../../imports";
import { useUserFriendsService } from "../../services";
import { UsersList } from "../Users/UsersList";


export const UserFriends = () => {

  const { user } = useContext(AuthContext);

  const userFriendsService = useUserFriendsService(user?.ID ?? '');

  const { data: userFriends } = useQuery([userFriendsService], () => userFriendsService.fetch());

  const users = userFriends?.map(uf => uf.Friend) ?? [];

  return (
    <>
      <Card title="User Friends">

        <UsersList users={users || []} />
      </Card>

    </>
  ); 
  
}