import Header from "@/components/header/Header";
import MarDownMain from "@/components/markdown/MarkDownMain";

const Markdown = () => {
  return (
    <>
      <Header />
      <div className="markdown-body markdown">
        <MarDownMain />
      </div>
    </>
  );
};

export default Markdown;
