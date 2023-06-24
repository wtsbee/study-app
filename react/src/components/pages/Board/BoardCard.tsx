import { Task } from "@/types";
import { Link } from "react-router-dom";

interface ChildrenProps {
  children: Task;
}

const BoardCard = ({ children: task }: ChildrenProps) => {
  return (
    <div className="card p-5 bg-pink-200 rounded mt-2">
      <Link to={`/board/${task.id}`} state={task}>
        {task.title}
      </Link>
    </div>
  );
};

export default BoardCard;
