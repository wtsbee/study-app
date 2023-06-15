export type CsrfToken = {
  csrf_token: string;
};

export type Credential = {
  email: string;
  password: string;
};

export type TaskList = {
  id?: number;
  name: string;
  tasks: Task[];
};

export type Task = {
  id: number | null;
  title: string;
};
