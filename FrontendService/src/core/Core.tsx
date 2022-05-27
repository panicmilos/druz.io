import { FC, useEffect, useState } from "react";
import { ThemeProvider } from "react-jss";
import { BrowserRouter } from "react-router-dom";
import { TreeListWrapper } from "./components/TreeListWrapper";
import { ContextPlugin, CoreContext, FeaturePlugin, MenuItem, Page, Plugin, SidebarPlugin, Theme } from "./Context";
import { Routes } from "./Routes";

export type InitParams = {
  registerPlugin: (plugin: Plugin) => any;
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

export const Core: FC<Props> = ({ initialize }) => {
  const [theme] = useState<Theme>(defaultTheme);

  const [registeredFeaturePluginIds, setRegisteredFeaturePluginIds] = useState<string[]>([]);
  const [menuItems, setMenuItems] = useState<MenuItem[]>([]);
  const [pages, setPages] = useState<Page[]>([]);
  
  const [Sidebar, setSidebar] = useState<JSX.Element>(<></>);

  const [registeredContextPluginIds, setRegisteredContextPluginIds] = useState<string[]>([]);
  const [providers, setProviders] = useState<FC[]>([]);

  const registerPlugin = (plugin: Plugin) => {
    if(plugin.type === "FeaturePlugin") {
      const featurePlugin = plugin as FeaturePlugin;
      if(registeredFeaturePluginIds.includes(featurePlugin.id)) return;
      setRegisteredFeaturePluginIds(prevPluginIds => [...prevPluginIds, featurePlugin.id]);
      setMenuItems(prevMenuItems => [...prevMenuItems, ...featurePlugin.menuItems]);
      setPages(prevPages => [...prevPages, ...featurePlugin.pages]);
    } else if(plugin.type === "SidebarPlugin") {
      const sidebarPlugin = plugin as SidebarPlugin;
      setSidebar(sidebarPlugin.component);
    } else if(plugin.type === "ContextPlugin") {
      const contextPlugin = plugin as ContextPlugin;
      if(registeredContextPluginIds.includes(contextPlugin.id)) return;
      setRegisteredContextPluginIds(prevPluginIds => [...prevPluginIds, contextPlugin.id]);
      setProviders(prevProviders => [...prevProviders, contextPlugin.Provider]);
    }
  }

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => initialize({ registerPlugin }), []);

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