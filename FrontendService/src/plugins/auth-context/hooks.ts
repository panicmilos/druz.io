import { useContext } from "react";
import { AuthContext } from "./Context";

export type IsAuthorizedParams = {
  roles?: string[];
};

export type IsAuthorizedFunction = (params?: IsAuthorizedParams) => boolean;

export function useIsAuthorized(): IsAuthorizedFunction {
  const { isAuthenticated, user } = useContext(AuthContext);

  return ({ roles = undefined } = {}) => {
    const userRole = user?.role ?? '';

    const hasOneOfRoles =
      !roles ||
      (!!userRole &&
      roles.reduce<boolean>(
          (acc, curr) => acc || curr === userRole,
          false
        ));

    return hasOneOfRoles && isAuthenticated;
  };
}

export type IsUnauthorizedFunction = () => boolean;

export function useIsUnauthorized(): IsUnauthorizedFunction {
  const { isAuthenticated } = useContext(AuthContext);

  return () => !isAuthenticated;
}