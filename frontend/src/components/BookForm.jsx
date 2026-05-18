import React, { useState } from "react";

function BookForm({ onSubmit, onCancel, initialData = null }) {
  const [formData, setFormData] = useState({
    title: initialData?.title || "",
    author: initialData?.author || "",
    year: initialData?.year || "",
  });

  const [errors, setErrors] = useState({});

  const validate = () => {
    const newErrors = {};
    if (!formData.title.trim()) newErrors.title = "Title is required";
    if (!formData.author.trim()) newErrors.author = "Author is required";
    if (!formData.year) {
      newErrors.year = "Publication year is required";
    } else if (
      isNaN(formData.year) ||
      formData.year < 1000 ||
      formData.year > new Date().getFullYear()
    ) {
      newErrors.year = "Please enter a valid year";
    }
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
    if (errors[name]) {
      setErrors((prev) => ({ ...prev, [name]: "" }));
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validate()) {
      onSubmit({ ...formData, year: parseInt(formData.year, 10) });
    }
  };

  return (
    <form onSubmit={handleSubmit} className="book-form">
      <div className="form-group">
        <label htmlFor="title">Book Title</label>
        <input
          type="text"
          id="title"
          name="title"
          value={formData.title}
          onChange={handleChange}
          className={errors.title ? "error-input" : ""}
          placeholder="e.g. Crime and Punishment"
        />
        {errors.title && <span className="error-text">{errors.title}</span>}
      </div>

      <div className="form-group">
        <label htmlFor="author">Author</label>
        <input
          type="text"
          id="author"
          name="author"
          value={formData.author}
          onChange={handleChange}
          className={errors.author ? "error-input" : ""}
          placeholder="e.g. Fyodor Dostoevsky"
        />
        {errors.author && <span className="error-text">{errors.author}</span>}
      </div>

      <div className="form-group">
        <label htmlFor="year">Publication Year</label>
        <input
          type="number"
          id="year"
          name="year"
          value={formData.year}
          onChange={handleChange}
          className={errors.year ? "error-input" : ""}
          placeholder="e.g. 1866"
        />
        {errors.year && <span className="error-text">{errors.year}</span>}
      </div>

      <div className="form-actions">
        <button type="button" onClick={onCancel} className="btn btn-secondary">
          Cancel
        </button>
        <button type="submit" className="btn btn-primary">
          {initialData ? "Update" : "Add Book"}
        </button>
      </div>
    </form>
  );
}

export default BookForm;
