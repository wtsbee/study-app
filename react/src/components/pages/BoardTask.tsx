import { useLocation } from "react-router-dom";
import Header from "@/components/header/Header";
import BoardTaskMain from "@/components/boardTask/BoardTaskMain";

const BoardTask = () => {
  const location = useLocation();

  return (
    <>
      <Header />
      <div className="pt-12 md:pt-16 px-5 pb-2 border-b-2 border-black text-2xl">
        {location.state}
      </div>
      <div className="markdown-body markdown">
        <BoardTaskMain />
      </div>
    </>
  );
};

export default BoardTask;
