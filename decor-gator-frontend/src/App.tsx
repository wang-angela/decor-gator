import React from "react";
import "./App.css";
import { BrowserRouter as Router, Routes, Route }
    from "react-router-dom";
import Login from "./components/Login";
import Dashboard from "./components/Dashboard";
  
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/redirect" element={<Dashboard />}/> 
        <Route path="/" element={<Login/>}/>
      </Routes>
    </Router>
  );
}
  
export default App;