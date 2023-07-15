// import { useLocation } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import "github-markdown-css";
import emoji from "remark-emoji";
import { useEffect, useRef, useState } from "react";
import { useQueryTaskDetailByTaskId } from "@/hooks/useQueryTasks";
import { useMutateTask } from "@/hooks/useMutateTask";
import { TaskDetail } from "@/types";
import { useCallback } from "react";
import { useDropzone } from "react-dropzone";

const markdownString = ``;

const html = markdownString.replace(/\n/g, "<br>");

const MarkdownMain = () => {
  const [text, setText] = useState(markdownString);
  const taskId = location.pathname.split("/")[2];
  const { data: taskDetail, refetch } = useQueryTaskDetailByTaskId(taskId);
  const {
    createTaskDetailMutation,
    updateTaskDetailMutation,
    uploadImageMutation,
  } = useMutateTask();
  const socketRef = useRef<WebSocket>();
  const [imagePath, setImagePath] = useState("");
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  const inputText = async (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value);
    await updateTaskDetailMutation.mutateAsync({
      ...(taskDetail as TaskDetail),
      detail: e.target.value,
    });
    socketRef.current?.send(JSON.stringify(e.target.value));
  };

  const onDrop = useCallback(async (acceptedFiles: File[]) => {
    const formData = new FormData();
    formData.append("image", acceptedFiles[0]);
    const res = await uploadImageMutation.mutateAsync(formData);
    // TODO:以下のように記述したいが、insertText内でtaskDetailがundefinedとなるためuseEffectで書いているため要改善
    setImagePath(res.data);
    // const imgUrl = `![](${import.meta.env.VITE_BACKEND_URL}/task/${res.data})`;
    // insertText(imgUrl);
  }, []);

  const insertText = async (imgUrl: string) => {
    const textarea = textareaRef.current;
    if (textarea) {
      const startPos = textarea.selectionStart;
      const endPos = textarea.selectionEnd;

      const currentValue = textarea.value;
      const newValue =
        currentValue.substring(0, startPos) +
        imgUrl +
        currentValue.substring(endPos);
      setText(newValue);
      await updateTaskDetailMutation.mutateAsync({
        ...(taskDetail as TaskDetail),
        detail: newValue,
      });
      socketRef.current?.send(JSON.stringify(newValue));
    }
  };

  const { getRootProps } = useDropzone({
    onDrop,
    noClick: true, // クリックによるアップロードを無効化
    accept: {
      "image/png": [".png", ".jpg", ".jpeg"],
    },
  });

  useEffect(() => {
    if (imagePath != "") {
      const imgUrl = `![](${
        import.meta.env.VITE_BACKEND_URL
      }/task/${imagePath})`;
      insertText(imgUrl);
    }
  }, [imagePath]);

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
          {...getRootProps()}
          onChange={inputText}
          value={text}
          spellCheck={false}
          ref={textareaRef}
          className="resize-none w-1/2 px-5 pt-2 pb-28 text-white bg-light-black overflow-scroll border-none outline-none"
        ></textarea>
        <div className="w-1/2 px-5 pt-2 pb-28 overflow-scroll">
          <ReactMarkdown
            remarkPlugins={[remarkGfm, emoji]}
            className="text-black"
          >
            {text}
          </ReactMarkdown>
        </div>
      </div>
    </>
  );
};

export default MarkdownMain;
