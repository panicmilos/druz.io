import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { useQuery } from "react-query";
import { useStatusesMap } from "../../hooks";
import { Button, Modal, Profile } from "../../imports";
import { Chat as ChatT } from "../../models/Chat";
import { useChatService } from "../../services";
import { ChatMessages } from "./ChatMessages";
import { SelectUserToChatForm } from "./SelectUserToChatForm";

const useStyles = createUseStyles({
  container: {
    margin: '1% 3% 0% 3%',
    '& button': {
      margin: '0.5em 0.5em 0.5em 0.5em'
    },
  },
  buttons: {
    display: 'flex',
    marginTop: '15px'
  },
  parentContainer: {
    display: 'flex',
    width: '100%'
  },
  twoContainer: {
    minWidth: '20%',
    maxWidth: '100%'
  },
  firstContainer: {
    marginRight: '2%',
  },
  growFlex: {
    flexGrow: 1
  }
});

const fixId = (id: string) => id.replace('users/', '');

export const Chat: FC = () => {

  const [isAddNewChatOpen, setIsAddNewChatOpen] = useState(false);
  const [chats, setChats] = useState<ChatT[]>([]);
  const [selectedChat, setSelectedChat] = useState<any>(undefined);
  const [fetchChats, setFetchChats] = useState(true);
  
  const chatService = useChatService();

  useQuery([chatService, fetchChats], () => chatService.fetch(), {
    onSuccess: (chatsReversed) => {
      const chats = chatsReversed?.reverse().map((c) => { c.User.ID = fixId(c.User.ID); return c; });

      if (!chats || !chats[0]) return;

      setChats(chats);
      setSelectedChat({ chatId: chats[0].Chat, friendId: chats[0].User.ID });
    }
  });
  
  const classes = useStyles();

  const statusesMap = useStatusesMap();

  // const { client } = useContext(SocketContext);
  // client?.on('messages', function (data: any) {
  //   console.log('Received message', data);
  // })

  return (
    <>

      <Modal title={"New Chat"} open={isAddNewChatOpen} onClose={() => setIsAddNewChatOpen(false)}>
        <SelectUserToChatForm
          onSubmit={(userId: string) => {
            const existingChatId = chats?.find(chat => fixId(chat.User.ID) === userId)?.Chat;
            setSelectedChat({ chatId: existingChatId ? existingChatId : 'NOT_CREATED_YET', friendId: userId });

            setIsAddNewChatOpen(false);
          }}
        />
      </Modal>


      <div className={classes.parentContainer}>

        <div className={`${classes.firstContainer} ${classes.twoContainer}`} >
          <div className={classes.buttons} style={{ marginTop: '-0.10em', marginBottom: '0.7em' }}>
            <Button onClick={() => { setIsAddNewChatOpen(true)} }>New Chat</Button>         
          </div>

          <div style={{display: 'flex', flexDirection: 'column'}}>
            {
              chats?.map(chat => {

                const friendId = chat.User.ID;
                const formName = (user: Profile) => `${user.FirstName} ${user.LastName}`;

                return (                    
                    <Button
                      key={chat.Chat}
                      onClick={() => setSelectedChat({ chatId: chat.Chat, friendId: friendId })}
                    >
                      {formName(chat.User)} ({ statusesMap[friendId] || 'offline'}) { selectedChat?.friendId === friendId ? 'Selected' : '' }
                    </Button>
                )
              })
            }
          </div>
        </div>

        <div className={`${classes.growFlex}`}>
          {
            selectedChat ?
              <ChatMessages key={selectedChat?.chatId} chat={selectedChat} onInitial={() => { setFetchChats(!fetchChats); setSelectedChat(undefined); }} /> :
              <></>
          }
        </div>
        
      </div>
    </>
  )
}