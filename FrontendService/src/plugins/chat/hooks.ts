import { useContext } from "react";
import { useQuery } from "react-query";
import { createEntitiesMap } from "../../core";
import { AuthContext } from "../auth-context";
import { useUserFriendsService } from "../users";

export const useUserFriendNamesMap = () => {

  const { user } = useContext(AuthContext);

  const userFriendsService = useUserFriendsService(user?.ID ?? '');
  
  const { data: userFriends } = useQuery([userFriendsService], () => userFriendsService.fetch());

  return createEntitiesMap(userFriends, userFriend => userFriend.FriendId, userFriend => `${userFriend.Friend.FirstName} ${userFriend.Friend.LastName}`);
}

export const useUserFriendsMap = () => {

  const { user } = useContext(AuthContext);

  const userFriendsService = useUserFriendsService(user?.ID ?? '');
  
  const { data: userFriends } = useQuery([userFriendsService], () => userFriendsService.fetch());

  return createEntitiesMap(userFriends, userFriend => userFriend.FriendId, userFriend => userFriend.Friend);

}