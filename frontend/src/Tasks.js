import React from "react";
import Card from "./Card";

const Tasks = (props) => {
  React.useEffect(props.fetchTasks, []);

  return (
    <div>
      {props.tasks === undefined
        ? null
        : props.tasks.map((task, index) => (
            <Card
              key={index}
              taskID={task.id}
              title={task.title}
              description={task.description}
              date={task.date}
              fetchTasks={props.fetchTasks}
            />
          ))}
    </div>
  );
};

export default Tasks;
