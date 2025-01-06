import React, { useEffect, useState } from "react";
import "./List.css";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { toast } from "react-toastify";
import { MdEdit } from "react-icons/md";

const List = ({ url }) => {
  const navigate = useNavigate();
  const [list, setList] = useState([]);

  const fetchList = async () => {
    try {
      const response = await axios.get(`${url}/api/food/list`);

      setList(response.data);
    } catch (error) {
      toast.error("Error in load data");
      console.error(error);
    }
  };

  const removeFood = async (ItemID) => {
    try {
      const response = await axios.delete(
        `http://localhost:8080/api/food/remove/${ItemID}`,
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      if (response.status === 204) {
        toast.success("Food removed successfully");
        fetchList();
      } else {
        toast.error("Error removing food");
      }
    } catch (error) {
      toast.error("Error removing food");
      console.error(error);
    }
  };

  const handleEditClick = (item) => {
    navigate("/add", { state: { foodItem: item } });
  };

  const [search, setSerch] = useState("");
  function searchItem(event) {
    setSerch(event.target.value);
  }

  useEffect(() => {
    fetchList();
  }, []);

  return (
    <div className="list add flex-col">
      <input
        onChange={searchItem}
        className="searchBar"
        type="search"
        placeholder="Search...."
        name="searchQuery"
      />
      <p>All Foods List</p>
      <div className="list-table-format title">
        <b>Item ID</b>
        <b>Name</b>
        <b>Category</b>
        <b>Price</b>
        <b></b>
        <b>{list.length}</b>
      </div>
      {list
        .filter(
          (e) =>
            e.Name.toLowerCase().includes(search) || e.Category.toLowerCase().includes(search)
        )
        .map((item, index) => {
          return (
            <div key={index} className="list-table-format">
              <p>{item.ItemID}</p>
              <p>{item.Name}</p>
              <p>{item.Category}</p>
              <p>{item.Price}</p>
              <p onClick={() => handleEditClick(item)} className="edit">
                <MdEdit />
              </p>
              <p onClick={() => removeFood(item.ItemID)} className="cross">
                X
              </p>
            </div>
          );
        })}
    </div>
  );
};

export default List;
