import { useResult } from "./imports";

export const useUsersResult = () => useResult('users');
export const useReportsResult = () => useResult('reports');
export const useUserBlocksResult = () => useResult('user-blocks');
export const useFriendRequestsResult = () => useResult('friend-requests');
export const useUserFriendsResult = () => useResult('user-friends');