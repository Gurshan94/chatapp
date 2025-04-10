import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Message{
    roomid: number,
    senderid: number,
    content: string,
    username: string
}

export interface BackendRoom {
    id: number;
    roomname: string;
    maxusers: number;
    adminid: number;
    currentusers: number;
  }

export interface Room extends BackendRoom {
    unreadCount: number;
    isopen: boolean;
    messages: Message[];
  }

interface roomState {
    joinedRooms: Room[],
    selectedRoomId: number | null,
    activeTab: "joined" | "discover";
}

const initialState: roomState = {
    joinedRooms: [],
    selectedRoomId: null,
    activeTab: "joined",
}

const roomSlice = createSlice({
    name:'room',
    initialState,
    reducers:{
        setJoinedRooms:(state, action:PayloadAction<BackendRoom[]>) => {
            state.joinedRooms=action.payload.map(room=>({
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
        joinRoom: (state, action: PayloadAction<BackendRoom>) => {
          
            state.joinedRooms.push({
              ...action.payload,
              unreadCount: 0,
              isopen: false,
              messages: [],
            });
        },
        leaveRoom: (state, action: PayloadAction<number>) => {
            const roomId = action.payload;
          
            state.joinedRooms = state.joinedRooms.filter(
              r => r.id !== roomId
            );
            
            if (state.selectedRoomId === roomId) {
              state.selectedRoomId = null;
            }
        },
    }
})

export const {setJoinedRooms,openRoom,closeRoom,addMessage,setActiveTab,joinRoom,leaveRoom} = roomSlice.actions;
export default roomSlice.reducer;