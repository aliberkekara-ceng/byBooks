import React, { useContext, useState } from "react";
import { BookContext } from "../context/BookContext";
import Modal from "../components/Modal";
import BookForm from "../components/BookForm";
import BookCard from "../components/BookCard";

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
        <h2>Kitap Paneli</h2>
        <button
          className="btn btn-primary"
          onClick={() => setIsAddModalOpen(true)}
          disabled={isLoading}
        >
          Kitap Ekle
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
          <p>Veriler yükleniyor...</p>
        </div>
      ) : (
        <div className="book-grid">
          {books.map((book) => (
            <BookCard key={book.id} book={book} onEdit={openEditModal} />
          ))}
          {books.length === 0 && (
            <div className="empty-state">
              <p>Henüz kitap yok. Hemen ekleyin!</p>
            </div>
          )}
        </div>
      )}

      <Modal
        isOpen={isAddModalOpen}
        onClose={() => setIsAddModalOpen(false)}
        title="Kitap Ekle"
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
        title="Kitabı Düzenle"
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
    </div>
  );
}

export default Dashboard;
