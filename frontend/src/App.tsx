import React from "react";
import "./App.css"
import { BrowserRouter as Router, Routes, Route} from "react-router-dom"
import FlipMenu from "./components/FlipMenu"
import BuyPage from "./components/BuyPage"
import PostPage from "./components/PostPage"

function App() {
  return (
    <div className = "App">
      <Router>
        <Routes>
          <Route path="/" element={<FlipMenu />} />
          <Route path="/LoginMenu" element={<FlipMenu />} />
          <Route path="/BuyPage" element={<BuyPage />} />
          <Route path="/PostPage" element={<PostPage />} />
        </Routes>
      </Router>
    </div>
  )
}
  
export default App;