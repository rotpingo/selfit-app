export interface ISign {
  email: string,
  password: string
}

export interface IUser {
  id?: number,
  name?: string,
  email: string,
  password: string,
  createdAt?: Date,
  updatedAt?: Date,
}

export interface INote {
  id?: number
  title: string
  content: string
  createdAt?: Date
  updatedAt?: Date
}

export interface IFabOption {
  label: string;
  icon?: string;
  action: () => void;
}

export interface ITask {
  id?: number,
  parentId?: number,
  title: string,
  content: string,
  status?: 'progress' | 'aborted' | 'done',
  isRepeat: boolean,
  interval?: number,
  notes?: string,
  dueDate: string,
  execAt?: Date,
  createdAt?: Date,
  updatedAt?: Date,
  userId: number,
}
