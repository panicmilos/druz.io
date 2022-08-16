import { FC, useContext } from "react"
import { AuthContext, Container } from "../../imports"
import { Profile } from "../../models/User";
import { ChangePasswordForm } from "./ChangePasswordForm";
import { DeactivateProfile } from "./DeactivateProfile";

export const ChangeProfile: FC = () => {

  const { user } = useContext(AuthContext);

  return (
    <>
      <Container flexDirection="column">
        
        <ChangePasswordForm user={user as Profile} />

        <DeactivateProfile user={user as Profile} />

      </Container>
    </>
  )
}