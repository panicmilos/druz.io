import { useContext, useEffect, useState } from "react";
import { useQuery } from "react-query";
import { createEntitiesMap } from "../../core";
import { AuthContext } from "../auth-context";
import { SocketContext } from "../socket-context";
import { useUserFriendsService } from "../users";
import { useChatService } from "./services";

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

export const useStatusesMap = () => {
  const chatService = useChatService();

  const [statusesMap, setStatusesMap] = useState<any>({});

  useQuery([chatService], () => chatService.fetchStatuses(), {
    onSuccess: (statuses) => {
      const statusesMap = createEntitiesMap(statuses, status => status.UserId, status => status.Status);
      setStatusesMap(statusesMap);
    }
  });

  const { client } = useContext(SocketContext);

  useEffect(() => {
    client?.on('statuses', function (data: any) {
      const notification = JSON.parse(data.text);
      const { Status, User: { ID } } = notification;
      const newStatusesMap = {...statusesMap};
      newStatusesMap[ID.replace('users/', '')] = Status;
      setStatusesMap(newStatusesMap);
    })

    return () => { client?.removeAllListeners('statuses'); };
  }, [client])

  return statusesMap;
}