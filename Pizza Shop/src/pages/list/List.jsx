import React, { useEffect, useState } from "react";
import "./List.css";
import axios from "axios";
import { toast } from "react-toastify";

const List = ({ url }) => {
  const [list, setList] = useState([]);
  const fetchList = async () => {
    const response = await axios.get(`${url}/api/food/list`);
    // console.log(response.data);

    if (response.data.success) {
      setList(response.data.data);
    } else {
      toast.error("Error in load data");
    }
  };

  const removeFood = async (foodId) => {
    console.log(foodId);

    const response = await axios.post(`${url}/api/food/remove/`, {
      id: foodId,
    });
    if (response.data.success) {
      toast.success(response.data.message);
      await fetchList();
    } else {
      toast.error(response.data.message);
    }
  };

  useEffect(() => {
    fetchList();
  }, []);

  return (
    <div className="list add flex-col">
      <p>All Foods List</p>
      <div className="list-table-format title">
        <b>Name</b>
        <b>Category</b>
        <b>Price</b>
        <b>{list.length}</b>
      </div>
      {list.map((item, index) => {
        return (
          <div key={index} className="list-table-format">
            <p>{item.name}</p>
            <p>{item.category}</p>
            <p>{item.price}</p>
            <p onClick={() => removeFood(item._id)} className="cross">
              X
            </p>
          </div>
        );
      })}
    </div>
  );
};

export default List;
