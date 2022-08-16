import { AxiosError } from "axios";
import { FC, useContext, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useMutation } from "react-query";
import { useUserBlocksResult } from "../../hooks";
import { AuthContext, Button, ConfirmationModal, extractErrorMessage, Table, TableBody, TableHead, TableRow, useNotificationService } from "../../imports";
import { UserBlock } from "../../models/UserBlock";
import { useUserBlocksService } from "../../services";

type Props = {
  userBlocks: UserBlock[]
}

const useStyles = createUseStyles({
  container: {
    '& button': {
      margin: '0em 0.5em 0.5em 0.5em'
    }
  }
});

export const BlockedUsersTable: FC<Props> = ({ userBlocks }) => {

  const [isUnblockUserOpen, setIsUnblockUserOpen] = useState(false);
  const [selectedUserBlock, setSelectedUserBlock] = useState<UserBlock|undefined>();

  const { user } = useContext(AuthContext);

  const userBlocksService = useUserBlocksService(user?.ID ?? '');
  const notificationService = useNotificationService();
  
  const { result, setResult } = useUserBlocksResult();

  const unblockUserMutation = useMutation(() => userBlocksService.delete(selectedUserBlock?.BlockedId + ''), {
    onSuccess: () => {
      notificationService.success('You have successfully unblocked user.');
      setResult({ status: 'OK', type: 'UNBLOCK_USER' });
    },
    onError: (error: AxiosError) => {
      notificationService.error(extractErrorMessage(error.response?.data));
      setResult({ status: 'ERROR', type: 'UNBLOCK_USER' });
    }
  });
  const unblockUser = () => unblockUserMutation.mutate();

  useEffect(() => {
    if (!result) return;

    if (result.status === 'OK' && result.type === 'UNBLOCK_USER') {
      setIsUnblockUserOpen(false);
    }

  }, [result]);

  const ActionsButtonGroup = ({ userBlok }: any) =>
    <>
      <Button onClick={() => { setSelectedUserBlock(userBlok); setIsUnblockUserOpen(true); }}>Unlbock</Button>
    </>
    
  const classes = useStyles();

  return (
    <div className={classes.container}>

      <ConfirmationModal title="Unlblock User" open={isUnblockUserOpen} onClose={() => setIsUnblockUserOpen(false)} onYes={unblockUser}>
        <p>Are you sure you want to unblock this user?</p>
      </ConfirmationModal>

      <Table hasPagination={false}>
        <TableHead columns={['Blocked', 'Action']}/>
        <TableBody>
          {
            userBlocks?.map((userBlock: UserBlock) => 
            <TableRow 
              key={userBlock.ID}
              cells={[
                `${userBlock.Blocked.FirstName} ${userBlock.Blocked.LastName}`,
                <ActionsButtonGroup userBlok={userBlock}/>
            ]}/>
            )
          }
        </TableBody>
      </Table>
    </div>
    );
  }