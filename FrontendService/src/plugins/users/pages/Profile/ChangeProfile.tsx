import { FC, useContext } from "react"
import { AuthContext, Container } from "../../imports"
import { Profile } from "../../models/User";
import { ChangePasswordForm } from "./ChangePasswordForm";

export const ChangeProfile: FC = () => {

  const { user } = useContext(AuthContext);

  return (
    <>
      <Container>
        
        <ChangePasswordForm user={user as Profile} />

      </Container>
    </>
  )
}