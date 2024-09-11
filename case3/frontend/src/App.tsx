import './App.css'
import '../node_modules/bulma/css/bulma.min.css'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from './pages/home/index'
import Login from './pages/login/index'
import Product from "./pages/product";
import Cart from "./pages/cart";
import Transaction from "./pages/transaction";
import Register from "./pages/register";
import NotFound from "./pages/not_found";

function App() {
  return (
      <Router>
          <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/product" element={<Product />} />
              <Route path="/cart" element={<Cart />} />
              <Route path="/transaction" element={<Transaction />} />
              <Route path="*" element={<NotFound />} />
          </Routes>
      </Router>
  )
}

export default App
