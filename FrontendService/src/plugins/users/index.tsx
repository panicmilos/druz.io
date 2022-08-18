import { FeaturePlugin, PaddingContainer, USER_ROLE } from './imports';
import GroupIcon from '@mui/icons-material/Group';
import { ADMIN_ROLE, authorizedFor, unauthorized } from '../auth-context';
import { Registration } from './pages/Register/Registration';
import { ChangeProfile } from './pages/ChangeProfile/ChangeProfile';
import { RequestReactivation } from './pages/ReactivationRequests/RequestReactivation';
import { Reactivation } from './pages/Reactivation/Reactivation';
import { ForgotPasswordRequest } from './pages/ForgotPasswordRequest/ForgotpasswordRequest';
import { ForgotPassword } from './pages/ForgotPassword/ForgotPassword';
import { Profile } from './pages/Profile/Profile';
import { Reports } from './pages/Reports/Reports';
import { Users } from './pages/Users/Users';
import { BlockedUsers } from './pages/BlockedUsers/BlockedUsers';
import { FriendRequests } from './pages/FriendRequests/FriendRequests';
import { UserFriends } from './pages/UserFriends/UserFriends';

export * from './exports';

export function getPluginDefinition(): FeaturePlugin {
  return {
    id: 'Users',
    type: 'FeaturePlugin',
    menuItems: [
      {
        label: 'Register',
        path: 'register',
        icon: <GroupIcon/>,
        shouldShow: unauthorized()
      },
      {
        label: 'Request Reactivation',
        path: 'request-reactivation',
        icon: <GroupIcon/>,
        shouldShow: unauthorized()
      },
      {
        label: 'Request Password Recovery',
        path: 'password-recovery',
        icon: <GroupIcon/>,
        shouldShow: unauthorized()
      },
      {
        label: 'Profile Settings',
        path: 'profile/settings',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      },
      {
        label: 'User Reports',
        path: 'users/reports',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [ADMIN_ROLE] })
      },
      {
        label: 'Users',
        path: 'users',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [USER_ROLE, ADMIN_ROLE] })
      },
      {
        label: 'Blocked Users',
        path: 'users/blocked',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      },
      {
        label: 'Friends List',
        path: 'users/friends',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      },
      {
        label: 'Friend Requests',
        path: 'users/friends/requests',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      }
    ],
    pages: [
      {
        component: <PaddingContainer>
            <Registration/>
          </PaddingContainer>,
        path: 'register',
        shouldShow: unauthorized()
      },
      {
        component: <PaddingContainer>
            <RequestReactivation />
          </PaddingContainer>,
        path: 'request-reactivation',
        shouldShow: unauthorized()
      },
      {
        component: <PaddingContainer>
            <Reactivation />
          </PaddingContainer>,
        path: 'users/:id/reactivation',
        shouldShow: unauthorized()
      },
      {
        component: <PaddingContainer>
            <ForgotPasswordRequest />
          </PaddingContainer>,
        path: 'password-recovery',
        shouldShow: unauthorized()
      },
      {
        component: <PaddingContainer>
            <ForgotPassword />
          </PaddingContainer>,
        path: 'users/:id/password/recovery',
        shouldShow: unauthorized()
      },
      {
        component: <PaddingContainer>
            <Profile />
          </PaddingContainer>,
        path: 'users/:id/',
        shouldShow: authorizedFor({ roles: [USER_ROLE, ADMIN_ROLE]})
      },
      {
        component: <PaddingContainer>
            <ChangeProfile/>
          </PaddingContainer>,
        path: 'profile/settings',
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      },
      {
        component: <PaddingContainer>
            <Reports/>
          </PaddingContainer>,
        path: 'users/reports',
        shouldShow: authorizedFor({ roles: [ADMIN_ROLE] })
      },
      {
        component: <PaddingContainer>
            <Users/>
          </PaddingContainer>,
        path: 'users',
        shouldShow: authorizedFor({ roles: [USER_ROLE, ADMIN_ROLE] })
      },
      {
        component: <PaddingContainer>
            <BlockedUsers/>
          </PaddingContainer>,
        path: 'users/blocked',
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      },
      {
        component: <PaddingContainer>
            <UserFriends/>
          </PaddingContainer>,
        path: 'users/friends',
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      },
      {
        component: <PaddingContainer>
            <FriendRequests/>
          </PaddingContainer>,
        path: 'users/friends/requests',
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      }
    ]
  }
}