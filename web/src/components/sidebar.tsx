import { Plus, Rss } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Popover, PopoverContent, PopoverTrigger } from '@radix-ui/react-popover'
import { Input } from './ui/input'
import useFetch from '@/hooks/use-fetch'
import { useState } from 'react'

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

export function Sidebar({ selectedFeed, setSelectedFeed }: { selectedFeed: Feed, setSelectedFeed: (feed: Feed) => void }) {

  const [newFeed, setNewFeed] = useState('')
  const [open, setOpen] = useState(false)
  const { data, loading, run } = useFetch<Feed[]>('/feeds', 
  {
    method: 'POST',
    body: newFeed
    },
    {
      onFinally: () => {
        setOpen(false)
        setNewFeed('')
      }
    }
  )

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
        <li>
          <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
              <Button variant="secondary" className="w-full justify-start bg-[#FACC14]/10 hover:bg-[#FACC14]/20">
                <Plus className="mr-2 h-4 w-4" />
                Add Feed
              </Button>
            </PopoverTrigger>
            <PopoverContent className='z-50'>
              <Input placeholder="New Feed URL, press Enter to add" size={40} className='ml-4 mt-2' value={newFeed} onChange={(e) => setNewFeed(e.target.value)} onKeyDown={async (e) => {
                if (e.key === 'Enter' && newFeed.length > 0) {
                  run()
                }
              }} />
            </PopoverContent>
          </Popover>
        </li>
      </ul>
    </aside>
  )
}
