import { FeaturePlugin, PaddingContainer } from './imports';
import GroupIcon from '@mui/icons-material/Group';
import { authorizedFor } from '../auth-context';
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
        shouldShow: authorizedFor({ roles: ["0"] })
      }
    ],
    pages: [
      {
        component: <PaddingContainer>
            <Posts/>
          </PaddingContainer>,
        path: 'posts',
        shouldShow: authorizedFor({ roles: ["0"] })
      }
    ]
  }
}