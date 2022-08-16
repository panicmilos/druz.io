import { FC, useContext, useEffect } from "react";
import { useQuery } from "react-query";
import { useUserBlocksResult } from "../../hooks";
import { AuthContext, Card } from "../../imports";
import { useUserBlocksService } from "../../services";
import { BlockedUsersTable } from "./BlockedUsersTable";

export const BlockedUsers: FC = () => {

  const { user } = useContext(AuthContext);

  const userBlocksService = useUserBlocksService(user?.ID ?? '');

  const { result, setResult } = useUserBlocksResult();

  const { data: userBlocks } = useQuery([userBlocksService], () => userBlocksService.fetch(), { enabled: !result });

  console.log(userBlocks);

  useEffect(() => {
    if (!result) return;
        
    setResult(undefined);
  }, [result]);
  
  return (
    <>
      <Card title="Block List">

        <BlockedUsersTable userBlocks={userBlocks || []} />
      </Card>
    </>
  )
}