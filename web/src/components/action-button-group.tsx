import { CheckSquare, Edit, Loader2, RefreshCcw, Trash } from 'lucide-react'
import { Button } from '@/components/ui/button'
import useFetch from '@/hooks/use-fetch'

function ActionButtonGroup({ selectedFeed, fetchFeeds }: { selectedFeed: Feed | null, fetchFeeds: () => void }) {
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
    <>
      <Button variant="secondary" className='w-8 h-8 p-0 rounded-r-none border-r'>
        <RefreshCcw className='w-4 h-4' />
      </Button>
      <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none rounded-r-none border-r' onClick={markAllAsRead}>
        <CheckSquare className='w-4 h-4' />
      </Button>
      <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none rounded-r-none border-r'>
        <Edit className='w-4 h-4' />
      </Button>
      <Button variant="secondary" className='w-8 h-8 p-0 rounded-l-none' onClick={deleteFeed} disabled={deletingFeed}>
        {
          deletingFeed ? <Loader2 className='w-4 h-4 animate-spin' /> : <Trash className='w-4 h-4' />
        }
      </Button>
    </>
  )
}

export default ActionButtonGroup
