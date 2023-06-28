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
  task_list_id?: number;
  rank?: number;
};

export type TaskDetail = {
  id: number;
  detail: string;
  task_id?: number;
};
