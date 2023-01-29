import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css"
import Home from "./pages/Home"
import RegisterHome from "./pages/RegisterHome";
import RegisterEmployee from "./pages/RegisterEmployee";
import RegisterEmployer from "./pages/RegisterEmployer";
import LoginHome from "./pages/LoginHome";
import LoginEmployee from "./pages/LoginEmployee";
import EmployeeHome from "./pages/EmployeeHome";
import EmployeeProfile from "./pages/EmployeeProfile";
import EmployeeEmployers from "./pages/EmployeeEmployers";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/register" element={<RegisterHome />} />
        <Route path="/register/employee" element={<RegisterEmployee />} />
        <Route path="/register/employer" element={<RegisterEmployer />} />
        <Route path="/login" element={<LoginHome />} />
        <Route path="/login/employee" element={<LoginEmployee />} />
        <Route path="/employee/home" element={<EmployeeHome />} />
        <Route path="/employee/profile" element={<EmployeeProfile />} />
        <Route path="/employee/employers" element={<EmployeeEmployers />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
