import { FC, useContext, useEffect } from "react";
import { useQuery } from "react-query";
import { useFriendRequestsResult } from "../../hooks";
import { AuthContext, Card } from "../../imports";
import { useFriendRequestsService } from "../../services";
import { ReceivedFriendRequestsTable } from "./ReceivedFriendRequestsTable";
import { SentFriendRequestsTable } from "./SentFriendRequestsTable";


export const FriendRequests: FC = () => {

  const { user } = useContext(AuthContext);

  const friendRequestsService = useFriendRequestsService(user?.ID ?? '');

  const { result, setResult } = useFriendRequestsResult();

  const sent = useQuery(['sent', result, friendRequestsService], () => friendRequestsService.fetchSent(), { enabled: !result });
  const received = useQuery(['received', result, friendRequestsService], () => friendRequestsService.fetchReceived(), { enabled: !result });

  const sentFriendRequests = sent.data;
  const receivedFriendRequests = received.data;

  useEffect(() => {
    if (!result) return;
        
    setResult(undefined);
  }, [result]);
  
  return (
    <>

      <Card title="Sent Requests">
        <SentFriendRequestsTable friendRequests={sentFriendRequests || []} />
      </Card>

      <Card title="Received Requests">
        <ReceivedFriendRequestsTable friendRequests={receivedFriendRequests || []} />
      </Card>
    </>
  );
}