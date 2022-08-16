import axios from "axios";
import { createContext, FC, useEffect, useState } from "react";
import {
  Profile,
  UsersService,
} from './imports';
import {
  getToken,
  getUserIdFromToken,
  setAxiosInterceptors,
} from './utils';

type AuthContextValue = {
  isAuthenticated: boolean;
  user?: Profile,
  setUser: (u: Profile|undefined) => any,
  setAuthenticated: (a: boolean) => any
}

export const AuthContext = createContext<AuthContextValue>({
  isAuthenticated: false,
  user: undefined,
  setUser: () => {},
  setAuthenticated: () => {}
});

export const AuthContextProvider: FC = ({ children }) => {
  const [isAuthenticated, setAuthenticated] = useState<boolean>(!!getToken());
  const [user, setUser] = useState<Profile>();

  const configureInterceptors = () => {
    setAxiosInterceptors(axios, () => {
      setAuthenticated(false);
      window.location.href = "/";
      sessionStorage.removeItem('jwt-token');
    });
  }

  configureInterceptors();

  useEffect(() => {
    if(isAuthenticated) {
      const userId = getUserIdFromToken();
      const userService = new UsersService();
      userService.fetch(userId)
        .then(setUser)
        .catch(console.log);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated]);

  return (
    <AuthContext.Provider value={{ user, setUser, isAuthenticated, setAuthenticated }}>
      {children}
    </AuthContext.Provider>
  );
}