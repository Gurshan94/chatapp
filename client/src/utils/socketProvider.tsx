import { useContext,createContext, useState, ReactNode } from 'react';

// Define types for the context
interface SocketContextType {
  socket: WebSocket | null;
  connectSocket: (userId: number,token:string) => void;
}

// Create the context with default values (null for socket, empty function for connectSocket)
const SocketContext = createContext<SocketContextType>({
  socket: null,
  connectSocket: () => {}, // default empty function
});

interface SocketProviderProps {
  children: ReactNode;
}

export const SocketProvider = ({ children }: SocketProviderProps) => {
  const [socket, setSocket] = useState<WebSocket | null>(null);

  // Function to establish WebSocket connection with userId
  const connectSocket = (userId: number,token:string) => {
    const ws = new WebSocket(`ws://localhost:8080/ws/${userId}?token=${token}`);
    console.log("websocket connection made")
    // Handle WebSocket events
    setSocket(ws); // Store the WebSocket instance

    // Cleanup on unmount
    return () => {
      if (ws) {
        ws.close();
      }
    };
  };

  return (
    <SocketContext.Provider value={{ socket, connectSocket }}>
      {children}
    </SocketContext.Provider>
  );
};

// Export SocketContext and SocketProvider for use in other components
export default SocketContext;
export const useSocket = () => useContext(SocketContext);
