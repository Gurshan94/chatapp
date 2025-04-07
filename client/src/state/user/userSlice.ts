import { createSlice, PayloadAction } from "@reduxjs/toolkit"

interface AuthState {
    isLoggedin: boolean;
    user: {
        id: number,
        username: string,
        email: string,
    } | null;
}

const initialState: AuthState = {
    isLoggedin: false,
    user: null,
}

const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        login:(state,action:PayloadAction<{id:number; username:string; email:string; password:string}>) => {
            state.isLoggedin=true;
            state.user={
                id:action.payload.id,
                username: action.payload.username,
                email: action.payload.email,
            }
        },
        logout:(state) => {
            state.isLoggedin=false;
            state.user=null;
        },
    },
})

export const {login,logout} = authSlice.actions
export default authSlice.reducer;