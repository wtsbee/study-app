import { useMutateTask } from "@/hooks/useMutateTask";
import { Task, TaskList } from "@/types";
import { Link } from "react-router-dom";

interface ChildrenProps {
  children: Task;
  socketRef: React.MutableRefObject<WebSocket | undefined>;
  taskListArray: TaskList[];
  listIndex: number;
  cardIndex: number;
}

const BoardCard = ({
  children: task,
  socketRef,
  taskListArray,
  listIndex,
  cardIndex,
}: ChildrenProps) => {
  const { deleteTaskMutation } = useMutateTask();

  const deleteTask = () => {
    const newArray = [...taskListArray];
    const updatedArray = newArray.map((item, index) => {
      if (index === listIndex) {
        const updatedTasks = item.tasks.filter(
          (_, taskIndex) => taskIndex !== cardIndex
        );
        return { ...item, tasks: updatedTasks };
      }
      return item;
    });
    deleteTaskMutation.mutate(task.id as number);
    socketRef.current?.send(JSON.stringify(updatedArray));
  };

  return (
    <div className="card p-5 bg-pink-200 rounded mt-2 flex justify-between">
      <Link to={`/board/${task.id}?`}>
        <div className="break-all">{task.title}</div>
      </Link>
      <div className="hover:cursor-pointer">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-6 h-6 rounded hover:bg-pink-400"
          onClick={deleteTask}
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </div>
    </div>
  );
};

export default BoardCard;
