import { Rss } from 'lucide-react'
import { Button } from '@/components/ui/button'

interface Feed {
  id: number
  name: string
  url: string
}

const mockFeeds: Feed[] = [
  { id: 1, name: 'TechCrunch', url: 'https://techcrunch.com/rss' },
  { id: 2, name: 'The Verge', url: 'https://www.theverge.com/rss/index.xml' },
  { id: 3, name: 'Engadget', url: 'https://www.engadget.com/rss.xml' },
]

// Sidebar 组件
export function Sidebar({ selectedFeed, setSelectedFeed }: { selectedFeed: Feed, setSelectedFeed: (feed: Feed) => void }) {
  return (
    <aside className="w-64 border-r p-4 overflow-y-auto">
      <h2 className="text-lg font-semibold mb-4">RSS Feeds</h2>
      <ul className="space-y-2">
        {mockFeeds.map((feed) => (
          <li key={feed.id}>
            <Button
              variant={selectedFeed.id === feed.id ? 'secondary' : 'ghost'}
              className="w-full justify-start"
              onClick={() => setSelectedFeed(feed)}
            >
              <Rss className="mr-2 h-4 w-4" />
              {feed.name}
            </Button>
          </li>
        ))}
      </ul>
    </aside>
  )
}
