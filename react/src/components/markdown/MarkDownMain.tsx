import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import "github-markdown-css";
import emoji from "remark-emoji";

const markdownString = `
# GFM
- [https://example.com](https://example.com)

- :smile:

## aaa
- **aaa**
  - aaa
    - aaa
      - aaa
* aaa

### bbb

#### ccc

- **ccc**
- :star:

Lorem ipsum dolor sit amet consectetur, adipisicing elit. Magni, nemo!

## Autolink literals

www.example.com, https://example.com, and contact@example.com.

## Footnote

A note[^1]

[^1]: Big note.

## Strikethrough

~one~ or ~~two~~ tildes.

## Table

| a | b  |  c |  d  |
| - | :- | -: | :-: |

## Tasklist

* [ ] to do
* [x] done
`;

const html = markdownString.replace(/\n/g, "<br>");

const MarkdownMain = () => {
  return (
    <>
      <div className="flex h-[calc(100vh-3rem)] md:h-[calc(100vh-3.5rem)]">
        <div
          className="w-1/2 px-5 pb-10 text-white bg-light-black overflow-scroll"
          dangerouslySetInnerHTML={{ __html: html }}
        />
        <div className="w-1/2 px-5 pb-10 overflow-scroll">
          <ReactMarkdown remarkPlugins={[remarkGfm, emoji]}>
            {markdownString}
          </ReactMarkdown>
        </div>
      </div>
    </>
  );
};

export default MarkdownMain;
