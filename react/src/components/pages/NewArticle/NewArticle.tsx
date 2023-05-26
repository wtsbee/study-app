const NewArticle = () => {
  return (
    <div className="main">
      <div className="mx-10 py-5">
        <h1 className="text-3xl md:text-4xl font-bold text-center mb-5">
          投稿画面
        </h1>
        <div>
          <label className="font-bold">タイトル</label>
          <br />
          <input type="text" className="border border-gray-400 my-2 w-full" />
        </div>
        <div>
          <label className="font-bold">内容</label>
          <br />
          <textarea className="border border-gray-400 my-2 w-full h-40 md:h-60" />
        </div>
        <div className="flex justify-center">
          <button className="border text-white font-bold bg-orange-500 my-2 p-3 w-1/4 md:w-2/12 rounded-md">
            登録
          </button>
        </div>
      </div>
    </div>
  );
};

export default NewArticle;
