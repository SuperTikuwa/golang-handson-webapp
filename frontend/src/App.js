import React from "react";
import Header from "./Header";
import Form from "./Form";
import "./styles/App.css";
import Tasks from "./Tasks";
import axios from "axios";
import { URL } from "./config";

const App = () => {
  const [tasks, setTasks] = React.useState([]);

  const fetchTasks = async () => {
    const res = await axios.get(URL.task);
    const json = await res.data;
    setTasks(json);
  };

  return (
    <div>
      <Header />
      <Form fetchTasks={fetchTasks} />
      <hr />
      <Tasks tasks={tasks} fetchTasks={fetchTasks} />
    </div>
  );
};

export default App;
