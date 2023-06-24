import { useLocation } from "react-router-dom";
import Header from "@/components/header/Header";
import BoardTaskMain from "@/components/boardTask/BoardTaskMain";
import { useEffect, useRef, useState } from "react";
import { useMutateTask } from "@/hooks/useMutateTask";
import { useQueryTaskById } from "@/hooks/useQueryTasks";
import { Task } from "@/types";

const BoardTask = () => {
  const location = useLocation();
  const taskId = location.pathname.split("/")[2];
  const [isEdit, setIsEdit] = useState(false);
  const insideRef = useRef<HTMLInputElement>(null);
  const { updateTaskMutation } = useMutateTask();
  const { data } = useQueryTaskById(taskId);
  const [input, setInput] = useState("");

  const editTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.target.value);
  };

  useEffect(() => {
    if (data) {
      console.log("aaaaaaa");
      setInput(data.title);
    }
  }, [data]);

  useEffect(() => {
    //対象の要素を取得
    const el = insideRef.current;

    //対象の要素がなければ何もしない
    if (!el) return;

    //クリックした時に実行する関数
    const handleClickOutside = async (e: MouseEvent) => {
      if (!el?.contains(e.target as Node)) {
        //ここに外側をクリックしたときの処理
        if (isEdit) {
          console.log("外側クリック");
          await updateTaskMutation.mutateAsync({
            ...(data as Task),
            title: input,
          });
          setIsEdit(false);
        }
      } else {
        //ここに内側をクリックしたきの処理
        if (!isEdit) {
          console.log("内側クリック");
          setIsEdit(true);
        }
      }
    };

    //クリックイベントを設定
    document.addEventListener("click", handleClickOutside);

    //クリーンアップ関数
    return () => {
      //コンポーネントがアンマウント、再レンダリングされたときにクリックイベントを削除
      document.removeEventListener("click", handleClickOutside);
    };
  }, [insideRef, isEdit, input]);

  return (
    <>
      <Header />
      <input
        ref={insideRef}
        onChange={editTitle}
        value={input}
        className="w-full pt-12 md:pt-16 px-5 pb-2 border-b-2 border-black text-2xl"
      ></input>
      <div className="markdown-body markdown">
        <BoardTaskMain />
      </div>
    </>
  );
};

export default BoardTask;
