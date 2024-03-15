import Header from './components/header'
import { Outlet } from 'react-router-dom'

const Layout = () => {
  return (
    <div className="p-4">
      <Header />
      <Outlet />
    </div>
  )
}

export default Layout
