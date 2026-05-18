import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { BookProvider } from "./context/BookContext";
import Dashboard from "./pages/Dashboard";
import BookDetail from "./pages/BookDetail";
import "./App.css";

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
              <Route path="/book/:id" element={<BookDetail />} />
            </Routes>
          </main>
        </div>
      </Router>
    </BookProvider>
  );
}

export default App;
