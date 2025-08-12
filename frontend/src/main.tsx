import App from "./app1";
import "./app.css";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

const node = document.getElementById("root");
const root = createRoot(node);
root.render(
  <StrictMode>
    <App />
  </StrictMode>,
);
