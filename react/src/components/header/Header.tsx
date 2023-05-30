import { useMutateAuth } from "@/hooks/useMutateAuth";

const Header = () => {
  const { logoutMutation } = useMutateAuth();
  const logout = async () => {
    await logoutMutation.mutateAsync();
  };
  return (
    <header className="bg-blue-700 z-0 w-100%">
      <div className="mx-4 md:mx-12 py-2 h-12 md:h-14">
        <div className="flex items-center justify-between text-white font-bold">
          <div className="text-2xl md:text-4xl">StudyApp</div>
          <button className="flex">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15m3 0l3-3m0 0l-3-3m3 3H9"
              />
            </svg>
            <div className="ml-1" onClick={logout}>
              ログアウト
            </div>
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;
