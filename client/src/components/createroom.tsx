import { motion } from "framer-motion";
import { useState } from "react";
import {useDispatch, useSelector} from "react-redux";
import { RootState } from "../state/store";
import { useCreateroomMutation, useJoinroomMutation } from "../state/room/roomApiSlice";
import { BackendRoom } from "../state/room/roomSlice";
import {joinRoom} from "../state/room/roomSlice"

export default function CreateRoomDialog() {
    const dispatch=useDispatch()
    const [roomname, setRoomName] = useState("");
    const [maxusers, setMaxUsers] = useState<number>(2);
    const [Error, setError] = useState("");
    const [Success, setSuccess] = useState("");
    const userid =useSelector((state:RootState)=>state.user.user?.id)

    const [createroom] = useCreateroomMutation()
    const [joinRoomMutation] = useJoinroomMutation()

    const handleCreateRoom = async (e: React.FormEvent) => {
        e.preventDefault();
        
        console.log("Creating room:");
        try {
            if (userid!==undefined){
                console.log(roomname, maxusers, userid);
                const roomData:BackendRoom = await createroom({roomname,maxusers,adminid:Number(userid)}).unwrap()
                if (roomData) {
                    var roomid:number=roomData.id;
                    await joinRoomMutation({roomid,userid:Number(userid)}).unwrap();
                    dispatch(joinRoom({...roomData}));
                    setSuccess("Room created successfully!"); // Set success message
                }
            }
            
        } catch (error) {
            setError("Failed to create room. Please try again.");      
        }
    };

return (
    <div className="min-h-screen bg-gray-900 text-white flex items-center justify-center px-4">
    <motion.div
        className="bg-gray-800 p-8 rounded-2xl shadow-lg w-full max-w-md"
        initial={{ opacity: 0, y: 40 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
    >
        <h2 className="text-3xl font-bold text-center mb-6">Create New Room </h2>

        {Error && (
        <p className="text-red-400 text-sm text-center mb-4">{Error}</p>
        )}

        {Success && (
        <p className="text-green-400 text-sm text-center mb-4">{Success}</p>
        )}

        <form className="space-y-4" onSubmit={handleCreateRoom}>
        <div>
            <label className="block mb-1 text-sm">RoomName</label>
            <input
            type="text"
            placeholder="room_1"
            value={roomname}
            onChange={(e) => setRoomName(e.target.value)}
            required
            className="w-full px-4 py-2 rounded bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
        </div>

        <div className="flex flex-col space-y-2 w-full">
        <label className="text-sm font-medium text-white">
            Max Users: {maxusers} {/* Display current slider value */}
        </label>
        <input
            type="range"
            min={2} // Minimum room size
            max={20} // Maximum room size
            step={1} // Increment of 1
            value={maxusers} // Bind the value of the slider to maxUsers state
            onChange={(e) => setMaxUsers(parseInt(e.target.value))} // Update maxUsers state on change
            required
            className="w-full h-2 bg-gray-600 rounded-full appearance-none focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        </div>

        <button
            type="submit"
            className="w-full bg-blue-600 hover:bg-blue-700 transition rounded px-4 py-2 font-semibold"
        >
            Create
        </button>
        </form>
    </motion.div>
    </div>
);
}