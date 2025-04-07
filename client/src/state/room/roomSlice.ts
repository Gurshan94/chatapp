import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Message{
    roomid: number,
    senderid: number,
    content: string,
    username: string
}

interface Room {
    id: number,
    roomname: string,
    maxusers: number,
    adminid: number,
    currentUsers: number,
    unreadCount: number,
    messages: Message[]
    isopen: boolean
}

interface roomState {
    joinedRooms: Room[],
    discoverableRooms: Room[],
    selectedRoomId: number | null,
    activeTab: "joined" | "discover";
}

const initialState: roomState = {
    joinedRooms: [],
    discoverableRooms: [],
    selectedRoomId: null,
    activeTab: "joined",
}

const roomSlice = createSlice({
    name:'room',
    initialState,
    reducers:{
        setJoinedRooms:(state, action:PayloadAction<Room[]>) => {
            state.joinedRooms=action.payload.map(room=>({
                ...room,
                unreadCount: 0,
                isopen: false,
                messages:[],
            }));
        },
        setDiscoverableRooms:(state, action:PayloadAction<Room[]>) => {
            state.discoverableRooms=action.payload.map(room=>({
                ...room,
                unreadCount: 0,
                isopen: false,
                messages:[],
            }));
        },
        openRoom:(state, action:PayloadAction<{roomId:number,messages:Message[]}>) => {
            const room=state.joinedRooms.find(r => r.id == action.payload.roomId)
            if (room) {
                room.isopen=true;
                room.messages=action.payload.messages;
                room.unreadCount=0;
                state.selectedRoomId=room.id;
            }
        },
        closeRoom:(state,action:PayloadAction<{roomId:number}>)=>{
            const room=state.joinedRooms.find(r => r.id == action.payload.roomId)
            if (room) {
                room.isopen=false;
                room.messages=[];
                if (state.selectedRoomId==room.id){
                    state.selectedRoomId=null;
                }
            }
        },
        addMessage: (state, action: PayloadAction<{ roomId: number; message: Message }>) => {
            const room = state.joinedRooms.find(r => r.id === action.payload.roomId);
            if (room) {
              if (room.isopen) {
                room.messages.push(action.payload.message);
              } else {
                room.unreadCount += 1;
              }
            }
        },
        setActiveTab: (state, action: PayloadAction<"joined" | "discover">) => {
            state.activeTab = action.payload;
        },
        joinRoom: (state, action: PayloadAction<Room>) => {
            state.discoverableRooms = state.discoverableRooms.filter(
              r => r.id !== action.payload.id
            );
          
            state.joinedRooms.push({
              ...action.payload,
              unreadCount: 0,
              isopen: false,
              messages: [],
            });
        },
        leaveRoom: (state, action: PayloadAction<number>) => {
            const roomId = action.payload;
          
            const index = state.joinedRooms.findIndex(room => room.id === roomId);
            if (index !== -1) {
              const [leftRoom] = state.joinedRooms.splice(index, 1);
          
              state.discoverableRooms.push({
                id: leftRoom.id,
                roomname: leftRoom.roomname,
                maxusers: leftRoom.maxusers,
                adminid: leftRoom.adminid,
                currentUsers: leftRoom.currentUsers,
                unreadCount: 0,
                isopen: false,
                messages: [],
              });
            }
          
            if (state.selectedRoomId === roomId) {
              state.selectedRoomId = null;
            }
        },
    }
})

export const {setJoinedRooms,setDiscoverableRooms,openRoom,closeRoom,addMessage,setActiveTab,joinRoom,leaveRoom} = roomSlice.actions;
export default roomSlice.reducer;