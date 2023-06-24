import axios from "axios";
import { useQuery } from "@tanstack/react-query";
import { TaskList, Task } from "../types";
import { useError } from "../hooks/useError";

export const useQueryTaskById = (id: string) => {
  const { switchErrorHandling } = useError();
  const getTaskById = async () => {
    const { data } = await axios<Task>({
      method: "get",
      url: `${import.meta.env.VITE_BACKEND_URL}/task/${id}`,
      withCredentials: true,
    });
    return data;
  };
  return useQuery<Task, Error>({
    queryKey: ["task", id],
    queryFn: getTaskById,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message);
      } else {
        switchErrorHandling(err.response.data);
      }
    },
  });
};

export const useQueryTasks = () => {
  const { switchErrorHandling } = useError();
  const getTasks = async () => {
    const { data } = await axios<TaskList[]>({
      method: "get",
      url: `${import.meta.env.VITE_BACKEND_URL}/tasks`,
      withCredentials: true,
    });
    return data;
  };
  return useQuery<TaskList[], Error>({
    queryKey: ["tasks"],
    queryFn: getTasks,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message);
      } else {
        switchErrorHandling(err.response.data);
      }
    },
  });
};
