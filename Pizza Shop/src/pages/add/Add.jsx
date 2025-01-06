import React, { useState, useEffect } from "react";
import "./Add.css";
import axios from "axios";
import { toast } from "react-toastify";
import { useLocation, useNavigate } from "react-router-dom";

const Add = ({ url }) => {
  const location = useLocation();
  const navigate = useNavigate();
  const foodItem = location.state?.foodItem;

  const [data, setData] = useState({
    Name: "",
    Price: "",
    Category: "Pizzas",
  });

  useEffect(() => {
    if (foodItem) {
      setData({
        Name: foodItem.Name,
        Price: foodItem.Price,
        Category: foodItem.Category,
      });
    }
  }, [foodItem]);

  const onChangeHandler = (event) => {
    const { name, value } = event.target;
    setData((prevData) => ({
      ...prevData,
      [name]: name === "Price" ? parseFloat(value) || 0 : value,
    }));
  };

  const onSubmitHandler = async (event) => {
    event.preventDefault();

    try {
      if (foodItem) {
        const response = await axios.put(`${url}/api/food/update/${foodItem.ItemID}`, data);
        toast.success("Food item updated successfully");
      } else {
        const response = await axios.post(`${url}/api/food/add`, data);
        toast.success("Food item added successfully");
      }

      setData({
        Name: "",
        Price: "",
        Category: "Pizzas",
      });

      navigate("/list");
    } catch (error) {
      toast.error("Failed to save item!");
      console.error(error);
    }
  };

  return (
    <div className="add">
      <form onSubmit={onSubmitHandler} className="flex-col">
        <div className="add-product-name flex-col">
          <p>Product Name<span> *</span></p>
          <input
            onChange={onChangeHandler}
            value={data.Name}
            type="text"
            name="Name"
            placeholder="Type here"
            required
          />
        </div>

        <div className="add-category-price">
          <div className="add-category flex-col">
            <p>Product Category<span> *</span></p>
            <select
              onChange={onChangeHandler}
              value={data.Category}
              name="Category"
              required
            >
              <option value="Pizzas">Pizzas</option>
              <option value="Beverages">Beverages</option>
              <option value="Toppings">Toppings</option>
            </select>
          </div>

          <div className="add-price flex-col">
            <p>Product Price <span> *</span></p>
            <input
              onChange={onChangeHandler}
              value={data.Price}
              type="number"
              name="Price"
              placeholder="Rs.500"
              required
            />
          </div>
        </div>

        <button className="add-btn" type="submit">
          {foodItem ? "Update" : "Add"}
        </button>
      </form>
    </div>
  );
};

export default Add;
