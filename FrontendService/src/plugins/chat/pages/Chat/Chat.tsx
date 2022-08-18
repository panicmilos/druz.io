import { FC, useState } from "react";
import { createUseStyles } from "react-jss";
import { useQuery } from "react-query";
import { Button, Modal } from "../../imports";
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
    minWidth: '30%',
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
  const [selectedChat, setSelectedChat] = useState<any>(undefined);
  const [fetchChats, setFetchChats] = useState(true);
  

  const chatService = useChatService();

  const { data: chatsReversed } = useQuery([chatService, fetchChats], () => chatService.fetch(), {
    onSuccess: (chats) => {
      chats = chats?.reverse() || [];

      if (!chats || !chats[0]) return;
      setSelectedChat({ chatId: chats[0].Chat, friendId: chats[0].User.ID.replace('users/', '') });
    }
  });
  const chats = chatsReversed?.reverse();
  

  console.log(chats);

  const classes = useStyles();

  return (
    <>
      <div className={classes.container}>

        <Modal title={"New Chat"} open={isAddNewChatOpen} onClose={() => setIsAddNewChatOpen(false)}>
          <SelectUserToChatForm
            onSubmit={(userId: string) => {
              
              const existingChat = chats?.find(chat => fixId(chat.User.ID) === userId);

              if (existingChat) {
                setSelectedChat({ chatId: existingChat.Chat, friendId: fixId(existingChat.User.ID) });
              } else {
                setSelectedChat({ chatId: 'NOT_CREATED_YET', friendId: userId });
              }

              setIsAddNewChatOpen(false);
            }}
          />
        </Modal>

        <div className={classes.buttons}>
          <Button onClick={() => { setIsAddNewChatOpen(true)} }>New Chat</Button>         
        </div>
      </div>

      <div className={classes.parentContainer}>

        <div className={`${classes.firstContainer} ${classes.twoContainer}`}>
          {
            chats?.map(chat => {

              return (
                <Button
                  key={chat.Chat}
                  onClick={() => setSelectedChat({ chatId: chat.Chat, friendId: fixId(chat.User.ID) })}
                >
                  {chat.Chat} {chat.User.FirstName} {chat.User.LastName} { selectedChat.friendId === fixId(chat.User.ID) ? 'Selected' : '' }
                </Button>
              )
            })
          }
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