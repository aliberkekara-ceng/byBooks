import React, { useContext } from "react";
import { useParams, Link, useNavigate } from "react-router-dom";
import { BookContext } from "../context/BookContext";

function BookDetail() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { books, deleteBook, isLoading, error, clearError } =
    useContext(BookContext);

  const book = books.find((b) => String(b.id) === id);

  const handleDelete = async () => {
    await deleteBook(id);
    navigate("/");
  };

  if (isLoading) {
    return (
      <div className="loading-container">
        <div className="spinner"></div>
        <p>Loading book details...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="book-detail">
        <div className="alert alert-error">
          <p>{error}</p>
          <button onClick={clearError} className="close-btn">
            &times;
          </button>
        </div>
        <Link to="/" className="back-link">
          &larr; Back to Dashboard
        </Link>
      </div>
    );
  }

  if (!book) {
    return (
      <div className="book-detail not-found">
        <div className="empty-state">
          <h2>Book Not Found</h2>
          <p>The book you are looking for does not exist or may have been deleted.</p>
          <Link
            to="/"
            className="btn btn-primary"
            style={{ marginTop: "1rem" }}
          >
            Back to Dashboard
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className="book-detail">
      <Link to="/" className="back-link">
        &larr; Back to Dashboard
      </Link>

      <div className="card detail-card">
        <div className="detail-header">
          <h2>{book.title}</h2>
          <span className="badge">{book.year}</span>
        </div>

        <div className="detail-content">
          <div className="detail-row">
            <span className="detail-label">Author:</span>
            <span className="detail-value">{book.author}</span>
          </div>
          <div className="detail-row">
            <span className="detail-label">Publication Year:</span>
            <span className="detail-value">{book.year}</span>
          </div>
          <div className="detail-row">
            <span className="detail-label">Book ID:</span>
            <span className="detail-value text-muted">{book.id}</span>
          </div>
        </div>

        <div className="detail-actions">
          <button onClick={handleDelete} className="btn btn-danger">
            Delete Book
          </button>
        </div>
      </div>
    </div>
  );
}

export default BookDetail;
