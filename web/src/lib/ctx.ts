import { createContext } from 'react'

export const AppCtx = createContext<{
  refreshFeeds: () => void
  user: User | null
  setUser: (user: User | null) => void
  removeUser: () => void
}>({
  refreshFeeds: () => {},
  user: null,
  setUser: () => {},
  removeUser: () => {},
})
