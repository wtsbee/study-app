import axios from "axios";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { Task, TaskList } from "../types";
import { useError } from "../hooks/useError";

export const useMutateTask = () => {
  const queryClient = useQueryClient();
  const { switchErrorHandling } = useError();

  const createTaskMutation = useMutation(
    (task: Task) =>
      axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/task`,
        data: task,
        withCredentials: true,
      }),
    {
      onSuccess: () => {},
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  const updateTaskMutation = useMutation(
    (task: TaskList[]) =>
      axios<TaskList[]>({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/tasks`,
        data: task,
        withCredentials: true,
      }),
    {
      onSuccess: () => {},
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  const deleteTaskListMutation = useMutation(
    (id: number) =>
      axios({
        method: "delete",
        url: `${import.meta.env.VITE_BACKEND_URL}/tasks/${id}`,
        withCredentials: true,
      }),
    {
      onSuccess: () => {},
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  return { createTaskMutation, updateTaskMutation, deleteTaskListMutation };
};
