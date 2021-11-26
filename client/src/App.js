import logo from "./logo.svg";
import "./App.css";
import Content from "./components/content";
import { Routes, Route } from "react-router-dom";
import SignIn from "./components/sign-in";
import * as React from "react";

function home() {
  return (
    <header className="App-header">
      <img src={logo} className="App-logo" alt="logo" />
      <Content />
    </header>
  );
}
function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<SignIn />} />
        <Route path="/home"element={home()}/>
      </Routes>
    </div>
  );
}

export default App;
