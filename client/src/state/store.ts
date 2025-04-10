import { configureStore } from "@reduxjs/toolkit";
import userReducer  from "./user/authSlice";
import roomReducer from './room/roomSlice'
import { userApi } from "./user/userApiSlice";
import { roomApi } from "./room/roomApiSlice";

export const store = configureStore({
    reducer:{
        [userApi.reducerPath]: userApi.reducer,
        user: userReducer,
        [roomApi.reducerPath]: roomApi.reducer,
        room: roomReducer
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(userApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch