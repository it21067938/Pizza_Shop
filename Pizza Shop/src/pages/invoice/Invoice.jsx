import React, { useState, useEffect } from "react";
import "./Invoice.css";
import axios from "axios";
import { toast } from "react-toastify";
import { FaWindowMinimize, FaPlus } from "react-icons/fa";

const Invoice = ({ url }) => {
  const [list, setList] = useState([]);
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
        
        <div className="cart-total">
          <div className="cart-total-details">
            <p>Subtotal</p>
            <p>Rs.</p>
          </div>
          <hr />
          <div className="cart-total-details">
            <p>Tax (10%)</p>
            <p>Rs.</p>
          </div>
          <hr />
          <div className="cart-total-details">
            <b>Total</b>
            <b>Rs. </b>
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
