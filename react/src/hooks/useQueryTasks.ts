import axios from "axios";
import { useQuery } from "@tanstack/react-query";
import { TaskList } from "../types";
import { useError } from "../hooks/useError";

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
