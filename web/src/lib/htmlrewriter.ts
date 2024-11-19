import { codeToHtml } from 'shiki/bundle/web'

async function rewriteLinks(origin: string, container: HTMLElement): Promise<void> {
  return new Promise((resolve) => {
    container.querySelectorAll('a').forEach(a => {
      const href = a.getAttribute('href')
      if (!href) return
      if (href.startsWith('/')) {
        a.setAttribute('href', `${origin}${href}`)
      }
    })
    resolve()
  })
}

async function rewriteImages(origin: string, container: HTMLElement): Promise<void> {
  return new Promise((resolve) => {
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
    resolve()
  })
}

async function rewriteCode(container: HTMLElement): Promise<void> {
  const promises = Array.from(container.querySelectorAll('pre code')).map(async code => {
    const divs = code.querySelectorAll('div');
    const content = Array.from(divs)
      .map(div => div.textContent || '')
      .join('\n');
    
    const html = await codeToHtml(content, { 
      lang: 'typescript',
      theme: 'nord'
    })
    
    const pre = code.parentElement
    if (pre) {
      pre.innerHTML = html
      console.log(html)
    }
  })
  
  return Promise.all(promises).then(() => {})
}

async function htmlRewriter(origin: string, container: HTMLElement) {
  await Promise.all([
    rewriteLinks(origin, container),
    rewriteImages(origin, container),
    rewriteCode(container)
  ])
}

export default htmlRewriter