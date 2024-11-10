import toast from 'react-hot-toast'

async function fetcher<T>(url: string, options?: RequestInit) {
  url = `/api${url}`
  const res = await fetch(url, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    ...options,
  })
  if (!res.ok) {
    switch (res.status) {
      case 401:
        window.location.href = '/login'
        break
      default:
        toast.error(await res.text())
    }
    return null
  }
  return await res.json() as T
}

export default fetcher