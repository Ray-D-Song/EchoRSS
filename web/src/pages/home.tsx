import { useState } from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Sidebar } from '@/components/sidebar'
import { Header } from '../components/header'
import { Article } from '../components/article'
import useView from '@/hooks/use-view'

// Mock data for RSS feeds
const mockFeeds = [
  { id: 1, name: 'Tech News', url: 'https://technews.com/rss' },
  { id: 2, name: 'World News', url: 'https://worldnews.com/rss' },
  { id: 3, name: 'Sports', url: 'https://sports.com/rss' },
  { id: 4, name: 'Science', url: 'https://science.com/rss' },
  { id: 5, name: 'Entertainment', url: 'https://entertainment.com/rss' },
]

// Mock data for articles
const mockArticles = [
  { id: 1, title: 'Latest Tech Innovations', summary: 'Exploring cutting-edge technologies shaping our future.', date: '2023-05-15' },
  { id: 2, title: 'Global Economic Trends', summary: 'Analysis of current economic patterns worldwide.', date: '2023-05-14' },
  { id: 3, title: 'Breakthrough in Quantum Computing', summary: 'Scientists achieve major milestone in quantum research.', date: '2023-05-13' },
  { id: 4, title: 'New Species Discovered', summary: 'Researchers find previously unknown species in the Amazon.', date: '2023-05-12' },
  { id: 5, title: 'Advancements in Renewable Energy', summary: 'Latest developments in sustainable power generation.', date: '2023-05-11' },
]

function Homepage() {
  useView()
  const [selectedFeed, setSelectedFeed] = useState(mockFeeds[0])


  return (
    <div className="flex flex-col h-screen">
      <Header />
      <main className="flex flex-1 overflow-hidden">
        <Sidebar selectedFeed={selectedFeed} setSelectedFeed={setSelectedFeed} />
        <section className="flex-1 p-6 overflow-y-auto max-w-sm border-r">
          <h2 className="text-2xl font-bold mb-6">{selectedFeed.name}</h2>
          <ScrollArea className="h-[calc(100vh-12rem)]">
            <div className="space-y-4">
              {mockArticles.map((article) => (
                <Card key={article.id}>
                  <CardHeader>
                    <CardTitle>{article.title}</CardTitle>
                    <CardDescription>{article.date}</CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p>{article.summary}</p>
                  </CardContent>
                </Card>
              ))}
            </div>
          </ScrollArea>
        </section>
        <section className="flex-1 p-6 overflow-y-auto">
          <Article />
        </section>
      </main>
    </div>
  )
}

export default Homepage