export interface ISign {
  email: string,
  password: string
}

export interface IAuthResponse {
  message: string;
  token: string;
}

export interface IUser {
  name?: string,
  email: string,
  password: string,
  createdAt?: Date,
  updatedAt?: Date,
}

export interface INote {
  id: number
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
}

export interface ITracker {
  id?: number,
  title: string,
  notes: string,
  startDate?: string,
  currentStreak?: number,
  bestStreak?: number,
  createdAt?: Date,
  updatedAt?: Date,
}

export interface IWeatherRequest {
  id?: number,
  name: string,
  country?: string,
}

export interface IWeatherResponse {
  name: string;
  main: {
    temp: number;
    feels_like: number;
    temp_min: number;
    temp_max: number;
    humidity: number;
  };
  wind: {
    speed: number;
    deg: number;
  };
  sys: {
    country: string;
    sunrise: number;
    sunset: number;
  };
  weather: {
    id: number;
    main: string;
    description: string;
  }[];
  clouds: {
    all: number;
  };
  coord: {
    lon: number;
    lat: number;
  };
}
