// import { useLocation } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import "github-markdown-css";
import emoji from "remark-emoji";
import { useEffect, useRef, useState } from "react";
import { useQueryTaskDetailByTaskId } from "@/hooks/useQueryTasks";
import { useMutateTask } from "@/hooks/useMutateTask";
import { TaskDetail } from "@/types";

const markdownString = ``;

const html = markdownString.replace(/\n/g, "<br>");

const MarkdownMain = () => {
  const [text, setText] = useState(markdownString);
  const taskId = location.pathname.split("/")[2];
  const { data: taskDetail, refetch } = useQueryTaskDetailByTaskId(taskId);
  const { createTaskDetailMutation, updateTaskDetailMutation } =
    useMutateTask();
  const socketRef = useRef<WebSocket>();

  const inputText = async (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value);
    await updateTaskDetailMutation.mutateAsync({
      ...(taskDetail as TaskDetail),
      detail: e.target.value,
    });
    socketRef.current?.send(JSON.stringify(e.target.value));
  };

  useEffect(() => {
    if (taskDetail) {
      if (taskDetail == "RecordNotFound") {
        (async () => {
          await createTaskDetailMutation.mutateAsync(Number(taskId));
          refetch();
        })();
      } else {
        const detail = (taskDetail as TaskDetail).detail;
        setText(detail);
      }
    }
  }, [taskDetail]);

  useEffect(() => {
    socketRef.current = new WebSocket(
      `${import.meta.env.VITE_BACKEND_WEBSOCKET_URL}/task/${taskId}/ws`
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
      setText(JSON.parse(event.data));
    };

    // コンポーネントのアンマウント時にWebSocket接続をクローズ
    return () => {
      if (socketRef.current === null) {
        return;
      }
      socketRef.current?.close();
    };
  }, []);

  return (
    <>
      <div className="flex h-screen fixed left-0 right-0">
        <textarea
          onChange={inputText}
          value={text}
          spellCheck={false}
          className="resize-none w-1/2 px-5 pt-2 pb-10 text-white bg-light-black overflow-scroll border-none outline-none"
        ></textarea>
        <div className="w-1/2 px-5 pt-2 pb-10 overflow-scroll">
          <ReactMarkdown remarkPlugins={[remarkGfm, emoji]}>
            {text}
          </ReactMarkdown>
        </div>
      </div>
    </>
  );
};

export default MarkdownMain;
