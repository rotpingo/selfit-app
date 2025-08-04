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
