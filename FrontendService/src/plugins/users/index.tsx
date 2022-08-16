import { FeaturePlugin, PaddingContainer } from './imports';
import GroupIcon from '@mui/icons-material/Group';
import { authorizedFor, unauthorized } from '../auth-context';
import { Registration } from './pages/Register/Registration';
import { ChangeProfile } from './pages/ChangeProfile/ChangeProfile';
import { RequestReactivation } from './pages/ReactivationRequests/RequestReactivation';
import { Reactivation } from './pages/Reactivation/Reactivation';
import { ForgotPasswordRequest } from './pages/ForgotPasswordRequest/ForgotpasswordRequest';
import { ForgotPassword } from './pages/ForgotPassword/ForgotPassword';
import { Profile } from './pages/Profile/Profile';
import { Reports } from './pages/Reports/Reports';
import { Users } from './pages/Users/Users';

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
        shouldShow: authorizedFor({ roles: ["0"] })
      },
      {
        label: 'User Reports',
        path: 'users/reports',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: ["0", "1"] })
      },
      {
        label: 'Users',
        path: 'users',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: ["0"] })
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
        path: 'users/:id/'
      },
      {
        component: <PaddingContainer>
            <ChangeProfile/>
          </PaddingContainer>,
        path: 'profile/settings',
        shouldShow: authorizedFor({ roles: ["0"] })
      },
      {
        component: <PaddingContainer>
            <Reports/>
          </PaddingContainer>,
        path: 'users/reports',
        shouldShow: authorizedFor({ roles: ["0", "1"] })
      },
      {
        component: <PaddingContainer>
            <Users/>
          </PaddingContainer>,
        path: 'users',
        shouldShow: authorizedFor({ roles: ["0"] })
      }
    ]
  }
}