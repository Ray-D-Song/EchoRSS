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
              refreshToken: JSON.parse(localStorage.getItem('user') ?? '{}')?.refreshToken
            })
          })
          if (refreshTokenRes.ok) {
            const userStr = localStorage.getItem('user')
            if (!userStr) return logout()
            localStorage.setItem('user', JSON.stringify({
              ...JSON.parse(userStr),
              ...(await refreshTokenRes.json())
            }))
            return await fetcher(url, options)
          } else logout()
        }
        break
      default:
        toast.error(await res.text())
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

function logout() {
  localStorage.removeItem('user')
  window.location.href = '/login'
}

export default fetcher