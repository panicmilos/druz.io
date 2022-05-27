import { FC } from "react";
import { SidebarPlugin } from "./imports";
import { Sidebar } from "./Sidebar";

export const PaddingContainer: FC = ({ children }) => <div style={{paddingTop: '3em'}}>{children}</div>

export function getPluginDefinition(): SidebarPlugin {
  return {
    type: 'SidebarPlugin',
    component: <Sidebar />
  }
}