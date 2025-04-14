import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { QueryClient,QueryClientProvider } from '@tanstack/react-query';
import App from './App.tsx'
import './index.css'
import { Provider } from 'react-redux';
import {store} from './state/store.ts'
import { 
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider
} from 'react-router-dom';
import Login from './screens/login.tsx';
import Signup from './screens/signup.tsx';
import Chat from './screens/chat.tsx';
import { SocketProvider } from './utils/socketProvider.tsx';

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path='/' element={<App />}>
      <Route path='login' element={<Login />} />
      <Route path='signup' element={<Signup />} />
      <Route path='chat' element={<Chat />} />
    </Route>
  )
);


const queryClient=new QueryClient();


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <SocketProvider>
    <Provider store={store}>
      <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      </QueryClientProvider>
    </Provider>
    </SocketProvider>
  </StrictMode>
);
