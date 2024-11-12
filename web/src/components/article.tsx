import fetcher from '@/lib/fetcher'
import { useEffect, useState } from 'react'
import { Readability } from '@mozilla/readability'
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from './ui/drawer'
import htmlRewriter from '@/lib/htmlrewriter'

interface ArticleProps {
  article: Article
  updateArticle: (article: Article) => void
}

function Article({ article, updateArticle }: ArticleProps) {
  const [remoteContent, setRemoteContent] = useState<ReturnType<Readability<string>['parse']> | null>(null)
  const [drawerVisible, setDrawerVisible] = useState(false)
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
        const resDoc = await fetcher<string>(`/tools/fetch-remote-content?url=${link.href}`)
        if (!resDoc) {
          // open in new tab
          window.open(link.href, '_blank')
        }
        const doc = new DOMParser().parseFromString(resDoc ?? '', 'text/html')
        htmlRewriter(new URL(article.link).origin, doc.documentElement)
        const docContent = new Readability(doc).parse()
        setRemoteContent(docContent)
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
  }, [article.read])

  return <div className='prose dark:prose-invert'>
    <section dangerouslySetInnerHTML={{ __html: article.content.length > 0 ? article.content : article.description }} />
    <Drawer open={drawerVisible} onOpenChange={setDrawerVisible}>
      <DrawerContent>
        <div className='max-h-[90vh] overflow-y-scroll'>
          <DrawerHeader>
            <DrawerTitle>{remoteContent?.title}</DrawerTitle>
          </DrawerHeader>
          <div className='items-center flex justify-center'>
            <section className='prose dark:prose-invert' dangerouslySetInnerHTML={{ __html: remoteContent?.content ?? '' }} />
          </div>
        </div>
      </DrawerContent>
    </Drawer>
  </div>
}

export default Article