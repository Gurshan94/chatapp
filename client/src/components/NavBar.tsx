import { useLocation, useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../state/store';
import { useLogoutMutation } from '../state/user/userApiSlice';
import { useDispatch } from "react-redux";
import { AppDispatch } from "../state/store"; // adjust this if needed
import { clearCredentials } from '../state/user/authSlice';

const Navbar=()=>{
    const dispatch = useDispatch<AppDispatch>();
  
    const username=useSelector((state:RootState)=>state.user.user?.username)
    const location = useLocation();
    const isRoot = location.pathname === '/';
    const navigate = useNavigate();

    const[logout]=useLogoutMutation();

    const handlerlogout=async ()=>{
      await logout()
      dispatch(clearCredentials())
      navigate('/');
    }

    return (
        <nav className="flex justify-between items-center px-8 py-4 bg-gray-800 shadow-md">
        <h1 className="text-3xl font-bold cursor-pointer" onClick={() => navigate('/')}>
          ChatZone
        </h1>
        {isRoot ? (<div className="space-x-4">
          <button
            onClick={() => navigate('/login')}
            className="px-4 py-2 bg-blue-600 rounded hover:bg-blue-700 transition"
          >
            Login
          </button>
          <button
            onClick={() => navigate('/signup')}
            className="px-4 py-2 bg-green-600 rounded hover:bg-green-700 transition"
          >
            Signup
          </button>
        </div>):null}

        {username?
        <div className="flex gap-4">
         <div className='text-2xl'>👋 {username} </div>
          <button
            onClick={() => handlerlogout()}
            className="px-3 py-2 bg-red-600 rounded hover:bg-red-700 transition"
          >
            Logout
          </button>
         </div>
        :null}
      </nav>
    )
}

export default Navbar