import { createContext, FC } from "react";

export type PluginType = 'SidebarPlugin' | 'FeaturePlugin' | 'ContextPlugin';

export type Plugin = {
  type: PluginType;
}

export type SidebarPlugin = Plugin & {
  component: JSX.Element;
}

export type ContextPlugin = Plugin & {
  id: string;
  Provider: FC;
}

export type FeaturePlugin = Plugin & {
  id: string;
  menuItems: MenuItem[];
  pages: Page[];
}

export type MenuItem = {
  label: string; 
  path: string;
  icon?: JSX.Element;
  shouldShow?: () => boolean;
}

export type Page = {
  component: JSX.Element;
  path: string;
  shouldShow?: () => boolean;
}

export type Breakpoint = 'sm'|'md'|'lg'|'xl';

export type Theme = {
  colors: {
    [key: string]: string;
  },
  breakpoints: {
    [key in Breakpoint]: string;
  }
}

type CoreContextType = {
  menuItems: MenuItem[];
  pages: Page[];
}

export const CoreContext = createContext<CoreContextType>({
  menuItems: [],
  pages: [],
} as any);