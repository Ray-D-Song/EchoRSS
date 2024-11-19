import { Popover, PopoverContent, PopoverTrigger } from '@radix-ui/react-popover'
import { Button } from '../ui/button'
import { Input } from '../ui/input'
import { Rss } from 'lucide-react'
import { useState } from 'react'
import useFetch from '@/hooks/use-fetch'

interface NewFeedProps {
  open: boolean
  setOpen: (open: boolean) => void
  categories: Category[]
  onSuccess: () => void
}

function NewFeed({ open, setOpen, categories, onSuccess }: NewFeedProps) {
  const [showDropdown, setShowDropdown] = useState(false)
  const [form, setForm] = useState({
    url: '',
    category: ''
  })

  const { run } = useFetch('/feeds', 
    {
      method: 'POST',
      body: JSON.stringify(form)
    }, {
      immediate: false,
      onSuccess: () => {
        setOpen(false)
        onSuccess()
      }
    }
  )

  // 过滤分类列表
  const filteredCategories = categories?.filter(category => 
    category.name.toLowerCase().includes(form.category.toLowerCase())
  )

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button variant="secondary" className="w-full justify-start bg-[#FACC14]/10 hover:bg-[#FACC14]/20">
          <Rss className="mr-2 h-4 w-4" />
          New Feed
        </Button>
      </PopoverTrigger>
      <PopoverContent className='z-50'>
        <div className='flex flex-col gap-2 justify-center items-center w-80 ml-4 border-solid border p-4 bg-background rounded-lg mt-2'>
          <Input 
            placeholder="New Feed URL" 
            value={form.url} 
            onChange={(e) => setForm({ ...form, url: e.target.value })} 
          />
          
          <div className="relative w-full">
            <Input
              placeholder="Select or input new category"
              value={form.category}
              onChange={(e) => setForm({ ...form, category: e.target.value })}
              onFocus={() => setShowDropdown(true)}
            />
            
            {showDropdown && (
              <div className="absolute w-full mt-1 max-h-[200px] overflow-y-auto border rounded-md bg-background shadow-lg">
                {filteredCategories?.map((category) => (
                  <div
                    key={category.id}
                    className="px-4 py-2 hover:bg-accent cursor-pointer"
                    onClick={() => {
                      setForm({ ...form, category: category.name })
                      setShowDropdown(false)
                    }}
                  >
                    {category.name}
                  </div>
                ))}
              </div>
            )}
          </div>

          <Button variant="default" className="w-full" onClick={run}>
            Save
          </Button>
        </div>
      </PopoverContent>
    </Popover>
  )
}

export default NewFeed