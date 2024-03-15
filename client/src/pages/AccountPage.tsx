import { NavLink, Outlet } from "react-router-dom"
import clsx from "clsx"
const AccountNavLink = ({ path, children }: { children: React.ReactNode, path: string }) => {
  return (
    <NavLink className={({ isActive }) => (clsx("py-2 px-6  rounded-full", isActive && "bg-primary text-white"))} to={path} >
      {children}
    </NavLink >
  )
}
const AccountPage = () => {
  return (
    <div>
      <nav className="flex w-full mt-8 gap-2 justify-center">
        <AccountNavLink path="/account/profile">My Profile</AccountNavLink>
        <AccountNavLink path="/account/bookings">My Bookings</AccountNavLink>
        <AccountNavLink path="/account/places">My Places</AccountNavLink>
      </nav>
      <Outlet />
    </div>
  )
}

export default AccountPage
