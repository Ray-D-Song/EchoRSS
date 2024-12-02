import fetcher from '@/lib/fetcher'
import { useEffect, useState } from 'react'
import { Readability } from '@mozilla/readability'
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from './ui/drawer'
import htmlRewriter from '@/lib/htmlrewriter'
import Markdown from 'react-markdown'
import turndownService from '@/lib/turndown'
import hljs from 'highlight.js'

interface ArticleProps {
  article: Article
  updateArticle: (article: Article) => void
}

function Article({ article, updateArticle }: ArticleProps) {
  const [remoteContent, setRemoteContent] = useState<string | null>(null)
  const [drawerVisible, setDrawerVisible] = useState(false)

  const [beautifiedContent, setBeautifiedContent] = useState<string | null>(null)
  useEffect(() => {
    const contentNeedBeautify = article.content.length > 0 ? article.content : article.description
    const docContentHtml = new DOMParser().parseFromString(contentNeedBeautify, 'text/html')
    htmlRewriter(new URL(article.link).origin, docContentHtml.documentElement).then(() => {
      setBeautifiedContent(turndownService.turndown(docContentHtml.documentElement.innerHTML))
    })
  }, [article.content, article.description])
  useEffect(() => {
    hljs.highlightAll()
  }, [beautifiedContent])

  useEffect(() => {
    if (article.read === 0) {
      fetcher(`/items/read?itemId=${article.id}`, {
        method: 'PUT',
      }).then((data) => {
        if(data) {
          updateArticle({
            ...article,
            read: 1,
          })
        }
      })
    }

    const handleClick = async (e: MouseEvent) => {
      const target = e.target as HTMLElement
      const link = target.closest('a')
      if (link) {
        e.preventDefault()
        let resDoc = await fetcher<string>(`/tools/fetch-remote-content?url=${link.href}`)
        if (!resDoc) {
          // open in new tab
          window.open(link.href, '_blank')
        }
        resDoc = resDoc?.replace('<head>', `<head><base href="${new URL(article.link).origin}" />`) ?? resDoc
        const doc = new DOMParser().parseFromString(resDoc ?? '', 'text/html')
        const docContent = new Readability(doc).parse()?.content
        const docContentHtml = new DOMParser().parseFromString(docContent ?? '', 'text/html')
        await htmlRewriter(new URL(article.link).origin, docContentHtml.documentElement)
        setRemoteContent(turndownService.turndown(docContentHtml.documentElement.innerHTML))
        setDrawerVisible(true)
      }
    }

    const articleContainer = document.querySelector('.prose') as HTMLElement
    if (articleContainer) {
      articleContainer.addEventListener('click', handleClick)
    }

    return () => {
      if (articleContainer) {
        articleContainer.removeEventListener('click', handleClick)
      }
    }
  }, [article.read, article.id])

  return <div className='prose dark:prose-invert'>
    <section>
      <Markdown>{beautifiedContent}</Markdown>
    </section>
    <Drawer open={drawerVisible} onOpenChange={setDrawerVisible}>
      <DrawerContent>
        <div className='max-h-[90vh] overflow-y-scroll'>
          <DrawerHeader>
            <DrawerTitle></DrawerTitle>
          </DrawerHeader>
          <div className='items-center flex justify-center'>
            <section className='prose dark:prose-invert'>
              <Markdown>{remoteContent}</Markdown>
            </section>
          </div>
        </div>
      </DrawerContent>
    </Drawer>
  </div>
}

export default Article