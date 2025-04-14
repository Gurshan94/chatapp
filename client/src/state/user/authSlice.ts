import { createSlice, PayloadAction } from "@reduxjs/toolkit"

interface AuthState {
    isLoggedin: boolean;
    user: {
        id: number,
        username: string,
        email: string,
    } | null;
}

export interface User {
    id:number,
    username:string,
    email:string,
    token:string
}

const initialState: AuthState = {
    isLoggedin: false,
    user: null,
}

const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        setCredentials:(state,action:PayloadAction<{id:number; username:string; email:string}>) => {
            state.isLoggedin=true;
            state.user={
                id:action.payload.id,
                username: action.payload.username,
                email: action.payload.email,
            }
        },
        clearCredentials:(state) => {
            state.isLoggedin=false;
            state.user=null;
        },
    },
})

export const {setCredentials,clearCredentials} = authSlice.actions
export default authSlice.reducer;