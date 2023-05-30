import { useEffect } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import axios from "axios";
import Home from "@/components/pages/Home";
import Auth from "@/components/pages/Auth";
import Board from "@/components/pages/Board/Board";
import { CsrfToken } from "@/types";

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true;
    const getCsrfToken = async () => {
      const { data } = await axios<CsrfToken>({
        method: "get",
        url: `${import.meta.env.VITE_BACKEND_URL}/csrf`,
      });
      axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
    };
    getCsrfToken();
  }, []);
  return (
    <div className="App">
      <div className="flex min-h-screen">
        <div className="flex-grow bg-blue-100 invisible md:visible"></div>
        <div className="main w-full min-h-full">
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Auth />} />
              <Route path="/home" element={<Home />} />
              <Route path="/board" element={<Board />} />
            </Routes>
          </BrowserRouter>
        </div>
        <div className="flex-grow bg-blue-100 invisible md:visible"></div>
      </div>
    </div>
  );
}

export default App;
