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
  status: 'progress' | 'aborted' | 'done',
  dueDate: Date,
  isRepeat: boolean,
  interval?: number,
  notes?: string,
  execAt?: Date,
  createdAt?: Date,
  updatedAt?: Date,
}
