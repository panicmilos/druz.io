import { FC, useContext } from "react"
import { AuthContext, Card, Container } from "../../imports"
import { Profile } from "../../models/User";
import { ChangeImageForm } from "./ChangeImageForm";
import { ChangePasswordForm } from "./ChangePasswordForm";
import { DeactivateProfile } from "./DeactivateProfile";
import { ProfileForm } from "./ProfileForm";

export const ChangeProfile: FC = () => {

  const { user } = useContext(AuthContext);

  return (
    <>
      <Container flexDirection="column">
        
        <Card title="Change Profile">
          <ChangeImageForm user={user as Profile} />

          <ProfileForm user={user as Profile} />
        </Card>
          
        <Card title="Change Password">
          <ChangePasswordForm user={user as Profile} />
        </Card>

        <Card title="Disable Profile">
          <DeactivateProfile user={user as Profile} />
        </Card>

      </Container>
    </>
  )
}