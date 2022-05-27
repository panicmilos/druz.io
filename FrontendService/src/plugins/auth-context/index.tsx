import { ContextPlugin } from "./imports";
import { AuthContextProvider } from "./Context";

export * from './exports';

export function getPluginDefinition(): ContextPlugin {
  return {
    id: 'AuthContext',
    type: 'ContextPlugin',
    Provider: AuthContextProvider
  }
}