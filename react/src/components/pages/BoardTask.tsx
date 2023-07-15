import { useLocation } from "react-router-dom";
import Header from "@/components/header/Header";
import BoardTaskMain from "@/components/boardTask/BoardTaskMain";
import { useEffect, useRef, useState } from "react";
import { useMutateTask } from "@/hooks/useMutateTask";
import { useQueryTaskById, useQueryTasks } from "@/hooks/useQueryTasks";
import { Task, TaskList } from "@/types";

const BoardTask = () => {
  const location = useLocation();
  const taskId = location.pathname.split("/")[2];
  const [isEdit, setIsEdit] = useState(false);
  const insideRef = useRef<HTMLInputElement>(null);
  const { updateTaskMutation, updateTasksMutation } = useMutateTask();
  const { data: taskListArray } = useQueryTasks();
  const [data, setData] = useState<TaskList[]>([]);
  const { data: task } = useQueryTaskById(taskId);
  const [input, setInput] = useState("");
  const socketRef = useRef<WebSocket>();

  const editTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.target.value);
  };

  const handleUpdateTitle = async () => {
    await updateTaskMutation.mutateAsync({
      ...(task as Task),
      title: input,
    });
    setIsEdit(false);

    const newData = [...data];

    let listIndex = 0; // カードが格納されているリストのindex
    let cardIndex = 0; // リスト内でのカードのindex

    for (let i = 0; i < newData.length; i++) {
      if (newData[i].tasks) {
        cardIndex = newData[i].tasks.findIndex(
          (task) => task.id === Number(taskId)
        );
        if (cardIndex !== -1) {
          listIndex = i;
          break;
        }
      }
    }

    newData[listIndex].tasks[cardIndex].title = input;
    updateTasksMutation.mutate(newData);
    socketRef.current?.send(JSON.stringify(newData));
  };

  useEffect(() => {
    if (task) {
      setInput(task.title);
    }
  }, [task]);

  useEffect(() => {
    socketRef.current = new WebSocket(
      `${import.meta.env.VITE_BACKEND_WEBSOCKET_URL}/tasks/ws`
    );

    socketRef.current.onopen = () => {
      console.log("ws接続");
    };

    socketRef.current.onclose = () => {
      console.log("ws切断");
    };

    // メッセージ受信時の処理
    socketRef.current.onmessage = (event) => {
      console.log("ws受信");

      const taskListArray: TaskList[] = JSON.parse(event.data);
      const str = location.pathname;
      const pattern = /\/board\/(\d+)/;
      const match = str.match(pattern);

      const taskList = taskListArray.find(
        (obj) =>
          obj.tasks && obj.tasks.find((task) => task.id === Number(taskId))
      );

      let title: string = "";
      if (taskList) {
        const task = taskList.tasks.find((task) => task.id === Number(taskId));
        if (task != undefined) {
          title = task.title;
        }
      }

      setInput(title);
      setData(JSON.parse(event.data));
    };

    // コンポーネントのアンマウント時にWebSocket接続をクローズ
    return () => {
      if (socketRef.current === null) {
        return;
      }
      socketRef.current?.close();
    };
  }, []);

  useEffect(() => {
    if (taskListArray) {
      setData(taskListArray);
    }
  }, [taskListArray]);

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
          handleUpdateTitle();
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
