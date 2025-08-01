export interface INote {
  id?: number
  title: string
  content: string
}

export interface IFabOption {
  label: string;
  icon?: string;
  action: () => void;
}
