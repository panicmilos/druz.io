import { useState } from "react";
import { useQuery } from "react-query";
import { useResult } from "../result-context";
import { useUserService } from "../users";
import { createEntitiesMap } from "./imports";

export const usePostsResult = () => useResult('posts');

export const useUsersMap = () => {

  const userService = useUserService();
  const [usersMap, setUsersMap] = useState<any>({});
  
  useQuery([userService, 'userNames'], () => userService.search({}), {
    onSuccess: (users) => setUsersMap(createEntitiesMap(users, user => user.ID, user => `${user.FirstName} ${user.LastName}`))
  });

  return usersMap;
}

export const useUserImagesMap = () => {

  const userService = useUserService();
  const [usersMap, setUsersMap] = useState<any>({});
  
  useQuery([userService, 'images'], () => userService.search({}), {
    onSuccess: (users) => setUsersMap(createEntitiesMap(users, user => user.ID, user => user?.Image))
  });

  return usersMap;
}