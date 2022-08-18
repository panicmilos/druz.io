import { FC, useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext, USER_ROLE } from "../imports";
import { LoginForm } from "./LoginForm";

type Props = {};

export const Login: FC<Props> = () => {

  const nav = useNavigate();

  const { user } = useContext(AuthContext);

  useEffect(() => {

    if (!user) return;

    if (user.Role === USER_ROLE) {
      nav('/posts');
    } else {
      nav('/users/reports')
    }
  }, [user]);
  return (
    <>
      <LoginForm />
    </>
  )
}