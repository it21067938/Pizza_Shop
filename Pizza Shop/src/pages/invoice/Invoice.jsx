import React, { useState, useEffect } from "react";
import "./Invoice.css";
import axios from "axios";
import { toast } from "react-toastify";
import { FaWindowMinimize, FaPlus } from "react-icons/fa";

const Invoice = ({ url }) => {
  const [list, setList] = useState([]);
  const [cart, setCart] = useState([]);
  const [search, setSearch] = useState("");

  const fetchList = async () => {
    try {
      const response = await axios.get(`${url}/api/food/list`);
      setList(response.data);
    } catch (error) {
      toast.error("Error loading data");
      console.error(error);
    }
  };

  useEffect(() => {
    fetchList();
  }, []);

  const addToCart = (item) => {
    setCart((prevCart) => [...prevCart, item]);
  };

  const removeFromCart = (index) => {
    setCart((prevCart) => prevCart.filter((_, i) => i !== index));
  };

  const calculateSubtotal = () => {
    return cart.reduce((sum, item) => sum + parseFloat(item.Price), 0);
  };

  const TAX_RATE = 0.1; // 10% tax rate
  const subtotal = calculateSubtotal();
  const tax = subtotal * TAX_RATE;
  const total = subtotal + tax;

  return (
    <div className="xx">
      <div className="invoice add flex-col">
        <input
          onChange={(e) => setSearch(e.target.value.toLowerCase())}
          className="searchBar"
          type="search"
          placeholder="Search...."
        />
        <div className="invoice-table-format title">
          <b></b>
          <b>Item ID</b>
          <b>Name</b>
          <b>Category</b>
          <b>Price</b>
          <b></b>
        </div>
        {list
          .filter(
            (e) =>
              e.Name.toLowerCase().includes(search) ||
              e.Category.toLowerCase().includes(search)
          )
          .map((item, index) => (
            <div key={index} className="invoice-table-format">
              <FaPlus className="icon" onClick={() => addToCart(item)} />
              <p>{item.ItemID}</p>
              <p>{item.Name}</p>
              <p>{item.Category}</p>
              <p>{item.Price}</p>
            </div>
          ))}
      </div>

      <div className="invoice">
        <p>Bill</p>
        {cart.map((item, index) => (
          <div key={index} className="cart-table-format">
            <FaWindowMinimize
              className="icon"
              onClick={() => removeFromCart(index)}
            />
            <p>{item.Name}</p>
            <p>{item.Category}</p>
            <p>{item.Price}</p>
          </div>
        ))}
        <div className="cart-total">
          <div className="cart-total-details">
            <p>Subtotal</p>
            <p>Rs. {subtotal.toFixed(2)}</p>
          </div>
          <hr />
          <div className="cart-total-details">
            <p>Tax (10%)</p>
            <p>Rs. {tax.toFixed(2)}</p>
          </div>
          <hr />
          <div className="cart-total-details">
            <b>Total</b>
            <b>Rs. {total.toFixed(2)}</b>
          </div>
          <button onClick={() => toast.success("Proceeding to checkout")}>
            PROCEED TO CHECKOUT
          </button>
        </div>
      </div>
    </div>
  );
};

export default Invoice;
