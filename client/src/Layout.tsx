import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify'
import Header from './components/Header'
import { Outlet } from 'react-router-dom';

const Layout = () => {

  return (
    <div className="p-4 flex flex-col min-h-screen">
      <ToastContainer />
      <Header />
      <Outlet />
    </div>
  )
}

export default Layout
