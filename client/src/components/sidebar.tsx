import { Users, Search } from 'lucide-react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../state/store';
import { setActiveTab } from '../state/room/roomSlice';

const Sidebar = () => {
    const dispatch = useDispatch();
    const activeTab = useSelector((state: RootState) => state.room.activeTab);
  
    return (
      <div className="h-full w-16 bg-gray-800 text-white py-4 flex flex-col items-center space-y-6 shadow-md">
  
        <div className="relative group">
          <button
            onClick={() => dispatch(setActiveTab("joined"))}
            className={`p-3 rounded-lg transition 
              ${activeTab === "joined" ? "bg-blue-600" : "hover:bg-gray-700"}`}
          >
            <Users size={20} />
          </button>
          <span className="absolute left-16 top-1/2 -translate-y-1/2 px-3 py-1 bg-gray-700 text-sm rounded-md opacity-0 group-hover:opacity-100 transition whitespace-nowrap z-10">
            Joined Rooms
          </span>
        </div>
  
        <div className="relative group">
          <button
            onClick={() => dispatch(setActiveTab("discover"))}
            className={`p-3 rounded-lg transition 
              ${activeTab === "discover" ? "bg-blue-600" : "hover:bg-gray-700"}`}
          >
            <Search size={20} />
          </button>
          <span className="absolute left-16 top-1/2 -translate-y-1/2 px-3 py-1 bg-gray-700 text-sm rounded-md opacity-0 group-hover:opacity-100 transition whitespace-nowrap z-10">
            Discover Rooms
          </span>
        </div>
  
      </div>
    );
  };
  
  export default Sidebar;