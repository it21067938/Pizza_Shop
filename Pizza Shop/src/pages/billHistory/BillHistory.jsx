import React, { useState, useEffect } from "react";
import "./BillHistory.css";
import axios from "axios";
import { toast } from "react-toastify";
import { assets } from "../../assets/assets";

const BillHistory = ({ url }) => {
  const [orders, setOrders] = useState([]);

  const fetchAllOrders = async () => {
    try {
      const response = await axios.get(`${url}/api/bill/history`);
      setOrders(response.data); // Assuming data is an array of bills
      console.log(response.data);
    } catch (error) {
      toast.error("Failed to fetch orders");
      console.error(error);
    }
  };

  useEffect(() => {
    fetchAllOrders();
  }, [url]);

  return (
    <div className="order add">
      <h3>Order Page</h3>
      <div className="order-list">
        {orders.map((order) => (
          <div key={order.BillID} className="order-item">
            <img src={assets.parcel_icon} alt="Parcel Icon" />
            <div>
              <p className="order-item-food">
                {order.CartItems ? (
                  order.CartItems.map((item, index) => {
                    const isLast = index === order.CartItems.length - 1;
                    return (
                      <span key={index}>
                        {item.Name} - {item.Price} x {item.Quantity}
                        {!isLast && ", "}
                      </span>
                    );
                  })
                ) : (
                  <span>No items in this bill</span>
                )}
              </p>
              <p className="order-item-date">Bill ID: {order.BillID}</p>
              <p className="order-item-date">Date: {order.Date}</p>
              <p>Total: Rs. {order.TotalAmount}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BillHistory;
