import axios from "axios";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { Task, TaskDetail, TaskList, TaskListRequest } from "../types";
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
    (task: Task) =>
      axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/task/${task.id}`,
        data: { title: task.title },
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
  const deleteTaskMutation = useMutation(
    (id: number) =>
      axios({
        method: "delete",
        url: `${import.meta.env.VITE_BACKEND_URL}/task/${id}`,
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
  const createTaskDetailMutation = useMutation(
    (taskId: number) =>
      axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/task/${taskId}/detail`,
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
  const updateTaskDetailMutation = useMutation(
    (taskDetail: TaskDetail) =>
      axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/task/${
          taskDetail.task_id
        }/update_task_detail`,
        data: {
          id: taskDetail.id,
          detail: taskDetail.detail,
          task_id: taskDetail.task_id,
        },
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
  const updateTaskListMutation = useMutation((taskList: TaskListRequest) =>
    axios({
      method: "put",
      url: `${import.meta.env.VITE_BACKEND_URL}/task_list/${taskList.id}`,
      data: taskList,
      withCredentials: true,
    })
  );
  const updateTasksMutation = useMutation(
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
  const uploadImageMutation = useMutation(
    (formData: FormData) =>
      axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/task/upload`,
        data: formData,
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
  return {
    createTaskMutation,
    updateTaskMutation,
    deleteTaskMutation,
    createTaskDetailMutation,
    updateTaskDetailMutation,
    updateTaskListMutation,
    updateTasksMutation,
    deleteTaskListMutation,
    uploadImageMutation,
  };
};
