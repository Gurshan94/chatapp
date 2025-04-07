import Sidebar from '../components/sidebar';
import RoomList from '../components/roomlist';
import Messagearea from '../components/messagearea.tsx';
import { useSelector } from 'react-redux';
import { RootState } from '../state/store.ts';

const Chat = () => {

  const openedroomid=useSelector((state:RootState)=>state.room.selectedRoomId)

  return (
    <div className="flex h-screen bg-gray-900">
        <Sidebar />
        <RoomList />
        {openedroomid ? 
        <Messagearea/>:
        <div className="flex flex-1 flex-col items-center justify-center text-center text-gray-400 px-4">
          <h2 className="text-2xl font-semibold text-gray-300 mb-2">Welcome to ChatZone</h2>
          <p className="text-base">Select a chat to start messaging</p>
        </div>}
    </div>
  );
};

export default Chat;