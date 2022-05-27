import { Plugin } from '../core/Context';

export const registerPlugins = (registerPlugins: (plugins: Plugin[]) => any) => {
  import('.').then((module: { [key: string]: () => Plugin }) => {
    const contextPlugins: Plugin[] = [];
    const featurePlugins: Plugin[] = [];
    const sidebarPlugins: Plugin[] = [];
    Object.keys(module).forEach((pluginKey: string) => {
      const getPluginDefinition = module[pluginKey];
      const plugin = getPluginDefinition();
      if(plugin.type === 'SidebarPlugin')
        sidebarPlugins.push(plugin);
      if(plugin.type === 'ContextPlugin')
        contextPlugins.push(plugin);
      if(plugin.type === 'FeaturePlugin')
        featurePlugins.push(plugin);
    });

    registerPlugins([
      ...contextPlugins,
      ...featurePlugins,
      ...sidebarPlugins
    ]);
  });
}