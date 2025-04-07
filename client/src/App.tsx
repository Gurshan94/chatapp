import { Outlet, useLocation } from 'react-router-dom';
import { motion } from 'framer-motion';
import Navbar from './components/NavBar';

const App = () => {
  const location = useLocation();
  const isRoot = location.pathname === '/';

  return (
    <div className="h-screen flex flex-col bg-gray-900 text-white">
      <Navbar /> 

      {isRoot ? (
        <div className="flex-1 flex flex-col items-center justify-center px-4">
          <motion.h1
            className="text-5xl md:text-6xl font-extrabold mb-6 text-center"
            initial={{ opacity: 0, y: -30 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 1, ease: 'easeOut' }}
          >
            Welcome to ChatZone 👋
          </motion.h1>

          <motion.p
            className="text-lg md:text-xl mb-10 text-center max-w-xl"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.5, duration: 1 }}
          >
            Connect and chat in real-time. Join rooms, talk with friends, and experience seamless messaging.
          </motion.p>
        </div>
      ) : (
        // For the chat page, use Outlet, which will take the rest of the space
        <div className="flex-1 overflow-hidden">
          <Outlet />
        </div>
      )}
    </div>
  );
};

export default App;