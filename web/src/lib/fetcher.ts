import { redirect } from '@/router'
import toast from 'react-hot-toast'

async function fetcher<T>(url: string, options?: RequestInit) {
  url = `/api${url}`
  const res = await fetch(url, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${JSON.parse(localStorage.getItem('user') ?? '{}')?.token}`
    },
    ...options,
  })
  if (!res.ok) {
    switch (res.status) {
      case 401:
        logout()
        break
      default: {
        const errMsg = await res.json()
        toast.error(errMsg.error)
      }
    }
    return null
  }
  if (res.headers.get('Content-Type') === 'text/html') {
    return await res.text() as T
  }
  if (res.headers.get('Content-Type') === 'application/json') {
    return await res.json() as T
  }
  return await res.text() as T
}

export function logout() {
  localStorage.removeItem('user')
  if (window.location.hash !== '#/login') {
    redirect('/login')
  }
}

export default fetcher