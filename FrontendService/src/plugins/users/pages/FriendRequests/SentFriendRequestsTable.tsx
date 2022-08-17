import { FC } from "react";
import { Table, TableBody, TableHead, TableRow } from "../../imports";
import { FriendRequest } from "../../models/FriendRequest";

type Props = {
  friendRequests: FriendRequest[]
}



export const SentFriendRequestsTable: FC<Props> = ({ friendRequests }) => {

  
  return (
    <Table hasPagination={false}>
      <TableHead columns={['Sent to']}/>
      <TableBody>
        {
          friendRequests?.map((friendRequest: FriendRequest) => 
          <TableRow 
            key={friendRequest.ID}
            cells={[
              `${friendRequest.Friend.FirstName} ${friendRequest.Friend.LastName}`,
          ]}/>
          )
        }
      </TableBody>
    </Table>
    );
  }