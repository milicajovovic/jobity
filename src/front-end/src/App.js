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
import LoginEmployer from "./pages/LoginEmployer";
import LoginAdmin from "./pages/LoginAdmin";
import AdminEmployees from "./pages/AdminEmployees";
import AdminEmployers from "./pages/AdminEmployers";
import AdminAds from "./pages/AdminAds";
import AdminReviews from "./pages/AdminReviews";
import EmployerHome from "./pages/EmployerHome";
import EmployerProfile from "./pages/EmployerProfile";
import CreateAd from "./pages/CreateAd";
import EmployerReviews from "./pages/EmployerReviews";
import EmployerApplications from "./pages/EmployerApplications";
import EmployerInterviews from "./pages/EmployerInterviews";

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
        <Route path="/login/employer" element={<LoginEmployer />} />
        <Route path="/login/admin" element={<LoginAdmin />} />
        <Route path="/employee/home" element={<EmployeeHome />} />
        <Route path="/employee/profile" element={<EmployeeProfile />} />
        <Route path="/employee/employers" element={<EmployeeEmployers />} />
        <Route path="/admin/employees" element={<AdminEmployees />} />
        <Route path="/admin/employers" element={<AdminEmployers />} />
        <Route path="/admin/ads" element={<AdminAds />} />
        <Route path="/admin/reviews" element={<AdminReviews />} />
        <Route path="/employer/home" element={<EmployerHome />} />
        <Route path="/employer/profile" element={<EmployerProfile />} />
        <Route path="/employer/ad" element={<CreateAd />} />
        <Route path="/employer/reviews" element={<EmployerReviews />} />
        <Route path="/employer/applications" element={<EmployerApplications />} />
        <Route path="/employer/interviews" element={<EmployerInterviews />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
