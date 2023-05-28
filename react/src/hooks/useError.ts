import axios from "axios";
import { useNavigate } from "react-router-dom";
import { CsrfToken } from "../types";

export const useError = () => {
  const navigate = useNavigate();
  const getCsrfToken = async () => {
    const { data } = await axios<CsrfToken>({
      method: "get",
      url: `${import.meta.env.VITE_BACKEND_URL}/csrf`,
    });
    axios.defaults.headers.common["X-CSRF-TOKEN"] = data.csrf_token;
  };
  const switchErrorHandling = (msg: string) => {
    switch (msg) {
      case "invalid csrf token":
        getCsrfToken();
        alert("CSRF token is invalid, please try again");
        break;
      case "invalid or expired jwt":
        alert("access token expired, please login");
        navigate("/");
        break;
      case "missing or malformed jwt":
        alert("access token is not valid, please login");
        navigate("/");
        break;
      case "duplicated key not allowed":
        alert("email already exist, please use another one");
        break;
      case "crypto/bcrypt: hashedPassword is not the hash of the given password":
        alert("password is not correct");
        break;
      case "record not found":
        alert("email is not correct");
        break;
      default:
        alert(msg);
    }
  };
  return { switchErrorHandling };
};
