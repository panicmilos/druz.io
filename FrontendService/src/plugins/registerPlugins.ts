import { Plugin } from '../core/Context';

export const registerPlugins = (registerPlugin: (plugin: Plugin) => any) => {
  import('.').then((module: { [key: string]: () => Plugin }) => {
    Object.keys(module).forEach((pluginKey: string) => {
      const getPluginDefinition = module[pluginKey];
      const plugin = getPluginDefinition();
      registerPlugin(plugin);
    });
  });
}