import { useEffect, useState } from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Sidebar } from '@/components/sidebar'
import { Header } from '../components/header'
import Article from '../components/article'
import useView from '@/hooks/use-view'
import useFetch from '@/hooks/use-fetch'

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

  function updateArticles(article: Article) {
    setArticles(articles => articles?.map(a => a.id === article.id ? article : a) ?? [])
  }

  return (
    <div className="flex flex-col h-screen">
      <Header />
      <main className="flex flex-1 overflow-hidden">
        <Sidebar selectedFeed={selectedFeed} setSelectedFeed={setSelectedFeed} />
        <section className="flex-1 p-6 overflow-y-auto max-w-sm border-r">
          <h2 className="text-2xl font-bold mb-6">{selectedFeed?.title}</h2>
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
      </main>
    </div>
  )
}

export default Homepage