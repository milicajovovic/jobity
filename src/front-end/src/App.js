import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css"
import Home from "./pages/Home"
import RegisterHome from "./pages/RegisterHome";
import RegisterEmployee from "./pages/RegisterEmployee";
import RegisterEmployer from "./pages/RegisterEmployer";
import LoginHome from "./pages/LoginHome";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/register" element={<RegisterHome />} />
        <Route path="/register/employee" element={<RegisterEmployee />} />
        <Route path="/register/employer" element={<RegisterEmployer />} />
        <Route path="/login" element={<LoginHome />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
