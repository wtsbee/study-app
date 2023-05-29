interface ChildrenProps {
  children: string;
}

const BoardCard = ({ children }: ChildrenProps) => {
  return <div className="card p-5 bg-pink-200 rounded mt-2">{children}</div>;
};

export default BoardCard;
