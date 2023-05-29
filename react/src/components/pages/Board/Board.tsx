import {
  DragDropContext,
  Draggable,
  Droppable,
  DropResult,
} from "react-beautiful-dnd";
import dummyData from "./dummyData";
import { useState } from "react";
import BoardCard from "./BoardCard";
import Header from "@/components/header/Header";

interface TaskData {
  id: string;
  title: string;
}

interface SectionData {
  id: string;
  title: string;
  tasks: TaskData[];
}

const Task = () => {
  const [data, setData] = useState<SectionData[]>(dummyData);

  const onDragEnd = (result: DropResult) => {
    if (!result.destination) return;
    const { source, destination } = result;

    if (source.droppableId !== destination.droppableId) {
      // 動かし始めたcolumnが違うcolumnに移動する場合
      // 動かし始めたcolumnの配列の番号を取得
      const sourceColIndex = data.findIndex((e) => e.id === source.droppableId);
      // 動かし終わったcolumnの配列の番号を取得
      const destinationColIndex = data.findIndex(
        (e) => e.id === destination.droppableId
      );

      const sourseCol = data[sourceColIndex];
      const destinationCol = data[destinationColIndex];

      // 動かし始めたタスクに属していたカラムの中のタスクを全て取得
      const sourceTask = [...sourseCol.tasks];
      // 動かし終わったタスクに属していたカラムの中のタスクを全て取得
      const destinationTask = [...destinationCol.tasks];

      // 前のカラムから削除
      const [removed] = sourceTask.splice(source.index, 1);
      // 後のカラムに追加
      destinationTask.splice(destination.index, 0, removed);

      data[sourceColIndex].tasks = sourceTask;
      data[destinationColIndex].tasks = destinationTask;

      setData(data);
    } else {
      // 同じカラム内でタスクを入れ替える場合
      const sourceColIndex = data.findIndex((e) => e.id === source.droppableId);
      const sourseCol = data[sourceColIndex];
      const sourceTask = [...sourseCol.tasks];

      const [removed] = sourceTask.splice(source.index, 1);
      sourceTask.splice(destination.index, 0, removed);

      data[sourceColIndex].tasks = sourceTask;

      setData(data);
    }
  };

  return (
    <>
      <Header></Header>
      <div className="mx-1">
        <h1 className="ml-2 mt-2 font-bold">タスク管理</h1>
        <DragDropContext onDragEnd={onDragEnd}>
          <div className="trello flex self-start">
            {data.map((section) => (
              <Droppable key={section.id} droppableId={section.id}>
                {(provided) => (
                  <div
                    className="trello-section w-96 bg-pink-400 m-1 p-2 rounded"
                    ref={provided.innerRef}
                    {...provided.droppableProps}
                  >
                    <div className="trello-section-title font-bold">
                      {section.title}
                    </div>
                    <div className="trello-section-content">
                      {section.tasks.map((task, index) => (
                        <Draggable
                          draggableId={task.id}
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
                                opacity: snapshot.isDragging ? "0.3" : "1",
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
            ))}
          </div>
        </DragDropContext>
      </div>
    </>
  );
};

export default Task;
