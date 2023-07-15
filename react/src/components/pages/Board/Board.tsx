import {
  DragDropContext,
  Draggable,
  Droppable,
  DropResult,
} from "react-beautiful-dnd";
import { useEffect, useRef, useState } from "react";
import BoardCard from "./BoardCard";
import NewCard from "./newCard";
import Header from "@/components/header/Header";
import { useQueryTasks } from "@/hooks/useQueryTasks";
import { useMutateTask } from "@/hooks/useMutateTask";
import { Task, TaskList } from "@/types";

const Board = () => {
  const { data: resData, isLoading, isError, refetch } = useQueryTasks();
  const { updateTasksMutation, deleteTaskListMutation } = useMutateTask();
  const [data, setData] = useState<TaskList[]>([]);
  const [isEdit, setIsEdit] = useState(false);
  const [input, setInput] = useState("");

  const socketRef = useRef<WebSocket>();

  const openList = () => {
    setIsEdit(true);
  };

  const closeList = () => {
    setIsEdit(false);
  };

  const inputText = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.target.value);
  };

  const addList = async () => {
    if (input !== "") {
      await updateTasksMutation.mutateAsync([
        ...data,
        { name: input, tasks: [] },
      ]);
      setInput("");
      await refetch();
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.nativeEvent.isComposing || e.key !== "Enter") return;
    addList();
  };

  const deleteList = (tasklist: TaskList, index: number) => {
    const newData = [...data];
    newData.splice(index, 1);
    deleteTaskListMutation.mutate(tasklist.id as number);
    socketRef.current?.send(JSON.stringify(newData));
  };

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
    if (resData) {
      setData(resData);
      if (socketRef.current?.readyState === 1) {
        socketRef.current?.send(JSON.stringify(resData));
      }
    }
  }, [resData]);

  if (isLoading) {
    return <div>Loading...</div>;
  }
  if (isError) {
    return <div>Error occurred</div>;
  }
  const onDragEnd = (result: DropResult) => {
    if (!result.destination) return;
    const { source, destination } = result;

    const newData = [...data];
    if (source.droppableId === "board") {
      // 要素の移動
      const startIndex = source.index;
      const endIndex = destination.index;
      const element = newData.splice(startIndex, 1)[0];
      newData.splice(endIndex, 0, element);

      // setData(newData);
    } else if (source.droppableId !== destination.droppableId) {
      // 動かし始めたcolumnが違うcolumnに移動する場合
      // 動かし始めたcolumnの配列の番号を取得
      const sourceColIndex = newData.findIndex(
        (e) => e.id?.toString() === source.droppableId
      );
      // 動かし終わったcolumnの配列の番号を取得
      const destinationColIndex = newData.findIndex(
        (e) => e.id?.toString() === destination.droppableId
      );

      const sourceCol = newData[sourceColIndex];
      const destinationCol = newData[destinationColIndex];

      // 動かし始めたタスクに属していたカラムの中のタスクを全て取得
      const sourceTask = [...sourceCol.tasks];
      // TODO:2番目以降のリストの箱が空の場合、tasksが空配列ではなくnullになるため要見直し
      if (destinationCol.tasks == null) {
        destinationCol.tasks = [];
      }
      // 動かし終わったタスクに属していたカラムの中のタスクを全て取得
      const destinationTask = [...destinationCol.tasks];

      // 前のカラムから削除
      const [removed] = sourceTask.splice(source.index, 1);
      // 後のカラムに追加
      destinationTask.splice(destination.index, 0, removed);

      newData[sourceColIndex].tasks = sourceTask;
      newData[destinationColIndex].tasks = destinationTask;

      // setData(newData);
    } else {
      // 同じカラム内でタスクを入れ替える場合
      const sourceColIndex = newData.findIndex(
        (e) => e.id?.toString() === source.droppableId
      );
      const sourceCol = newData[sourceColIndex];
      const sourceTask = [...sourceCol.tasks];

      const [removed] = sourceTask.splice(source.index, 1);
      sourceTask.splice(destination.index, 0, removed);

      newData[sourceColIndex].tasks = sourceTask;

      // setData(newData);
    }
    updateTasksMutation.mutate(newData);
    socketRef.current?.send(JSON.stringify(newData));
  };

  return (
    <>
      <Header></Header>
      <div className="mx-1 overflow-auto min-h-[calc(100%_-_48px) md:min-h-[calc(100%_-_56px)] pt-12 md:pt-14">
        <div className="ml-2">
          <h1 className="my-2 font-bold text-xl">タスク管理</h1>
          {isEdit ? (
            <div className="flex flex-col w-80">
              <input
                onChange={inputText}
                onKeyDown={handleKeyDown}
                value={input}
                placeholder="リスト名を入力してください"
                className="mb-2 p-2 rounded border border-gray-500"
              />
              <div className="flex items-center">
                <button
                  onClick={addList}
                  className="mb-1 py-2 px-4 rounded text-white bg-orange-500 font-bold"
                >
                  リストを追加
                </button>
                <div onClick={closeList} className="hover:cursor-pointer">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="w-6 h-6"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M6 18L18 6M6 6l12 12"
                    />
                  </svg>
                </div>
              </div>
            </div>
          ) : (
            <button
              onClick={openList}
              className="mb-1 py-2 px-4 rounded text-white bg-orange-500 font-bold"
            >
              リストの新規作成
            </button>
          )}
        </div>
        <DragDropContext onDragEnd={onDragEnd}>
          <div className="trello ml-1 mb-10">
            <Droppable droppableId="board" direction="horizontal" type="board">
              {(provided) => (
                <div
                  className="flex items-start"
                  ref={provided.innerRef}
                  {...provided.droppableProps}
                >
                  {data?.map((section, index) => (
                    <Draggable
                      draggableId={`${section.id.toString()} + ${section.name}`}
                      index={index}
                      key={section.id}
                    >
                      {(provided) => (
                        <div
                          ref={provided.innerRef}
                          {...provided.draggableProps}
                          {...provided.dragHandleProps}
                        >
                          <Droppable
                            key={section.id}
                            droppableId={section.id.toString()}
                          >
                            {(provided) => (
                              <div
                                className="trello-section w-96 bg-pink-400 m-1 p-2 rounded"
                                ref={provided.innerRef}
                                {...provided.droppableProps}
                              >
                                <div className="flex justify-between trello-section-title font-bold">
                                  <span>{section.name}</span>
                                  <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    strokeWidth={1.5}
                                    stroke="currentColor"
                                    className="w-6 h-6 hover:bg-pink-200 cursor-pointer rounded"
                                    onClick={() => deleteList(section, index)}
                                  >
                                    <path
                                      strokeLinecap="round"
                                      strokeLinejoin="round"
                                      d="M6 18L18 6M6 6l12 12"
                                    />
                                  </svg>
                                </div>
                                <div className="trello-section-content">
                                  {section.tasks?.map((task, cardIndex) => (
                                    <Draggable
                                      draggableId={task.id.toString()}
                                      index={cardIndex}
                                      key={task.id}
                                    >
                                      {(provided, snapshot) => (
                                        <div
                                          ref={provided.innerRef}
                                          {...provided.draggableProps}
                                          {...provided.dragHandleProps}
                                          style={{
                                            ...provided.draggableProps.style,
                                            opacity: snapshot.isDragging
                                              ? "0.3"
                                              : "1",
                                          }}
                                        >
                                          <BoardCard>{task}</BoardCard>
                                        </div>
                                      )}
                                    </Draggable>
                                  ))}
                                </div>
                                {provided.placeholder}
                                <NewCard taskList={section} />
                              </div>
                            )}
                          </Droppable>
                        </div>
                      )}
                    </Draggable>
                  ))}
                  {provided.placeholder}
                </div>
              )}
            </Droppable>
          </div>
        </DragDropContext>
      </div>
    </>
  );
};

export default Board;
