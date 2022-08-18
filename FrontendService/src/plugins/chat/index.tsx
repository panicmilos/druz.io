import { FeaturePlugin, PaddingContainer } from './imports';
import GroupIcon from '@mui/icons-material/Group';
import { authorizedFor } from '../auth-context';
import { Chat } from './pages/Chat/Chat';

export * from './exports';

export function getPluginDefinition(): FeaturePlugin {
  return {
    id: 'Chats',
    type: 'FeaturePlugin',
    menuItems: [
      {
        label: 'Chat',
        path: 'chat',
        icon: <GroupIcon/>,
        shouldShow: authorizedFor({ roles: ["0"] })
      }
    ],
    pages: [
      {
        component: <PaddingContainer>
            <Chat/>
          </PaddingContainer>,
        path: 'chat',
        shouldShow: authorizedFor({ roles: ["0"] })
      }
    ]
  }
}