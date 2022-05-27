import { FC, useEffect, useState } from "react";
import { ThemeProvider } from "react-jss";
import { BrowserRouter } from "react-router-dom";
import { TreeListWrapper } from "./components/TreeListWrapper";
import { ContextPlugin, CoreContext, FeaturePlugin, MenuItem, Page, Plugin, SidebarPlugin, Theme } from "./Context";
import { Routes } from "./Routes";

export type InitParams = {
  registerPlugin: (plugin: Plugin) => any;
  registerAllPlugins: (plugins: Plugin[]) => any;
}

type Props = {
  initialize: (params: InitParams) => any;
}

const defaultTheme: Theme = {
  colors: {
    primary: '#00A7EC',
    secondary: '#118CFF',
    warning: '#FFC700',
    danger: '#FF513A',
    textLight: '#fafdff',
    textDark: '#757678',
  },
  breakpoints: {
    sm: `@media (min-width: 576px)`,
    md: `@media (min-width: 768px)`,
    lg: `@media (min-width: 992px)`,
    xl: `@media (min-width: 1200px)`,
  }
}

type CoreState = {
  registeredFeaturePluginIds: string[];
  menuItems: MenuItem[];
  pages: Page[];
  Sidebar: JSX.Element;
  registeredContextPluginIds: string[];
  providers: FC[];
}

export const Core: FC<Props> = ({ initialize }) => {
  const [theme] = useState<Theme>(defaultTheme);

  const [state, setState] = useState<CoreState>({
    registeredFeaturePluginIds: [],
    menuItems: [],
    pages: [],
    Sidebar: <></>,
    registeredContextPluginIds: [],
    providers: []
  });

  const { registeredFeaturePluginIds, menuItems, pages, Sidebar, registeredContextPluginIds, providers } = state;

  const registerPlugin = (plugin: Plugin) => {
    if(plugin.type === "FeaturePlugin") {
      const featurePlugin = plugin as FeaturePlugin;
      if(registeredFeaturePluginIds.includes(featurePlugin.id)) return;
      setState(prevState => ({
        ...prevState,
        registeredFeaturePluginIds: [...prevState.registeredFeaturePluginIds, featurePlugin.id],
        menuItems: [...prevState.menuItems, ...featurePlugin.menuItems],
        pages: [...prevState.pages, ...featurePlugin.pages],
      }));
    } else if(plugin.type === "SidebarPlugin") {
      const sidebarPlugin = plugin as SidebarPlugin;
      setState(prevState => ({
        ...prevState,
        Sidebar: sidebarPlugin.component,
    }));
    } else if(plugin.type === "ContextPlugin") {
      const contextPlugin = plugin as ContextPlugin;
      if(registeredContextPluginIds.includes(contextPlugin.id)) return;
      setState(prevState => ({
        ...prevState,
        registeredContextPluginIds: [...prevState.registeredContextPluginIds, contextPlugin.id],
        providers: [...prevState.providers, contextPlugin.Provider],
      }));
    }
  }

  const registerAllPlugins = (plugins: Plugin[]) => {
    const featurePlugins = plugins.filter(p => p.type === "FeaturePlugin") as FeaturePlugin[];
    const sidebarPlugin = plugins.find(p => p.type === "SidebarPlugin") as SidebarPlugin;
    const contextPlugins = plugins.filter(p => p.type === "ContextPlugin") as ContextPlugin[];
    setState(prevState => ({
      ...prevState,
      registeredFeaturePluginIds: featurePlugins.map(p => p.id),
      menuItems: featurePlugins.reduce<MenuItem[]>((acc, p) => [...acc, ...p.menuItems], []),
      pages: featurePlugins.reduce<Page[]>((acc, p) => [...acc, ...p.pages], []),
      Sidebar: sidebarPlugin?.component,
      registeredContextPluginIds: contextPlugins.map(p => p.id),
      providers: contextPlugins.map(p => p.Provider),
    }));
  }

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => initialize({ registerPlugin, registerAllPlugins }), []);

  return (
    <CoreContext.Provider value={{ menuItems, pages }}>
      <ThemeProvider theme={theme}>
        <BrowserRouter>
          <TreeListWrapper wrappers={providers}>
            {Sidebar}
            <Routes />
          </TreeListWrapper>
        </BrowserRouter>
      </ThemeProvider>
    </CoreContext.Provider>
  );
}