import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Users from "./components/Activity";

import UserCreate from "./components/ActivityCreate";

export default function App() {

return (

  <Router>

   <div>

   <Navbar />

   <Routes>

       <Route path="/" element={<Users />} />

       <Route path="/create" element={<UserCreate />} />

   </Routes>

   </div>

  </Router>

);

}