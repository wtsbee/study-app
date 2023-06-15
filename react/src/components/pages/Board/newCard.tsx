import { useState } from "react";
import { Task, TaskList } from "@/types";

type Props = {
  taskList: TaskList;
  index: number;
};

const newCard = ({ taskList, index }: Props) => {
  const [isEdit, setIsEdit] = useState(false);

  const openCard = () => {
    console.log(taskList, index);
    setIsEdit(true);
  };

  const closeCard = () => {
    setIsEdit(false);
  };

  return (
    <>
      {isEdit ? (
        <>
          <div className="mt-2 p-5 bg-pink-300 rounded">
            <textarea
              placeholder="タイトルを入力してください"
              className=" w-full placeholder-black bg-pink-300 outline-none"
            />
          </div>
          <div className="flex items-center mt-2 ">
            <button className="py-2 px-4 rounded text-white bg-blue-500 font-bold hover:cursor-pointer">
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
          className="flex mt-2 py-2 rounded hover:bg-pink-300"
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
