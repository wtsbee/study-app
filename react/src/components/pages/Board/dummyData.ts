import { v4 as uuidv4 } from "uuid";

interface Task {
  id: string;
  title: string;
}

interface Column {
  id: string;
  title: string;
  tasks: Task[];
}

const dummyData: Column[] = [
  {
    id: uuidv4(),
    title: "ToDo",
    tasks: [
      {
        id: uuidv4(),
        title: "ToDo1",
      },
      {
        id: uuidv4(),
        title: "ToDo2",
      },
      {
        id: uuidv4(),
        title: "ToDo3",
      },
    ],
  },
  {
    id: uuidv4(),
    title: "作業中",
    tasks: [
      {
        id: uuidv4(),
        title: "作業中1",
      },
      {
        id: uuidv4(),
        title: "作業中2",
      },
    ],
  },
  {
    id: uuidv4(),
    title: "完了",
    tasks: [
      {
        id: uuidv4(),
        title: "完了1",
      },
    ],
  },
];

export default dummyData;
