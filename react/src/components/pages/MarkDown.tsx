import Header from "@/components/header/Header";
import MarDownMain from "@/components/markdown/MarkDownMain";

const Markdown = () => {
  return (
    <>
      <Header />
      <div className="markdown-body markdown ml-2 mt-2">
        <MarDownMain />
      </div>
    </>
  );
};

export default Markdown;
