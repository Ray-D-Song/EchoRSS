import fetcher from '@/lib/fetcher'
import { useEffect, useState } from 'react'

interface UseFetchOptions<T> {
  onSuccess?: (data: T) => void
  onError?: (error?: Error) => void
  onFinally?: () => void
  immediate?: boolean
}

interface UseFetchReturn<T> {
  data: T | null
  loading: boolean
  run: () => void
  setData: (data: T | null) => void
}

function useFetch<T>(url: string, reqOptions: RequestInit, hookOptions: UseFetchOptions<T> = {}): UseFetchReturn<T> {
  const [data, setData] = useState<T | null>(null)
  const [loading, setLoading] = useState(false)
  const [reqTr, setReqTr] = useState(0)

  useEffect(() => {
    if (!hookOptions.immediate && reqTr === 0) return
    setLoading(true)
    fetcher<T>(url, reqOptions)
      .then(data => {
        if (data) {
          setData(data)
          hookOptions.onSuccess?.(data)
        } else {
          hookOptions.onError?.()
        }
      })
      .catch(error => {
        hookOptions.onError?.(error)
      })
      .finally(() => {
        hookOptions.onFinally?.()
        setLoading(false)
      })
  }, [reqTr])

  const run = () => {
    if (loading) return
    setReqTr(prev => prev + 1)
  }

  return { data, loading, run, setData }
}

export default useFetch