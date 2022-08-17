import { useQuery } from "react-query";
import { useResult } from "../result-context";
import { useUserService } from "../users";
import { createEntitiesMap } from "./imports";

export const usePostsResult = () => useResult('posts');

export const useUsersMap = () => {

  const userService = useUserService();
  
  const { data: users } = useQuery([userService], () => userService.search({}));

  return createEntitiesMap(users, user => user.ID, user => `${user.FirstName} ${user.LastName}`);

}