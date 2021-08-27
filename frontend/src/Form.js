import React from "react";
import axios from "axios";
import { URL } from "./config";

const Form = (props) => {
  const title = React.useRef(null);
  const date = React.useRef(null);
  const description = React.useRef(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (title.current.value === "" || date.current.value === "") {
      alert("Please fill out all fields");
      return;
    }
    let data = {
      title: title.current.value,
      date: date.current.value,
      description: description.current.value,
    };
    await axios.post(URL.task, data);
    title.current.value = "";
    date.current.value = "";
    description.current.value = "";

    await props.fetchTasks();
  };

  return (
    <div className="task-register-form">
      <h2>タスク登録フォーム</h2>
      <div className="task-register-form-input">
        <input type="text" placeholder="Title" ref={title} />

        <input
          type="text"
          placeholder="Description (optional)"
          ref={description}
        />
        <input type="date" ref={date} />
        <input
          type="button"
          onClick={handleSubmit}
          value="Submit"
          className="submit-button"
        />
      </div>
    </div>
  );
};

export default Form;
