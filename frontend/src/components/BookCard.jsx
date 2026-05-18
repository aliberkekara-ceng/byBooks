import React, { useContext } from "react";
import { Link } from "react-router-dom";
import { BookContext } from "../context/BookContext";

function BookCard({ book, onEdit }) {
  const { deleteBook } = useContext(BookContext);

  return (
    <div className="card book-card">
      <div className="book-info">
        <h3>{book.title}</h3>
        <p className="author">{book.author}</p>
        <p className="year">{book.year}</p>
      </div>
      <div className="book-actions">
        <Link to={`/book/${book.id}`} className="btn btn-secondary btn-sm">
          Details
        </Link>
        <button
          onClick={() => onEdit(book)}
          className="btn btn-secondary btn-sm"
        >
          Edit
        </button>
        <button
          onClick={() => deleteBook(book.id)}
          className="btn btn-danger btn-sm"
        >
          Delete
        </button>
      </div>
    </div>
  );
}

export default BookCard;
