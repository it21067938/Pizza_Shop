import React, { useState, useEffect } from "react";
import "./BillHistory.css";
import axios from "axios";
import { toast } from "react-toastify";
import { assets } from "../../assets/assets";

const BillHistory = ({ url }) => {
  const [orders, setOrders] = useState([]);

  const fetchAllOrders = async () => {
    try {
      const response = await axios.get(`${url}/api/bills/history`);
      setOrders(response.data); // Assuming data is an array of bills
      console.log(response.data); // Debugging the response
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
          <div key={order.billID} className="order-item">
            <img src={assets.parcel_icon} alt="Parcel Icon" />
            <div>
              <p className="order-item-food">
                {order.items && order.items.length > 0 ? (
                  order.items.map((item, index) => {
                    const isLast = index === order.items.length - 1;
                    return (
                      <span key={index}>
                        Item ID: {item.itemID} - Price: Rs. {item.price} x {item.quantity}
                        {!isLast && ", "}
                      </span>
                    );
                  })
                ) : (
                  <span>No items in this bill</span>
                )}
              </p>
              <p className="order-item-date">Bill ID: {order.billID}</p>
              <p className="order-item-date">Date: {order.date}</p>
              <p>Total: Rs. {order.totalAmount}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BillHistory;
