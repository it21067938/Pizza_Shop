import { Route, Routes } from "react-router-dom";
import Navbar from "./components/navbar/Navbar";
import SlideBar from "./components/slideBar/SlideBar";
import Add from "./pages/add/Add";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import List from "./pages/list/List"
import Invoice from './pages/invoice/Invoice';
import BillHistory from './pages/billHistory/BillHistory';

function App() {

  const url = "http://localhost:8080";

  return (
    <div>
      <ToastContainer />
      <Navbar />
      <hr />
      <div className="app-content">
        <SlideBar />
        <Routes>
          <Route path="/add" element={<Add url={url}/>} />
          <Route path="/list" element={<List url={url}/>} />
          <Route path="/invoice" element={<Invoice url={url}/>} />
          <Route path="/history" element={<BillHistory url={url}/>} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
