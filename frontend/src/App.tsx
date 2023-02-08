import Home from "./components/Home";
import HomeStudent from "./components/Home-Student";
import HomeAdmin from "./components/Home-Admin";
import './App.css';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import CreateAdmin from "./components/Admin/CreateAdmin";
import DataAdmin from "./components/Admin/DataAdmin";
import UpdateAdmin from "./components/Admin/UpdateAdmin";
import StudentLogin from "./components/Login-Student";
import AdminLogin from "./components/Login-Admin";

import CreatePostponement from "./components/Postponement/CreatePostponement";
import DataPostponement from "./components/Postponement/DataPostponement";
import UpdatePostponement from "./components/Postponement/UpdatePostponement";
import SearchPostponement from "./components/Postponement/SearchPostponement";

function App() {
  return (
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/StudentLogin" element={<StudentLogin />} />
          <Route path="/AdminLogin" element={<AdminLogin />} />

          <Route path="/HomeStudent" element={<HomeStudent />} />
          <Route path="/HomeAdmin" element={<HomeAdmin />} />
          <Route path="/CreateAdmin" element={<CreateAdmin />} />         
          <Route path="/DataAdmin" element={<DataAdmin />} />
          <Route path="/DataAdmin/UpdateAdmin/:id" element={<UpdateAdmin/>} />
          <Route path="/CreatePostponement" element={<CreatePostponement />} />         
          <Route path="/DataPostponement" element={<DataPostponement />} />
          <Route path="/DataPostponement/UpdatePosponement/:id" element={<UpdatePostponement />} />
          <Route path="/DataPostponement/SearchPostponement/:id" element={<SearchPostponement />} />
        </Routes>

  );
}
export default App;
