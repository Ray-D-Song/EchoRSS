import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import '@/styles/globals.css'
import { Routes } from '@generouted/react-router'
import { ThemeProvider } from './components/theme-provider'
import { Toaster } from 'react-hot-toast'
import { AppCtx } from './lib/ctx'
import { useLocalStorage } from 'react-use'

function RenderWrapper() {
  const [user, setUser, removeUser] = useLocalStorage<User | null>('user', null)
  return (
    <StrictMode>
      <ThemeProvider defaultTheme='dark' storageKey='echo-rss-theme'>
        <Toaster position='top-center' />
      <AppCtx.Provider value={{
        user: user ?? null,
        setUser,
        removeUser,
      }}>
        <Routes />
        </AppCtx.Provider>
      </ThemeProvider>
    </StrictMode>
  )
}

createRoot(document.getElementById('root')!).render(
  <RenderWrapper />
)
