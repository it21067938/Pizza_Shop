import React from 'react'
import "./SlideBar.css"
import { assets } from '../../assets/assets'
import { NavLink } from 'react-router-dom'

function SlidBar() {
  return (
    <div className='slidebar'>
      <div className="slidebar-options">
        <NavLink to="/add" className="slidebar-option">
          <img src={assets.add_icon} alt="" />
          <p>Add Items</p>
        </NavLink>
        <NavLink to="/list" className="slidebar-option">
          <img src={assets.order_icon} alt="" />
          <p>List Items</p>
        </NavLink>
        <NavLink to="/history" className="slidebar-option">
          <img src={assets.order_icon} alt="" />
          <p>Bill History</p>
        </NavLink>
        <NavLink to="/invoice" className="slidebar-option">
          <img src={assets.order_icon} alt="" />
          <p>Invoice</p>
        </NavLink>
      </div>
    </div>
  )
}

export default SlidBar