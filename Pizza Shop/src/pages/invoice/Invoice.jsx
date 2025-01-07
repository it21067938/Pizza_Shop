import React, { useState, useEffect } from "react";
import "./Invoice.css";
import axios from "axios";
import { toast } from "react-toastify";
import { printInvoice } from "../../components/printInvoice/PrintInvoice";
import { FaWindowMinimize, FaPlus } from "react-icons/fa";
import { MdOutlineLocalPrintshop } from "react-icons/md";

const Invoice = ({ url }) => {
  const [list, setList] = useState([]);
  const [search, setSearch] = useState("");
  const [cart, setCart] = useState([]);
  const [subtotal, setSubtotal] = useState(0);

  // Fetch the item list from the API
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
    const savedCart = JSON.parse(localStorage.getItem("cart")) || [];
    setCart(savedCart);
    fetchList();
  }, []);

  useEffect(() => {
    localStorage.setItem("cart", JSON.stringify(cart));

    const total = cart.reduce(
      (sum, item) => sum + item.Price * item.Quantity,
      0
    );
    setSubtotal(total);
  }, [cart]);

  const addToCart = (item) => {
    const existing = cart.find((cartItem) => cartItem.ItemID === item.ItemID);
    if (existing) {
      setCart(
        cart.map((cartItem) =>
          cartItem.ItemID === item.ItemID
            ? { ...cartItem, Quantity: cartItem.Quantity + 1 }
            : cartItem
        )
      );
    } else {
      setCart([...cart, { ...item, Quantity: 1 }]);
    }
  };

  const removeFromCart = (itemID) => {
    const updatedCart = cart.filter((item) => item.ItemID !== itemID);
    setCart(updatedCart);
  };

  const handleCheckout = async () => {
    const tax = subtotal * 0.1;
    const totalAmount = subtotal + tax;

    const data = {
      totalAmount,
      items: cart.map((item) => ({
        ItemID: item.ItemID,
        Quantity: item.Quantity,
        Price: item.Price,
      })),
    };

    try {
      await axios.post(`${url}/api/bills`, data);
      toast.success("Bill created successfully");
      setCart([]);
      setSubtotal(0);
      localStorage.removeItem("cart");
    } catch (error) {
      toast.error("Failed to create bill");
      console.error(error);
    }
  };

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
        <p>Invoice</p>
        <div className="cart">
          {cart.map((item, index) => (
            <div key={index} className="bill-table-format">
              <p>{item.Name}</p>
              <p>Qty: {item.Quantity}</p>
              <p>Price: Rs. {item.Price * item.Quantity}</p>
              <FaWindowMinimize
                className="icon"
                onClick={() => removeFromCart(item.ItemID)}
              />
            </div>
          ))}
        </div>

        <div className="cart-total">
          <div className="cart-total-details">
            <p>Subtotal</p>
            <p>Rs. {subtotal.toFixed(2)}</p>
          </div>
          <hr />
          <div className="cart-total-details">
            <p>Tax (10%)</p>
            <p>Rs. {(subtotal * 0.1).toFixed(2)}</p>
          </div>
          <hr />
          <div className="cart-total-details">
            <b>Total</b>
            <b>Rs. {(subtotal + subtotal * 0.1).toFixed(2)}</b>
          </div>
          <button onClick={handleCheckout}>PROCEED TO CHECKOUT</button>
        </div>

        <button className="printBTN"  onClick={() => printInvoice(cart, subtotal)}>
         <MdOutlineLocalPrintshop/> Print Invoice
        </button>
      </div>
    </div>
  );
};

export default Invoice;
