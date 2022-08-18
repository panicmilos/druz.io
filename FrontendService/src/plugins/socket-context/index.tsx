import { SocketContextProvider } from "./Context";
import { ContextPlugin } from "./imports";

export * from './Context';

export function getPluginDefinition(): ContextPlugin {
  return {
    id: 'SocketsContext',
    type: 'ContextPlugin',
    Provider: SocketContextProvider
  }
}