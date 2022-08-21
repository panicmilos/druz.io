import { FC, useContext, useEffect, useState } from "react";
import { createUseStyles } from "react-jss";
import { useQuery } from "react-query";
import { useStatusesMap, useUserFriendsMap } from "../../hooks";
import { Button, Modal, Profile, SocketContext } from "../../imports";
import { Chat as ChatT } from "../../models/Chat";
import { Message } from "../../models/Message";
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
  },
  userButton: {
    borderRadius: '0px',
    textAlign: 'left'
  },
  notSelected: {
    backgroundColor: '#64caf5'
  },
  nameContainer: {
    display: 'flex',
    alignItems: 'center',
    '& p': {
      marginLeft: '5px',
    },
    '& img': {
      width: '36px',
      height: '36px',
      borderRadius: '50%'
    }
  },
});

const fixId = (id: string) => id.replace('users/', '');

var globalChats: ChatT[] = [];
var globalNotificationsMap: any = {};
var globalUserFriendsMap: any = {};
var globalSeletedChat: any = {};

export const Chat: FC = () => {

  const [isAddNewChatOpen, setIsAddNewChatOpen] = useState(false);
  const [chats, setChats] = useState<ChatT[]>([]);
  globalChats = chats;
  const [selectedChat, setSelectedChat] = useState<any>(undefined);
  globalSeletedChat = selectedChat;
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
  const [notificationsMap, setNotificationsMap] = useState<any>({})
  globalNotificationsMap = notificationsMap;

  const { client } = useContext(SocketContext);
  const userFriendsMap = useUserFriendsMap();
  globalUserFriendsMap = userFriendsMap;

  useEffect(() => {
    if (!client) return;

    client?.on('messages_sidebar', function(data: any) {
      const message = JSON.parse(data.text).Message as Message;
      const messageChatId = message.ID.split('/')[1];

      if (messageChatId === globalSeletedChat?.chatId) { return; }

      if (!globalChats.find(c => c.Chat === messageChatId)) {
        setChats([{ Chat: messageChatId, User: globalUserFriendsMap[message.FromId] }, ...globalChats]);
      }
      setNotificationsMap({...globalNotificationsMap, [messageChatId]: (globalNotificationsMap[messageChatId] || 0) + 1 })
    });

    return () => { client.removeAllListeners('messages_sidebar'); }
  }, [client]);

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

                const chatId = chat.Chat;
                const friendId = chat.User.ID;
                const formName = (user: Profile) => `${user.FirstName} ${user.LastName} `;

                return (                    
                    <Button
                      key={chatId}
                      className={`${classes.userButton} ${selectedChat?.friendId !== friendId ? classes.notSelected : ''}`}
                      onClick={() => { setSelectedChat({ chatId: chatId, friendId: friendId }); setNotificationsMap({...notificationsMap, [chatId]: 0 }); }}
                    >
                      <div className={classes.nameContainer}>
                      <img src={chat.User?.Image || '/images/no-image.png'}></img>
                      <p>
                        {formName(chat.User)}
                        {notificationsMap[chatId] ? `(${notificationsMap[chatId]})`: ``}
                        <span style={{margin: '0px 5px 0px 5px', fontSize: '20px', color: statusesMap[friendId] === 'online' ? 'green': 'red'}}>●</span>
                      </p>
                      </div>
                    </Button>
                )
              })
            }
          </div>
        </div>

        <div className={`${classes.growFlex}`}>
          {
            selectedChat ?
              <ChatMessages key={selectedChat?.chatId} chat={selectedChat} onInitial={() => { setTimeout(() => { setFetchChats(!fetchChats) }, 300); setSelectedChat(undefined); }} /> :
              <></>
          }
        </div>
        
      </div>
    </>
  )
}