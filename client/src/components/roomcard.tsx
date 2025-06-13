import React from "react";
import { BackendRoom, Room } from "../state/room/roomSlice";

interface RoomCardProps {
    room: BackendRoom | Room;
  }

const RoomCard: React.FC<RoomCardProps> = ({ room }) => {
  return (
    <div className="p-4 rounded bg-gray-800 text-gray-200 shadow-sm hover:bg-gray-700 transition-colors duration-200 flex justify-between items-center">
      <div>
        <h3 className="text-lg font-semibold mb-1">{room.roomname}</h3>
        <div className="text-sm flex items-center gap-2">
          🧑 <strong>{room.currentusers}</strong> / {room.maxusers} users
        </div>
      </div>

      {'unreadCount' in room && room.unreadCount > 0 && (
        <div className="bg-red-600 text-white rounded-full w-6 h-6 flex items-center justify-center text-xs font-bold">
          {room.unreadCount}
        </div>
      )}

      {!('unreadCount' in room) && (
        <button className = " bg-blue-600 hover:bg-blue-700 transition rounded px-4 py-2 font-semibold">
          Join
        </button>
      )}
    </div>
  );
};

export default RoomCard;