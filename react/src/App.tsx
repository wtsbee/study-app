import { Routes, Route } from "react-router-dom";
import Header from "./components/header/Header";
import Home from "./components/pages/Home/Home";
import NewArticle from "./components/pages/NewArticle/NewArticle";

function App() {
  return (
    <div className="App">
      <Header></Header>
      <div className="flex min-h-screen">
        <div className="flex-grow bg-blue-100 invisible md:visible"></div>
        <div className="main w-full md:w-3/5">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/article" element={<NewArticle />} />
          </Routes>
        </div>
        <div className="flex-grow bg-blue-100 invisible md:visible"></div>
      </div>
    </div>
  );
}

export default App;
