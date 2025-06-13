import Sidebar from '../components/sidebar';
import RoomList from '../components/joinedrooms.tsx';
import Messagearea from '../components/messagearea.tsx';
import DiscoverRoomList from '../components/discoverrooms.tsx';
import { useSelector } from 'react-redux';
import { RootState } from '../state/store.ts';
import { useEffect } from 'react';
import { useGetMeQuery } from '../state/user/userApiSlice.ts';
import { useDispatch } from "react-redux";
import { AppDispatch } from "../state/store"; // adjust this if needed
import { setCredentials } from "../state/user/authSlice";
import CreateRoomDialog from '../components/createroom.tsx';

const Chat = () => {

  const dispatch = useDispatch<AppDispatch>();
  const { data, isSuccess } = useGetMeQuery(); // auto login if the user refreshes

  useEffect(() => {
    if (isSuccess && data) {
      dispatch(setCredentials(data));
    }
  }, [isSuccess, data, dispatch]);

  const activeTab = useSelector((state: RootState) => state.room.activeTab); // set active tag
  const openedroomid=useSelector((state:RootState)=> state.room.selectedRoomId) // get the opened room


  return (
    <div className="flex h-screen bg-gray-900">
        <Sidebar />
        {activeTab=='joined'?<RoomList />:activeTab=="discover"?<DiscoverRoomList/>:<CreateRoomDialog/>}
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