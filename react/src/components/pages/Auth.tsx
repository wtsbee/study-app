import { useState, FormEvent } from "react";
import { CheckBadgeIcon, ArrowPathIcon } from "@heroicons/react/24/solid";
import { useMutateAuth } from "@/hooks/useMutateAuth";

const Auth = () => {
  const [email, setEmail] = useState("");
  const [pw, setPw] = useState("");
  const [isLogin, setIsLogin] = useState(true);
  const { loginMutation, registerMutation } = useMutateAuth();

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (isLogin) {
      loginMutation.mutate({
        email: email,
        password: pw,
      });
    } else {
      await registerMutation
        .mutateAsync({
          email: email,
          password: pw,
        })
        .then(() =>
          loginMutation.mutate({
            email: email,
            password: pw,
          })
        );
    }
  };
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center">
        <CheckBadgeIcon className="h-8 w-8 mr-2 text-blue-500" />
        <span className="text-center text-3xl font-extrabold">StudyApp</span>
      </div>
      <h2 className="my-6">{isLogin ? "ログイン" : "アカウント作成"}</h2>
      <form onSubmit={submitAuthHandler}>
        <div>
          <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="email"
            type="email"
            autoFocus
            placeholder="メールアドレス"
            onChange={(e) => setEmail(e.target.value)}
            value={email}
          />
        </div>
        <div>
          <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="password"
            type="password"
            placeholder="パスワード"
            onChange={(e) => setPw(e.target.value)}
            value={pw}
          />
        </div>
        <div className="flex justify-center my-2">
          <button
            className="disabled:opacity-40 py-2 px-4 rounded text-white bg-orange-500"
            disabled={!email || !pw}
            type="submit"
          >
            {isLogin ? "ログイン" : "サインアップ"}
          </button>
        </div>
      </form>
      <ArrowPathIcon
        onClick={() => setIsLogin(!isLogin)}
        className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
      />
    </div>
  );
};

export default Auth;
