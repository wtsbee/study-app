import {
  DragDropContext,
  Draggable,
  Droppable,
  DropResult,
} from "react-beautiful-dnd";
import { useEffect, useState } from "react";
import BoardCard from "./BoardCard";
import Header from "@/components/header/Header";
import { useQueryTasks } from "@/hooks/useQueryTasks";
import { useMutateTask } from "@/hooks/useMutateTask";
import { Task, TaskList } from "@/types";

const Board = () => {
  const { data: resData, isLoading, isError } = useQueryTasks();
  const [data, setData] = useState<TaskList[]>([]);
  const { updateTaskMutation } = useMutateTask();

  useEffect(() => {
    if (resData) {
      setData(resData);
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
    if (source.droppableId == "board") {
      // 要素の移動
      const startIndex = source.index;
      const endIndex = destination.index;
      const element = newData.splice(startIndex, 1)[0];
      newData.splice(endIndex, 0, element);

      setData(newData);
    } else if (source.droppableId !== destination.droppableId) {
      // 動かし始めたcolumnが違うcolumnに移動する場合
      // 動かし始めたcolumnの配列の番号を取得
      const sourceColIndex = newData.findIndex(
        (e) => e.id.toString() === source.droppableId
      );
      // 動かし終わったcolumnの配列の番号を取得
      const destinationColIndex = newData.findIndex(
        (e) => e.id.toString() === destination.droppableId
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

      setData(newData);
    } else {
      // 同じカラム内でタスクを入れ替える場合
      const sourceColIndex = newData.findIndex(
        (e) => e.id.toString() === source.droppableId
      );
      const sourceCol = newData[sourceColIndex];
      const sourceTask = [...sourceCol.tasks];

      const [removed] = sourceTask.splice(source.index, 1);
      sourceTask.splice(destination.index, 0, removed);

      newData[sourceColIndex].tasks = sourceTask;

      setData(newData);
    }
    updateTaskMutation.mutate(newData);
  };

  return (
    <>
      <Header></Header>
      <div className="mx-1 overflow-auto min-h-[calc(100%_-_48px) md:min-h-[calc(100%_-_56px)]">
        <div className="ml-2">
          <h1 className="my-2 font-bold text-xl">タスク管理</h1>
          <button className="mb-1 py-2 px-4 rounded text-white bg-orange-500 font-bold">
            リストを追加
          </button>
        </div>
        <DragDropContext onDragEnd={onDragEnd}>
          <div className="trello ml-1">
            <Droppable droppableId="board" direction="horizontal" type="board">
              {(provided) => (
                <div
                  className="flex self-start"
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
                                <div className="trello-section-title font-bold">
                                  {section.name}
                                </div>
                                <div className="trello-section-content">
                                  {section.tasks?.map((task, index) => (
                                    <Draggable
                                      draggableId={task.id.toString()}
                                      index={index}
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
                                          <BoardCard>{task.title}</BoardCard>
                                        </div>
                                      )}
                                    </Draggable>
                                  ))}
                                </div>
                                {provided.placeholder}
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
