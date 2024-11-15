import { LogOut, Plus, Rss, Settings, User } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Popover, PopoverContent, PopoverTrigger } from '@radix-ui/react-popover'
import { Input } from './ui/input'
import useFetch from '@/hooks/use-fetch'
import { useContext, useState } from 'react'
import { Skeleton } from './ui/skeleton'
import { ScrollArea } from './ui/scroll-area'
import { AppCtx } from '@/lib/ctx'
import UserDialog from './user-dialog'

export function Sidebar({ selectedFeed, setSelectedFeed }: { selectedFeed: Feed | null, setSelectedFeed: (feed: Feed | null) => void }) {

  const { user } = useContext(AppCtx)
  const [newFeed, setNewFeed] = useState('')
  const [open, setOpen] = useState(false)
  const [userDialogVisible, setUserDialogVisible] = useState(false)
  const { data: feeds, loading: loadingFeeds, run: refreshFeeds } = useFetch<Feed[]>('/feeds', {}, {
    immediate: true,
    onSuccess: (data) => {
      if (!selectedFeed && data.length > 0) {
        setSelectedFeed(data[0])
      }
    }
  })
  useFetch('/feeds/refresh', {
    method: 'POST'
  }, {
    immediate: true
  })
  const { run } = useFetch<Feed[]>('/feeds', 
  {
    method: 'POST',
      body: JSON.stringify({
        url: newFeed
      })
    },
    {
      onFinally: () => {
        setOpen(false)
        setNewFeed('')
      }
    }
  )

  return (
    <aside className="w-64 border-r overflow-y-auto">
      <h2 className="text-lg font-semibold m-2 mb-0">RSS Feeds</h2>
      <ScrollArea className="h-[calc(100vh-16rem)]">
        <ul className="space-y-1 p-2">
          {
          loadingFeeds ? <Skeleton className="w-full h-10" /> :
          feeds?.map((feed) => (
          <li key={feed.id}>
            <Button
              variant={selectedFeed?.id === feed.id ? 'secondary' : 'ghost'}
              className="w-full justify-start"
              onClick={() => setSelectedFeed(feed)}
            >
              {
                feed.favicon ? <img src={`data:image/png;base64,${feed.favicon}`} alt="favicon" className="w-4 h-4 mr-2 rounded-sm" /> :
                <Rss className="mr-2 h-4 w-4" />
              }
              {feed.title}
              {feed.unreadCount > 0 && <span className="text-xs text-gray-500 ml-auto">({feed.unreadCount})</span>}
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
      </ScrollArea>
      <div className='border-t' />
      <ul className="space-y-1 p-2">
        <li>
          <Button variant="ghost" className="w-full justify-start">
            <Settings className="mr-2 h-4 w-4" />
            Settings
          </Button>
        </li>
        {
          user?.role === 'admin' && (
            <li>
              <UserDialog open={userDialogVisible} onOpenChange={setUserDialogVisible} />
              <Button variant="ghost" className="w-full justify-start" onClick={() => setUserDialogVisible(true)}>
                <User className="mr-2 h-4 w-4" />
                User
              </Button>
            </li>
          )
        }
        <li>
          <Button variant="ghost" className="w-full justify-start">
            <LogOut className="mr-2 h-4 w-4" />
            Logout
          </Button>
        </li>
      </ul>
    </aside>
  )
}
