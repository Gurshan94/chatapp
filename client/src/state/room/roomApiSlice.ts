import { createApi,fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const roomApi = createApi({
    reducerPath: 'roomApi',
    baseQuery: fetchBaseQuery({
        baseUrl: import.meta.env.VITE_BASE_URL ,
        credentials: 'include',
    }),
    endpoints:(builder) =>({
        getrooms: builder.mutation<any, {limit:number, offset:number}>({
            query:(credentials) =>({
                url: `/getrooms?limit=${credentials.limit}&offset=${credentials.offset}`,
                method:'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            }),
        }),
        getroomsjoinedbyuser: builder.mutation<any, {userid:number | undefined}>({
            query:(credentials) =>({
                url: `/getroomsjoinedbyuser/${credentials.userid}`,
                method:'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            }),
        }),
        fetchmessages: builder.mutation<any, {roomid:number| null, limit:number, cursor:string}>({
            query:(credentials) =>({
                url: `/fetchmessage/${credentials.roomid}?limit=${credentials.limit}&cursor=${credentials.cursor ?? ""}`,
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            }),
        }),
        deletemessage: builder.mutation<any, {messageId:number}>({
            query:(credentials) =>({
                url: `/deletemessage/${credentials.messageId}`,
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
            }),
        }),
        createroom: builder.mutation<any, {roomname:string, maxusers:number, adminid:number}>({
            query:(credentials)=>({
                url:'/createroom',
                method:'POST',
                body:credentials,
                headers:{
                    'Content-Type':'application/json',
                },
            })
        }),
        joinroom: builder.mutation<any, {roomid:number, userid:number}>({
            query:(credentials)=>({
                url:'/addusertoroom',
                method:'POST',
                body:credentials,
                headers:{
                    'Content-Type':'application/json',
                },
            })
        })
    })
})

export const {
    useGetroomsjoinedbyuserMutation,
    useGetroomsMutation,
    useFetchmessagesMutation,
    useDeletemessageMutation,
    useCreateroomMutation,
    useJoinroomMutation
}=roomApi