import { FeaturePlugin, PaddingContainer } from './imports';
import GroupIcon from '@mui/icons-material/Group';
import { authorizedFor, USER_ROLE } from '../auth-context';
import { Posts } from './pages/Posts/Posts';

export * from './exports';

export function getPluginDefinition(): FeaturePlugin {
  return {
    id: 'Posts',
    type: 'FeaturePlugin',
    menuItems: [
      {
        label: 'Posts',
        path: 'posts',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      }
    ],
    pages: [
      {
        component: <PaddingContainer>
            <Posts/>
          </PaddingContainer>,
        path: 'posts',
        shouldShow: authorizedFor({ roles: [USER_ROLE] })
      }
    ]
  }
}