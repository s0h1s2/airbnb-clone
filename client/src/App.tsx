import { Route, Routes, useNavigate } from "react-router-dom"
import IndexPage from "./pages/IndexPage"
import LoginPage from "./pages/LoginPage"
import Layout from "./Layout"
import RegisterPage from "./pages/RegisterPage"
import { useEffect } from "react"
import client, { OkResponseResult } from "./lib/client"
import { useBoundStore } from "./state/boundState"
import ProtectedPage from "./components/ProtectedPage"
import AccountPage from "./pages/AccountPage"
import ProfilePage from "./pages/ProfilePage"
import BookingsPage from "./pages/BookingsPage"
import PlacesPage from "./pages/PlacesPage"
type UserResponse = OkResponseResult<{
  user: {
    name: string,
    email: string
  }
}>;
function App() {
  const navigate = useNavigate()
  const setUser = useBoundStore((state) => state.setUser)
  useEffect(() => {
    const isAuth = async () => {
      try {
        const response = await client.get<UserResponse>("/users/me", { withCredentials: true })
        const { user } = response.data.data
        setUser({ name: user.name, email: user.email })
      } catch (e) {
        navigate("/login")
      }

    }
    isAuth()
  }, [])
  return (
    <Routes>
      <Route path="/" element={<Layout />} >
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route element={<ProtectedPage />}>
          <Route path="/home" element={<IndexPage />} />
          <Route path="/account" element={<AccountPage />}>
            <Route path="/account/profile" element={<ProfilePage />} />
            <Route path="/account/bookings" element={<BookingsPage />} />
            <Route path="/account/places" element={<PlacesPage />} />
          </Route>
        </Route>
      </Route>
    </Routes>
  )
}

export default App
