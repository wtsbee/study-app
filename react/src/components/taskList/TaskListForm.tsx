import { useState } from "react";
import { useMutateTask } from "@/hooks/useMutateTask";
import { TaskList } from "@/types";

interface Props {
  section: TaskList;
}

const TaskListForm = ({ section: section }: Props) => {
  const [isEditList, setIsEditList] = useState(false);
  const [input, setInput] = useState(section.name);
  const { updateTaskListMutation } = useMutateTask();

  const openEditList = () => {
    setIsEditList(true);
  };

  const closeEditList = async (id: number) => {
    await updateTaskListMutation.mutateAsync({
      id: id,
      name: input,
    });
    setIsEditList(false);
  };

  const inputTaskList = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.target.value);
  };

  return (
    <>
      {isEditList ? (
        <div className="relative my-3">
          <input
            type="text"
            className="w-full py-2 pl-2 pr-7 rounded"
            placeholder="リスト名を入力"
            value={input}
            onChange={inputTaskList}
          />
          <button className="absolute inset-y-0 right-0 pl-4 pr-1 flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6 hover:bg-pink-200 cursor-pointer rounded"
              onClick={() => closeEditList(section.id as number)}
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      ) : (
        <div className="flex justify-between trello-section-title font-bold my-3 p-2">
          <span className="">{input}</span>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="w-6 h-6 hover:bg-pink-200 cursor-pointer rounded ml-2"
            onClick={openEditList}
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10"
            />
          </svg>
        </div>
      )}
    </>
  );
};

export default TaskListForm;
