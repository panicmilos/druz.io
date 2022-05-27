import { FC } from "react";
import { useIsAuthorized } from "../auth-context/hooks";

type Props = {
  roles?: string[];
};

export const Authorized: FC<Props> = ({
  children,
  roles = undefined,
}) => {
  const isAuthorized = useIsAuthorized();

  return isAuthorized({ roles }) ? (
    <>{children}</>
  ) : (
    <></>
  );
};
