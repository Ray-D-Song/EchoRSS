interface AuthForm {
  account: string
  password: string
}

interface User {
  role: string
  token: string
  refreshToken: string
  username: string
}

interface UserListItem {
  id: string
  username: string
  password: string
  deleted: 0 | 1
}
