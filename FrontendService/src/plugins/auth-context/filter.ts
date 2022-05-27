import { IsAuthorizedParams, useIsAuthorized, useIsUnauthorized } from "./hooks";

export const authorizedFor = (params?: IsAuthorizedParams): () => boolean => {
  return () => {
    const isAuthorized = useIsAuthorized();
    return isAuthorized(params);
  }
}

export const unauthorized = (): () => boolean => {
  return () => {
    const isUnauthorized = useIsUnauthorized();
    return isUnauthorized();
  }
}