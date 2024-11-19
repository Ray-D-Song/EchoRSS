import { useNavigate } from '@/router'
import { useEffect, useState } from 'react'
import { useMedia } from 'react-use'

function useView() {
  const isMobile = useMedia('(max-width: 768px)')
  const [currentView, setCurrentView] = useState(isMobile ? 'mobile' : 'desktop')
  const navigate = useNavigate()
  useEffect(() => {
    if (isMobile && currentView !== 'mobile') {
      setCurrentView('mobile')
      navigate('/mobile/home')
    } else if (!isMobile && currentView !== 'desktop') {
      setCurrentView('desktop')
      navigate('/home')
    }
  }, [isMobile])
}

export default useView