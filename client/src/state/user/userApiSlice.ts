import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const userApi = createApi({
  reducerPath: 'userApi',
  baseQuery: fetchBaseQuery({
    baseUrl: import.meta.env.VITE_BASE_URL,
    credentials: 'include', // send cookies
  }),
  endpoints: (builder) => ({
    login: builder.mutation<any, { email: string; password: string }>({
      query: (credentials) => ({
        url: '/login',
        method: 'POST',
        body: credentials,
        headers: {
          'Content-Type': 'application/json',
        },
      }),
    }),
    signup: builder.mutation<any, { username:string; email: string; password: string }>({
        query: (credentials) => ({
          url: '/signup',
          method: 'POST',
          body: credentials,
          headers: {
            'Content-Type': 'application/json',
          },
        }),
      }),
    logout: builder.mutation<void, void>({
      query: () => ({
        url: '/logout',
        method: 'GET',
      }),
    }),
    getMe: builder.query<void, void>({
      query: () => ({
        url: "/me",
        credentials: "include", // important to send cookies
      }),
    }),
  }),
});

export const { useLoginMutation, useSignupMutation, useLogoutMutation,useGetMeQuery } = userApi;