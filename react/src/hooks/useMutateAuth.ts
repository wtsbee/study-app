import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { Credential } from "../types";
import { useError } from "../hooks/useError";

export const useMutateAuth = () => {
  const navigate = useNavigate();
  const { switchErrorHandling } = useError();
  const loginMutation = useMutation(
    async (user: Credential) =>
      await axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/login`,
        data: user,
      }),
    {
      onSuccess: () => {
        navigate("/board");
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  const registerMutation = useMutation(
    async (user: Credential) =>
      await axios({
        method: "post",
        url: `${import.meta.env.VITE_BACKEND_URL}/signup`,
        data: user,
      }),
    {
      onError: (err: any) => {
        console.log("err:", err);
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  const logoutMutation = useMutation(
    async () => await axios.post(`${import.meta.env.VITE_BACKEND_URL}/logout`),
    {
      onSuccess: () => {
        navigate("/");
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  return { loginMutation, registerMutation, logoutMutation };
};
