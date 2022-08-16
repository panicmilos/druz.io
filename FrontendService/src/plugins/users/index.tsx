import { FeaturePlugin, PaddingContainer } from './imports';
import GroupIcon from '@mui/icons-material/Group';
import { authorizedFor, unauthorized } from '../auth-context';
import { Registration } from './pages/Register/Registration';
import { ChangeProfile } from './pages/Profile/ChangeProfile';

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
        label: 'Profile Settings',
        path: 'profile/settings',
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
            <ChangeProfile/>
          </PaddingContainer>,
        path: 'profile/settings',
        shouldShow: authorizedFor({ roles: ["0"] })
      }
    ]
  }
}