import { useEffect, useState } from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Sidebar } from '@/components/sidebar/index'
import { Header } from '../components/header'
import Article from '../components/article'
import useView from '@/hooks/use-view'
import useFetch from '@/hooks/use-fetch'
import { SidebarProvider } from '@/components/ui/sidebar'
import { Edit, Loader2, RefreshCcw, Trash } from 'lucide-react'
import { Button } from '@/components/ui/button'

function Homepage() {
  useView()
  const [selectedFeed, setSelectedFeed] = useState<Feed | null>(null)
  const [selectedArticle, setSelectedArticle] = useState<Article | null>(null)
  const { data: articles, run: refreshItems, setData: setArticles } = useFetch<Article[]>(`/items?feedId=${selectedFeed?.id}`, {}, {
    immediate: false,
    onSuccess: (data) => {
      if (data.length > 0) {
        setSelectedArticle(data[0])
      }
    }
  })

  useEffect(() => {
    if (selectedFeed) {
      refreshItems()
    }
  }, [selectedFeed])

  // fetch feeds
  const { data: feeds, run: fetchFeeds, setData: setFeeds } = useFetch<Feed[]>('/feeds', {}, {
    immediate: true,
  })
  function updateArticles(changedArticle: Article) {
    const newArticles = articles?.map((a: Article) => a.id === changedArticle.id ? changedArticle : a) ?? []
    setArticles(newArticles)
    if (changedArticle.read === 1) {
      setFeeds(feeds?.map((f: Feed) => f.id === selectedFeed?.id ? {
        ...f,
        unreadCount: f.unreadCount - 1,
      } : f) ?? [])
    }
  }
  // delete feed
  const { run: deleteFeed, loading: deletingFeed } = useFetch(`/feeds?feedID=${selectedFeed?.id}`, {
    method: 'DELETE',
  }, {
    immediate: false,
    onSuccess: () => {
      fetchFeeds()
    }
  })

  return (
    <div className="flex flex-col h-screen">
      <Header />
      <main className="flex flex-1 overflow-hidden">
        <SidebarProvider>
        <Sidebar
          selectedFeed={selectedFeed}
          setSelectedFeed={setSelectedFeed}
          feeds={feeds}
          fetchFeeds={fetchFeeds}
        />
        <section className="flex-1 p-6 overflow-y-auto max-w-[18rem] border-r">
          <div className='flex flex-col justify-between gap-2 mb-6'>
            <h2 className="text-2xl font-bold">{selectedFeed?.title}</h2>
            <div className='flex items-center' style={{
              visibility: selectedFeed ? 'visible' : 'hidden',
            }}>
              <Button variant="secondary" className='w-8 h-8 p-0 rounded-r-none border-r'>
                <RefreshCcw className='w-4 h-4' />
              </Button>
              <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none rounded-r-none border-r'>
                <Edit className='w-4 h-4' />
              </Button>
              <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none' onClick={deleteFeed} disabled={deletingFeed}>
                {
                  deletingFeed ? <Loader2 className='w-4 h-4 animate-spin' /> : <Trash className='w-4 h-4' />
                }
              </Button>
            </div>
          </div>
          <ScrollArea className="h-[calc(100vh-12rem)]">
            <div className="space-y-4">
              {articles?.map((article) => (
                <Card key={article.id} onClick={() => setSelectedArticle(article)} className={`${article.read ? 'opacity-60' : ''}`}>
                  <CardHeader>
                    <CardTitle className="text-lg font-semibold leading-6">{article.title}</CardTitle>
                    <CardDescription>{article.pubDate}</CardDescription>
                  </CardHeader>
                  {
                    article.description.length > 0 && (
                      <CardContent>
                        <div dangerouslySetInnerHTML={{ __html: article.description.slice(0, 110) + '...' }} />
                      </CardContent>
                    )
                  }
                </Card>
              ))}
            </div>
          </ScrollArea>
        </section>
        <section className="flex-1 p-6 overflow-y-auto">
          {selectedArticle && <Article article={selectedArticle} updateArticle={updateArticles} />}
        </section>
        </SidebarProvider>
      </main>
    </div>
  )
}

export default Homepage