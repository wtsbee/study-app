import { Task } from "@/types";
import { useState } from "react";
import { Link } from "react-router-dom";

interface ChildrenProps {
  children: Task;
  listIndex: number;
  cardIndex: number;
}

const BoardCard = ({ children: task, listIndex, cardIndex }: ChildrenProps) => {
  return (
    <div className="card p-5 bg-pink-200 rounded mt-2">
      <Link to={`/board/${task.id}?list=${listIndex}&card=${cardIndex}`}>
        {task.title}
      </Link>
    </div>
  );
};

export default BoardCard;
