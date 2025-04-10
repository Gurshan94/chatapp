import { useDispatch,useSelector } from "react-redux";
import { useGetroomsjoinedbyuserMutation } from "../state/room/roomApiSlice";
import { AppDispatch, RootState } from "../state/store";
import { openRoom, setJoinedRooms } from "../state/room/roomSlice";
import { useEffect, useState } from "react";
import RoomCard from "./roomcard";
import { Room } from "../state/room/roomSlice";

  const RoomList = () => {
    
    const dispatch=useDispatch<AppDispatch>()
    const userid=useSelector((state:RootState)=>state.user.user?.id)
    const activeTab=useSelector((state:RootState)=>state.room.activeTab)

    const [fetchjoinedrooms]=useGetroomsjoinedbyuserMutation()
    const [Error, setError] = useState("");

    useEffect(()=>{

          const fetchRooms = async () => {
            try {
              const data = await fetchjoinedrooms({ userid }).unwrap();
              dispatch(setJoinedRooms(data));
            } catch (err: any) {
              setError(err.data?.message || 'couldn’t fetch rooms');
            }
          };
          fetchRooms();
      
    },[userid, activeTab])

    const handleOpenroom=(roomid:number)=>{
      

    }
  

    const allRooms=useSelector((state:RootState)=>state.room.joinedRooms)

    return (
      <div className="flex flex-col h-full w-72 bg-gray-850 border-r border-gray-700">
        <div className="flex-1 overflow-y-auto p-4 custom-scrollbar">
        <h2 className="text-xl font-semibold mb-4">Rooms</h2>
          <div className="flex flex-col gap-2">
            {allRooms.map((room: Room) => (
              <div
                key={room.id}
                onClick={() => handleOpenroom(room.id)}
                className="cursor-pointer"
              >
                <RoomCard room={room} />
              </div>            
          ))}        
          </div>
        </div>
     

      <div className="h-20" />
    </div>
    );
  };
  
  export default RoomList;