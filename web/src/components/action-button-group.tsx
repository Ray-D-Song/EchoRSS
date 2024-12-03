import { Bookmark, CheckSquare, Edit, Loader2, RefreshCcw, Trash } from 'lucide-react'
import { Button } from '@/components/ui/button'
import useFetch from '@/hooks/use-fetch'
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip'

interface ActionButtonGroupProps {
  selectedFeed: Feed | null
  fetchFeeds: () => void
  filter: string
  setFilter: (filter: string) => void
}

function ActionButtonGroup({ selectedFeed, fetchFeeds, filter, setFilter }: ActionButtonGroupProps) {
  // delete feed
  const { run: deleteFeed, loading: deletingFeed } = useFetch(`/feeds?feedID=${selectedFeed?.id}`, {
    method: 'DELETE',
  }, {
    immediate: false,
    onSuccess: () => {
      fetchFeeds()
    }
  })
  // mark all as read
  const { run: markAllAsRead } = useFetch(`/feeds/all-read?feedID=${selectedFeed?.id}`, {
    method: 'PUT',
  }, {
    immediate: false,
    onSuccess: () => {
      fetchFeeds()
    }
  })
  return (
    <TooltipProvider>
      <>
        <Tooltip>
          <TooltipTrigger asChild>
            <Button variant="secondary" className='w-8 h-8 p-0 rounded-r-none border-r'>
              <RefreshCcw className='w-4 h-4' />
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Refresh</p>
          </TooltipContent>
        </Tooltip>

        <Tooltip>
          <TooltipTrigger asChild>
            <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none rounded-r-none border-r' onClick={markAllAsRead}>
              <CheckSquare className='w-4 h-4' />
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Mark all as read</p>
          </TooltipContent>
        </Tooltip>

        <Tooltip>
          <TooltipTrigger asChild>
            <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none rounded-r-none border-r'>
              <Edit className='w-4 h-4' />
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Edit</p>
          </TooltipContent>
        </Tooltip>

        <Tooltip>
          <TooltipTrigger asChild>
            <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none rounded-r-none border-r' onClick={deleteFeed} disabled={deletingFeed}>
              {deletingFeed ? <Loader2 className='w-4 h-4 animate-spin' /> : <Trash className='w-4 h-4' />}
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Delete</p>
          </TooltipContent>
        </Tooltip>

        <Tooltip>
          <TooltipTrigger asChild>
            <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none' onClick={() => setFilter(filter === 'all' ? 'bookmark' : 'all')}>
              <Bookmark className={`w-4 h-4 ${filter === 'bookmark' ? 'text-yellow-500' : ''}`} />
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Show bookmark</p>
          </TooltipContent>
        </Tooltip>
      </>
    </TooltipProvider>
  )
}

export default ActionButtonGroup
