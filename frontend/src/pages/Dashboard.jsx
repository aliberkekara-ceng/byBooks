import React, { useContext, useState } from "react";
import { BookContext } from "../context/BookContext";
import Modal from "../components/Modal";
import BookForm from "../components/BookForm";
import BookCard from "../components/BookCard";
import UrlProcessor from "../components/UrlProcessor";

function Dashboard() {
  const { books, addBook, updateBook, isLoading, error, clearError } =
    useContext(BookContext);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [editingBook, setEditingBook] = useState(null);

  const handleAddBook = async (newBook) => {
    await addBook(newBook);
    setIsAddModalOpen(false);
  };

  const openEditModal = (book) => {
    setEditingBook(book);
    setIsEditModalOpen(true);
  };

  const handleEditBook = async (updatedData) => {
    await updateBook(editingBook.id, updatedData);
    setIsEditModalOpen(false);
    setEditingBook(null);
  };

  return (
    <div className="dashboard">
      <div className="dashboard-header">
        <h2>Books Dashboard</h2>
        <button
          className="btn btn-primary"
          onClick={() => setIsAddModalOpen(true)}
          disabled={isLoading}
        >
          Add Book
        </button>
      </div>

      {error && (
        <div className="alert alert-error">
          <p>{error}</p>
          <button onClick={clearError} className="close-btn">
            &times;
          </button>
        </div>
      )}

      {isLoading ? (
        <div className="loading-container">
          <div className="spinner"></div>
          <p>Loading data...</p>
        </div>
      ) : (
        <div className="book-grid">
          {books.map((book) => (
            <BookCard key={book.id} book={book} onEdit={openEditModal} />
          ))}
          {books.length === 0 && (
            <div className="empty-state">
              <p>No books found. Add one now!</p>
            </div>
          )}
        </div>
      )}

      <Modal
        isOpen={isAddModalOpen}
        onClose={() => setIsAddModalOpen(false)}
        title="Add Book"
      >
        {isAddModalOpen && (
          <BookForm
            onSubmit={handleAddBook}
            onCancel={() => setIsAddModalOpen(false)}
          />
        )}
      </Modal>

      <Modal
        isOpen={isEditModalOpen}
        onClose={() => {
          setIsEditModalOpen(false);
          setEditingBook(null);
        }}
        title="Edit Book"
      >
        {isEditModalOpen && editingBook && (
          <BookForm
            initialData={editingBook}
            onSubmit={handleEditBook}
            onCancel={() => {
              setIsEditModalOpen(false);
              setEditingBook(null);
            }}
          />
        )}
      </Modal>

      <UrlProcessor />
    </div>
  );
}

export default Dashboard;
