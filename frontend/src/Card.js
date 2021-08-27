import axios from "axios";
import React from "react";
import { URL } from "./config";

const Card = (props) => {
  const handleClick = async (e) => {
    e.preventDefault();
    const data = {
      id: props.taskID,
    };
    await axios.put(URL.task, data);
    await props.fetchTasks();
  };

  return (
    <div className="card">
      <h5 className="card-title">{props.title}</h5>
      <h6>{props.date}</h6>
      <div className="card-body">
        <p className="card-text">{props.description}</p>
        <input
          type="button"
          className="done-button"
          value="Done"
          onClick={handleClick}
        />
      </div>
    </div>
  );
};
export default Card;
