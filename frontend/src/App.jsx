import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";

function App() {
  return (
    <Router>
      <div className="app-container">
        <header className="app-header">
          <h1>Library Management</h1>
        </header>
        <main className="app-content"></main>
      </div>
    </Router>
  );
}

export default App;
