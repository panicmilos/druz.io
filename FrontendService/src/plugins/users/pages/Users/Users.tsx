import { useState } from "react";
import { useQuery } from "react-query";
import { Card } from "../../imports";
import { useUserService } from "../../services"
import { SearchUserForm } from "./SearchUserForm";
import { UsersList } from "./UsersList";


export const Users = () => {


  const userService = useUserService();
  const [params, setParams] = useState<any>({});

  const { data: users } = useQuery([params, userService], () => userService.search(params));

  return (
    <>
      <Card title="Users">
        <SearchUserForm onSearch={setParams} />

        <UsersList users={users || []} />
      </Card>

    </>
  ); 
  
}