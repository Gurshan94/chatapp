import { createApi,fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const roomApi = createApi({
    reducerPath: 'roomApi',
    baseQuery: fetchBaseQuery({
        baseUrl: 'http://localhost:8080',
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
    })
})

export const {useGetroomsjoinedbyuserMutation, useGetroomsMutation}=roomApi