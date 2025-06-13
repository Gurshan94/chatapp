import MessageList from "./messages"
import { useSelector,useDispatch } from "react-redux"
import { RootState } from "../state/store"
import { useSocket } from "../utils/socketProvider"
import { useEffect,useState} from "react"
import { addMessage, outMessage } from "../state/room/roomSlice"
import { Message } from "../state/room/roomSlice"
const Messagearea=()=>{

    const roomid=useSelector((state:RootState)=>state.room.selectedRoomId)
    const userId=useSelector((state:RootState)=>state.user.user?.id)
    const {socket}=useSocket();
    const dispatch = useDispatch();
    const [message, setMessage] = useState(''); // State for message input


    useEffect(() => {
        if (!socket) return;
    
        const handleMessage = (event: MessageEvent) => {
          try {
            const data:Message = JSON.parse(event.data);
    
            dispatch(addMessage(data));
            console.log("message dispatched")
          } catch (err) {
            console.error('Failed to parse message:', err);
          }
        };
    
        socket.addEventListener('message', handleMessage);
    
        return () => {
          socket.removeEventListener('message', handleMessage);
        };
      }, [socket, dispatch]);


      const handleMessageSend = () => {
        if (socket) {
          // Create message object
          const newMessage:outMessage = {
            senderid: Number(userId),
            content: message,
            roomid: roomid, // You might want to use dynamic room ID depending on the room you're in
          };
    
          // Send the message through WebSocket
          socket.send(JSON.stringify(newMessage));     
          // Clear the input after sending the message
          setMessage('');
        }
      };


    return (
        <div className="h-[calc(100vh-68px)] flex flex-col flex-1">
            
        <div className="flex-1 overflow-y-auto p-4" id="scrollable-message-container">
            <MessageList roomid={roomid}/>
        </div>

        <div className="p-4 border-t border-gray-700 bg-gray-800">
        <div className="flex gap-2">
            <input
            type="text"
            placeholder="Type your message..."
            className="flex-1 p-3 rounded-lg bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            onKeyDown={(e) => e.key === 'Enter' && handleMessageSend()}
            />
            <button
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            onClick={handleMessageSend}
            >
            Send
            </button>
        </div>
        </div>
        </div>
    )
}

export default Messagearea