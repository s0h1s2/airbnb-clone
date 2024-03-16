import { NavLink, Outlet } from "react-router-dom"
import clsx from "clsx"
import { CgProfile } from "react-icons/cg"
import { CiBookmark, CiLocationOn } from "react-icons/ci"
const AccountNavLink = ({ path, children }: { children: React.ReactNode, path: string }) => {
  return (
    <NavLink className={({ isActive }) => (clsx("inline-flex gap-1 py-2 px-6 rounded-full", isActive && "bg-primary text-white", !isActive && "bg-gray-300 text-black"))} to={path} >
      {children}
    </NavLink >
  )
}
const AccountPage = () => {
  return (
    <div>
      <nav className="flex w-full mt-8 gap-2 justify-center">
        <AccountNavLink path="/account/profile">
          <CgProfile className="mt-1" />
          My Profile
        </AccountNavLink>
        <AccountNavLink path="/account/bookings">
          <CiBookmark className="mt-1" />
          My Bookings
        </AccountNavLink>
        <AccountNavLink path="/account/places">
          <CiLocationOn className="mt-1" />
          My Places
        </AccountNavLink>
      </nav>
      <Outlet />
    </div>
  )
}

export default AccountPage
