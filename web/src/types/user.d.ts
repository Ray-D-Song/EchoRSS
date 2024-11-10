interface AuthForm {
  account: string
  password: string
}

interface User {
  id: number
  role: string
  token: string
  refreshToken: string
  username: string
}