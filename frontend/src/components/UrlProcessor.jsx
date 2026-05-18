import React, { useState, useContext } from "react";
import { BookContext } from "../context/BookContext";

function UrlProcessor() {
  const { cleanUrl, urlResult, isLoading } = useContext(BookContext);
  const [inputUrl, setInputUrl] = useState("");
  const [operation, setOperation] = useState("all");
  const [localError, setLocalError] = useState("");
  const [copied, setCopied] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLocalError("");
    setCopied(false);

    if (!inputUrl.trim()) {
      setLocalError("URL string cannot be empty.");
      return;
    }

    try {
      new URL(inputUrl.trim());
    } catch (_) {
      setLocalError(
        "Please enter a valid absolute URL (e.g. https://example.com).",
      );
      return;
    }

    try {
      await cleanUrl(inputUrl.trim(), operation);
    } catch (err) {}
  };

  const handleCopy = () => {
    if (urlResult) {
      navigator.clipboard.writeText(urlResult);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  return (
    <div className="card url-processor-card" style={{ marginTop: "2.5rem" }}>
      <div className="card-header" style={{ marginBottom: "1.5rem" }}>
        <h3
          style={{
            fontSize: "1.25rem",
            color: "var(--text-main)",
            marginBottom: "0.25rem",
          }}
        >
          URL Cleanup & Redirection Tool
        </h3>
        <p
          className="subtitle"
          style={{ fontSize: "0.875rem", color: "var(--text-muted)" }}
        >
          Test the URL sanitization and redirection service in real-time
        </p>
      </div>

      <form
        onSubmit={handleSubmit}
        className="url-form"
        style={{ display: "flex", flexDirection: "column", gap: "1.25rem" }}
      >
        <div className="form-group">
          <label htmlFor="inputUrl">Original URL</label>
          <input
            type="text"
            id="inputUrl"
            value={inputUrl}
            onChange={(e) => setInputUrl(e.target.value)}
            placeholder="e.g. https://BYFOOD.com/food-EXPeriences?query=abc/"
            className={localError ? "error-input" : ""}
          />
          {localError && <span className="error-text">{localError}</span>}
        </div>

        <div className="form-group">
          <label htmlFor="operation">Operation Type</label>
          <select
            id="operation"
            value={operation}
            onChange={(e) => setOperation(e.target.value)}
            style={{
              padding: "0.625rem",
              border: "1px solid var(--secondary-hover)",
              borderRadius: "var(--radius-md)",
              fontSize: "0.875rem",
              backgroundColor: "white",
              outline: "none",
              cursor: "pointer",
            }}
          >
            <option value="all">All (Redirection & Canonical)</option>
            <option value="canonical">
              Canonical (Remove query params & trailing slashes)
            </option>
            <option value="redirection">
              Redirection (Ensure www and lowercase)
            </option>
          </select>
        </div>

        <button type="submit" className="btn btn-primary" disabled={isLoading}>
          {isLoading ? "Processing..." : "Process URL"}
        </button>
      </form>

      {urlResult && (
        <div
          className="result-container"
          style={{
            marginTop: "1.5rem",
            paddingTop: "1.5rem",
            borderTop: "1px solid var(--secondary-hover)",
          }}
        >
          <label
            className="detail-label"
            style={{ display: "block", marginBottom: "0.5rem" }}
          >
            Processed URL
          </label>
          <div
            className="result-box"
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              backgroundColor: "var(--secondary-color)",
              padding: "0.75rem 1rem",
              borderRadius: "var(--radius-md)",
              border: "1px solid var(--secondary-hover)",
              gap: "1rem",
            }}
          >
            <span
              className="result-text"
              style={{
                fontSize: "0.875rem",
                fontFamily: "monospace",
                color: "var(--primary-hover)",
                wordBreak: "break-all",
              }}
            >
              {urlResult}
            </span>
            <button
              onClick={handleCopy}
              className="btn btn-secondary btn-sm copy-btn"
              style={{ flex: "none", width: "70px" }}
            >
              {copied ? "Copied!" : "Copy"}
            </button>
          </div>
        </div>
      )}
    </div>
  );
}

export default UrlProcessor;
