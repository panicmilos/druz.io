import { FeaturePlugin } from "./imports";
import { unauthorized } from "../auth-context/filter";
import { Login } from "./pages/Login";
import LoginIcon from '@mui/icons-material/Login';

export * from './exports';

export function getPluginDefinition(): FeaturePlugin {
  return {
    id: 'Auth',
    type: 'FeaturePlugin',
    menuItems: [
      {
        label: 'Login',
        path: '/',
        icon: <LoginIcon/>,
        shouldShow: unauthorized()
      }
    ],
    pages: [
      {
        component: <Login/>,
        path: '/'
      }
    ]
  }
}