import { FC, useContext } from "react";
import { Route, Routes as RoutesGroup } from "react-router-dom";
import { CoreContext, Page } from "./Context";

type Props = {}

const PageRoute: FC<{ page: Page }> = ({ page }) => {
  const shouldShow = !page.shouldShow || page.shouldShow();

  return shouldShow ? page.component : <></>;
}

export const Routes: FC<Props> = () => {

  const { pages } = useContext(CoreContext);

  return (
    <RoutesGroup>
      {pages.map((p, i) => <Route key={p.path + i} path={p.path} element={<PageRoute page={p} />} />)}
    </RoutesGroup>
  );
}