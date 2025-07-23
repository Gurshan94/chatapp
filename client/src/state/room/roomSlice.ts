import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Message{
    id:number
    roomid: number,
    senderid: number,
    content: string,
    username: string
}

export interface outMessage{
    roomid: number | null,
    senderid: number | undefined,
    content: string,
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
}

interface roomState {
    joinedRooms: Room[],
    discoverRooms: BackendRoom[],
    selectedRoomId: number | null,
    activeTab: "joined" | "discover" | "create",
    messages: Message[]
}

const initialState: roomState = {
    joinedRooms: [],
    discoverRooms: [],
    selectedRoomId: null,
    activeTab: "joined",
    messages: []
}

const roomSlice = createSlice({
    name:'room',
    initialState,
    reducers:{
        setJoinedRooms:(state, action:PayloadAction<BackendRoom[]>) => {
            state.joinedRooms=action.payload.map(room=>({
                ...room,
                unreadCount: 0,
            }));
        },
        closeroom:(state)=>{
            state.selectedRoomId=null
        }, 
        openRoom:(state, action:PayloadAction<{roomId:number}>) => {
            const room=state.joinedRooms.find(r => r.id == action.payload.roomId)

            if (room) {
                room.unreadCount=0;
                state.selectedRoomId=room.id;
                state.messages=[]
            }
        },
        appendMessages: (state, action: PayloadAction<{ roomId: number| null, messages: Message[] }>) => {
            const room = state.joinedRooms.find(r => r.id === action.payload.roomId);
            if (room) {

              const existingIds = new Set(state.messages.map(m => m.id));
              const newMessages = action.payload.messages.filter(m => !existingIds.has(m.id));
              state.messages.push(...newMessages);
            }
        },

        deleteMessageWithID:(state,action:PayloadAction<{messageId:number}>)=>{
            state.messages=state.messages.filter(m=>m.id!=action.payload.messageId)
        },

        addMessage: (state, action: PayloadAction<Message>) => {
            const  message  = action.payload;
            const openedRoomId = state.selectedRoomId;
          
            if (message.roomid==openedRoomId) {
                state.messages = [message,...state.messages];
            } else {
                const room=state.joinedRooms.find(r => r.id == message.roomid)
                if (room){
                    room.unreadCount+=1
                }
            }
        },

        setActiveTab: (state, action: PayloadAction<"joined" | "discover" | "create">) => {
            state.activeTab = action.payload;
        },

        joinRoom: (state, action: PayloadAction<BackendRoom>) => {
            state.joinedRooms.push({
              ...action.payload,
              unreadCount: 0,
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

export const {
    setJoinedRooms,
    openRoom,
    addMessage,
    setActiveTab,
    joinRoom,
    leaveRoom,
    appendMessages,
    deleteMessageWithID,
    closeroom,
} = roomSlice.actions;
export default roomSlice.reducer;