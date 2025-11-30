import App from "./app";
import "./app.css";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

const node = document.getElementById("root");
if (!node) throw new Error("Root element not found");

const root = createRoot(node);
root.render(
  <StrictMode>
    <App />
  </StrictMode>,
);
