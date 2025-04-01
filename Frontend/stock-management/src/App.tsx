
import AppPrivade from "./private/AppPrivate"
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router'
import { Login } from "./public/components/Login";
import { useSelector } from "react-redux";
import { RootState } from "./store";

const PrivateRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated } = useSelector((state: RootState) => state.auth);
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
};

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route
          path="/"
          element={
            <PrivateRoute>
              <AppPrivade />
            </PrivateRoute>
          }
          />
      </Routes>
    </Router>



  )
}

export default App
