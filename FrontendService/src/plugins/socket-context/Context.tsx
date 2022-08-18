import { createContext, FC, useContext, useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { AuthContext, SOCKET_SERVICE_URL, useNotificationService } from "./imports";


export type SocketContextType = {
  client: any
}

const initialValues = {
  client: undefined
}


var globalLocation: any;

export const SocketContext = createContext<SocketContextType>(initialValues);

export const SocketContextProvider: FC = ({ children }) => {

  const { user } = useContext(AuthContext);
  
  const [client, setClient] = useState<any>();
  const notificationService = useNotificationService();
  const location = useLocation();
  globalLocation = location;

  const removeListeners = () => {
    client?.removeAllListeners('statuses');
    client?.removeAllListeners('messages_sidebar');
    client?.removeAllListeners('messages_chat');
    client?.removeAllListeners('messages_global');
    client?.removeAllListeners('messages_delete');
    client?.removeAllListeners('chat_delete');
    client?.removeAllListeners('disconnect');
    (window as any).io.removeAllListeners && (window as any).io.removeAllListeners('connection');
  };

  useEffect(() => {
    if (!user) return;

    const client = (window as any).io.connect(`${SOCKET_SERVICE_URL}`, {
      transports: ['websocket']
    });


    client.on('connect', () => {
      console.log(`Connected to socket server.`);
    });

    client.on('connection', (client: any) => {
      client.join(user?.ID)
    });

    client.on('messages_global', function(data: any) {
      if (globalLocation.pathname !== '/chat') {
        notificationService.info(`You have a new message.`);
      }
    });

    client.on('disconnect', () => {
      removeListeners();
      console.log(`Disconnected from socket server.`);
    });

    const interval = setInterval(() => {
      client.emit("heartbit", { text: `{ "UserId": "${user?.ID}" }` });
      console.log(`Heartbit sent for ${user?.ID}`);
      }, 5000)

    setClient(client);

    return () => { clearInterval(interval); removeListeners(); client?.close(); }
  }, [user]);

  return (
    <SocketContext.Provider value={{ client }}>
      {children}
    </SocketContext.Provider>
  );
}