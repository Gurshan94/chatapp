const RoomList = () => {
    return (
      <div className="w-72 bg-gray-850 border-r border-gray-700 p-4 overflow-y-auto">
        <h2 className="text-xl font-semibold mb-4">Rooms</h2>
        <ul className="space-y-3">
          {[1, 2, 3, 4,5,6,7,8,9,9,9,9,9,9,9,9].map((room) => (
            <li
              key={room}
              className="p-3 rounded-lg bg-gray-800 hover:bg-gray-700 transition cursor-pointer"
            >
              Room {room}
            </li>
          ))}
        </ul>
      </div>
    );
  };
  
  export default RoomList;