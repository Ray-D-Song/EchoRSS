function htmlRewriter(origin: string, container: HTMLElement) {
  // rewrite all links to be relative to the origin
  container.querySelectorAll('a').forEach(a => {
    const href = a.getAttribute('href')
    if (!href) return
    if (href.startsWith('/')) {
      a.setAttribute('href', `${origin}${href}`)
    }
  })
  container.querySelectorAll('img').forEach(img => {
    const src = img.getAttribute('src')
    if (!src) return
    if (src.startsWith('/')) {
      img.setAttribute('src', `${origin}${src}`)
    }

    const srcset = img.getAttribute('srcset')
    if (srcset) {
      const newSrcset = srcset.split(',').map(item => {
        if (item.startsWith(' ')) item = item.slice(1)
        const bks = item.split(' ')
        const src = bks[0]
        if (src.startsWith('/')) {
          bks[0] = `${origin}${src}`
        }
        return bks.join(' ')
      }).join(',')
      img.setAttribute('srcset', newSrcset)
    }
  })
}

export default htmlRewriter