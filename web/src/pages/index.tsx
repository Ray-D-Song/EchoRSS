import { AppCtx } from '@/lib/ctx'
import { useContext } from 'react'
import { Navigate } from 'react-router-dom'
import { useMedia } from 'react-use'

function Index() {
  const { user } = useContext(AppCtx)
  const isMobile = useMedia('(max-width: 768px)')
  return user ? <Navigate to={isMobile ? '/mobile/home' : '/home'} /> : <Navigate to="/login" />
}

export default Index
