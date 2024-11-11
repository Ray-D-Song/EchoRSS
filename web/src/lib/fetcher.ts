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
        {
          if (url === '/auth/refresh-token') return logout()
          const refreshTokenRes = await fetch('/api/auth/refresh-token', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              token: JSON.parse(localStorage.getItem('user') ?? '{}')?.refreshToken
            })
          })
          if (refreshTokenRes.ok) {
            localStorage.setItem('user', JSON.stringify(await refreshTokenRes.json()))
          } else logout()
        }
        break
      default:
        toast.error(await res.text())
    }
    return null
  }
  return await res.json() as T
}

function logout() {
  localStorage.removeItem('user')
  window.location.href = '/login'
}

export default fetcher