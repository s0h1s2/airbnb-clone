import { ToastContainer } from 'react-toastify'
import Header from './components/Header'
import { Outlet } from 'react-router-dom'
import 'react-toastify/dist/ReactToastify.css';
const Layout = () => {
  return (
    <div className="p-4 flex flex-col min-h-screen">
      <Header />
      <Outlet />
      <ToastContainer />
    </div>
  )
}

export default Layout
