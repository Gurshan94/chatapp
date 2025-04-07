import { useLocation, useNavigate } from 'react-router-dom';


const Navbar=()=>{

    const location = useLocation();
    const isRoot = location.pathname === '/';
    const navigate = useNavigate();

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
      </nav>
    )
}

export default Navbar