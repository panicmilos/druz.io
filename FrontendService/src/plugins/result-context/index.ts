import { ResultContextProvider } from "./Context";
import { ContextPlugin } from "./imports";

export * from './exports';

export function getPluginDefinition(): ContextPlugin {
  return {
    id: 'ResultContext',
    type: 'ContextPlugin',
    Provider: ResultContextProvider
  }
}