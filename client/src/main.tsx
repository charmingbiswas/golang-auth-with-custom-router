import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import LoginPage from "./pages/LoginPage";

document.title = "Golang Custom Login";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <LoginPage />
  </React.StrictMode>
);
