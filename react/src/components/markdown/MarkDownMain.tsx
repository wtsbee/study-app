import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import "github-markdown-css";
import emoji from "remark-emoji";
import { useState } from "react";

const markdownString = ``;

const html = markdownString.replace(/\n/g, "<br>");

const MarkdownMain = () => {
  const [text, setText] = useState(markdownString);

  const inputText = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value);
  };

  return (
    <>
      <div className="flex h-screen pt-12 md:pt-14 fixed left-0 right-0">
        <textarea
          onChange={inputText}
          value={text}
          spellCheck={false}
          className="resize-none w-1/2 px-5 pt-2 pb-10 text-white bg-light-black overflow-scroll border-none outline-none"
        ></textarea>
        <div className="w-1/2 px-5 pt-2 pb-10 overflow-scroll">
          <ReactMarkdown remarkPlugins={[remarkGfm, emoji]}>
            {text}
          </ReactMarkdown>
        </div>
      </div>
    </>
  );
};

export default MarkdownMain;
