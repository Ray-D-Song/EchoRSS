import { createContext } from 'react'

export const AppCtx = createContext<{
  user: User | null
  setUser: (user: User | null) => void
  removeUser: () => void
}>({
  user: null,
  setUser: () => {},
  removeUser: () => {},
})
