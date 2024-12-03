import fetcher from '@/lib/fetcher'
import { useEffect, useState } from 'react'
import { Readability } from '@mozilla/readability'
import htmlRewriter from '@/lib/htmlrewriter'
import Markdown from 'react-markdown'
import turndownService from '@/lib/turndown'
import hljs from 'highlight.js'
import { Button } from './ui/button'
import { BookmarkIcon, LanguagesIcon } from 'lucide-react'
import useTranslate from '@/hooks/use-trans'
import toast from 'react-hot-toast'
import useFetch from '@/hooks/use-fetch'

interface ArticleProps {
  article: Article
  updateArticle: (article: Article, actions: 'read' | 'bookmark') => void
  setRemoteContent: (content: {
    content: string
    url: string
  }) => void
  setDrawerVisible: (visible: boolean) => void
}

function Article({ article, updateArticle, setRemoteContent, setDrawerVisible }: ArticleProps) {
  const [beautifiedContent, setBeautifiedContent] = useState<string | null>(null)
  useEffect(() => {
    const contentNeedBeautify = article.content.length > 0 ? article.content : article.description
    const docContentHtml = new DOMParser().parseFromString(contentNeedBeautify, 'text/html')
    htmlRewriter(new URL(article.link).origin, docContentHtml.documentElement).then(() => {
      setBeautifiedContent(turndownService.turndown(docContentHtml.documentElement.innerHTML))
    })
  }, [article.content, article.description])


  useEffect(() => {
    if (article.read === 0) {
      fetcher(`/items/read?itemId=${article.id}`, {
        method: 'PUT',
      }).then((data) => {
        if(data) {
          updateArticle({
            ...article,
            read: 1,
          }, 'read')
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
        setRemoteContent({
          content: turndownService.turndown(docContentHtml.documentElement.innerHTML),
          url: link.href,
        })
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

  const { translate, result } = useTranslate({
    url: article.id,
    content: beautifiedContent ?? '',
  })

  useEffect(() => {
    hljs.highlightAll()
  }, [result])

  const { run: toggleBookmark } = useFetch(`/bookmark?itemId=${article.id}`, {
    method: 'PUT',
  }, {
    immediate: false,
    onSuccess: () => {
      updateArticle({
        ...article,
        bookmark: article.bookmark === 0 ? 1 : 0,
      }, 'bookmark')
    },
    onError: () => {
      toast.error(`Failed to ${article.bookmark === 0 ? 'bookmark' : 'remove bookmark'}`)
    }
  })

  return <section className='prose dark:prose-invert'>
      <div className='fixed bottom-6 right-6 z-10 shadow-lg hover:shadow-xl transition-shadow border'>
        <Button 
          size="icon" 
          variant="ghost"
          className='rounded-none'
          onClick={() => {
            translate()
          }}
        >
          <LanguagesIcon className='w-5 h-5' />
        </Button>
        <Button size="icon" variant="ghost" className='rounded-none' onClick={toggleBookmark}>
          <BookmarkIcon className={`w-5 h-5 ${article.bookmark === 1 ? 'text-yellow-500' : ''}`} />
        </Button>
      </div>
      <Markdown>{result}</Markdown>
    </section>
}

export default Article