import React, { createContext, useState, useEffect } from "react";

export const BookContext = createContext();

const API_URL = "http://localhost:8080/api/books";
const PROCESS_URL = "http://localhost:8080/api/url-process";

export const BookProvider = ({ children }) => {
  const [books, setBooks] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [urlResult, setUrlResult] = useState("");

  const fetchBooks = async () => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch(API_URL);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setBooks(data);
    } catch (err) {
      setError(
        "Connection error: Server is unreachable. (" + err.message + ")",
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const deleteBook = async (id) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch(`${API_URL}/${id}`, {
        method: "DELETE",
      });

      if (!response.ok) throw new Error("Delete operation failed.");

      setBooks(books.filter((book) => String(book.id) !== String(id)));
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const addBook = async (newBook) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch(API_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newBook),
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.error || "Add operation failed.");
      }

      const addedBook = await response.json();
      setBooks([...books, addedBook]);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const updateBook = async (id, updatedData) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedData),
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.error || "Update operation failed.");
      }

      const updatedBook = await response.json();
      setBooks(
        books.map((book) =>
          String(book.id) === String(id) ? updatedBook : book,
        ),
      );
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const cleanUrl = async (url, operation) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await fetch(PROCESS_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url, operation }),
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.error || "URL processing failed.");
      }

      const data = await response.json();
      setUrlResult(data.processed_url);
      return data.processed_url;
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  const clearError = () => setError(null);

  return (
    <BookContext.Provider
      value={{
        books,
        setBooks,
        deleteBook,
        addBook,
        updateBook,
        cleanUrl,
        urlResult,
        setUrlResult,
        isLoading,
        error,
        clearError,
      }}
    >
      {children}
    </BookContext.Provider>
  );
};
