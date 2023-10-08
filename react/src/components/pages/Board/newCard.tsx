import { useState } from "react";
import { Task, TaskList } from "@/types";
import { useMutateTask } from "@/hooks/useMutateTask";
import { useQueryTasks } from "@/hooks/useQueryTasks";

interface Props {
  taskList: TaskList;
}

const newCard = ({ taskList }: Props) => {
  const { refetch } = useQueryTasks();
  const { createTaskMutation } = useMutateTask();
  const [isEdit, setIsEdit] = useState(false);
  const [input, setInput] = useState("");

  const inputText = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInput(e.target.value);
  };

  const openCard = () => {
    setIsEdit(true);
  };

  const closeCard = () => {
    setIsEdit(false);
  };

  const addCard = async () => {
    if (input !== "") {
      const rank: number =
        taskList.tasks === null ? 1 : taskList.tasks.length + 1;

      await createTaskMutation.mutateAsync({
        id: null,
        title: input,
        task_list_id: taskList.id,
        rank,
      });
      setInput("");
      await refetch();
      setIsEdit(false);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.nativeEvent.isComposing || e.key !== "Enter") return;
    addCard();
  };

  return (
    <>
      {isEdit ? (
        <>
          <div className="mt-2 p-5 bg-pink-300 rounded">
            <textarea
              onChange={inputText}
              onKeyDown={handleKeyDown}
              value={input}
              placeholder="タスクを入力してください"
              className=" w-full placeholder-black bg-pink-300 outline-none"
            />
          </div>
          <div className="flex items-center mt-2 ">
            <button
              onClick={addCard}
              className="py-2 px-4 rounded text-white bg-blue-500 font-bold hover:cursor-pointer"
            >
              カードを追加
            </button>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6 hover:cursor-pointer"
              onClick={closeCard}
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </div>
        </>
      ) : (
        <div
          className="flex mt-2 py-2 rounded hover:bg-pink-300 hover:cursor-pointer"
          onClick={openCard}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="w-6 h-6 mx-2"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M12 4.5v15m7.5-7.5h-15"
            />
          </svg>
          <span>カードを追加</span>
        </div>
      )}
    </>
  );
};

export default newCard;
