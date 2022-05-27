import { ContextPlugin } from "./imports";
import { QueryContextProvider } from "./Context";

export function getPluginDefinition(): ContextPlugin {
  return {
    id: 'ReactQueryContext',
    type: 'ContextPlugin',
    Provider: QueryContextProvider
  }
}