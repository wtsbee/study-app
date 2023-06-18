import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import Header from "@/components/header/Header";
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

const MarkdownMain = () => {
  return (
    <>
      <ReactMarkdown remarkPlugins={[remarkGfm, emoji]}>
        {markdownString}
      </ReactMarkdown>
    </>
  );
};

export default MarkdownMain;
