import React from "react";
import { BookProvider } from "./context/BookContext";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";
import Dashboard from "./pages/Dashboard";

function App() {
  return (
    <BookProvider>
      <Router>
        <div className="app-container">
          <header className="app-header">
            <h1>Library Management</h1>
          </header>
          <main className="app-content">
            <Routes>
              <Route path="/" element={<Dashboard />} />
            </Routes>
          </main>
        </div>
      </Router>
    </BookProvider>
  );
}

export default App;
