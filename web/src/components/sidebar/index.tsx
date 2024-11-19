import { Folder, LogOut, Rss, Settings, User } from 'lucide-react'
import { Button } from '@/components/ui/button'
import useFetch from '@/hooks/use-fetch'
import { useContext, useState } from 'react'
import { Skeleton } from '@/components/ui/skeleton'
import { AppCtx } from '@/lib/ctx'
import UserDialog from './user-dialog'
import NewFeed from './new-feed'
import { SidebarContent, SidebarGroup, SidebarHeader, Sidebar as SidebarUI, SidebarMenuItem, SidebarMenuSub, SidebarMenuSubItem, SidebarFooter, SidebarMenuButton, SidebarMenu, SidebarMenuSubButton } from '../ui/sidebar'
import fetcher, { logout } from '@/lib/fetcher'
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible'
import { ChevronRight } from 'lucide-react'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-react"
import { toast } from 'react-hot-toast'

interface SidebarProps {
  selectedFeed: Feed | null
  setSelectedFeed: (feed: Feed | null) => void
  feeds: Feed[] | null
  fetchFeeds: () => void
}

export function Sidebar({ selectedFeed, setSelectedFeed, feeds, fetchFeeds }: SidebarProps) {

  const { user } = useContext(AppCtx)
  const [open, setOpen] = useState(false)
  const [userDialogVisible, setUserDialogVisible] = useState(false)
  // fetch categories
  const { data: categories, loading: loadingCategories, run: fetchCategories } = useFetch<Category[]>('/category', {}, {
    immediate: true,
  })

  // refresh feeds
  const { run: refreshFeeds } = useFetch('/feeds/refresh', {
    method: 'POST'
  }, {
    immediate: true
  })

  function onNewFeedSuccess() {
    refreshFeeds()
    fetchCategories()
    fetchFeeds()
  }

  const handleRename = (category: Category) => {
    const newName = window.prompt('请输入新的分类名称:', category.name)
    if (newName && newName !== category.name) {
      fetcher(`/category/rename?originalName=${category.name}&newName=${newName}`, {
        method: 'PUT'
      })
      .then((data) => {
        if (data) {
          fetchCategories()
          toast.success('Rename success')
          return
        }
        toast.error('Rename failed')
      })
    }
  }

  const handleDelete = (category: Category) => {
    if (window.confirm(`确定要删除分类 "${category.name}" 吗?`)) {
      fetcher(`/category?name=${category.name}`, {
        method: 'DELETE'
      })
      .then((data) => {
        if (data) {
          fetchCategories()
          toast.success('Delete success')
          return
        }
        toast.error('Delete failed')
      })
    }
  }

  return (
    <SidebarUI>
      <SidebarHeader className='bg-background'>
        <div className='flex items-center gap-2'>
          <img src="/logo.svg" alt="logo" className='w-10 h-10' />
          <div className='flex flex-col ml-[-0.5rem]'>
            <span className='text-sm mb-[-0.2rem] font-semibold'>Echo</span>
            <span className='text-lg mt-[-0.2rem] font-bold'>RSS</span>
          </div>
        </div>
        <UserDialog open={userDialogVisible} onOpenChange={setUserDialogVisible} />
      </SidebarHeader>
      <SidebarContent className='bg-background h-[calc(100vh-18rem)]'>
        {loadingCategories ? (
          <Skeleton className="w-full h-10" />
        ) : (
          <SidebarGroup>
            <SidebarMenu>
              {categories?.map((category) => (
                <SidebarMenuItem key={category.id}>
                  <Collapsible defaultOpen className='group/collapsible'>
                    <CollapsibleTrigger asChild>
                      <SidebarMenuButton className="w-full">
                        <div className="flex items-center justify-between w-full">
                          <div className='flex items-center'>
                            <Folder className="mr-2 h-4 w-4" />
                            {category.name}
                          </div>
                          <div className="flex items-center">
                            <ChevronRight className="h-4 w-4 shrink-0 transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
                            <DropdownMenu>
                              <DropdownMenuTrigger asChild>
                                <Button variant="ghost" size="icon" className="h-8 w-8">
                                  <MoreHorizontal className="h-4 w-4" />
                                </Button>
                              </DropdownMenuTrigger>
                              <DropdownMenuContent align="end">
                                <DropdownMenuItem onClick={() => handleRename(category)}>Rename</DropdownMenuItem>
                                <DropdownMenuItem className="text-destructive" onClick={() => handleDelete(category)}>Delete</DropdownMenuItem>
                              </DropdownMenuContent>
                            </DropdownMenu>
                          </div>
                        </div>
                      </SidebarMenuButton>
                    </CollapsibleTrigger>
                    <CollapsibleContent>
                      <SidebarMenuSub>
                        {feeds
                          ?.filter((feed) => feed.categoryID === category.id)
                          .map((feed) => (
                            <SidebarMenuSubItem key={feed.id}>
                              <SidebarMenuSubButton
                                isActive={selectedFeed?.id === feed.id}
                                onClick={() => setSelectedFeed(feed)}
                                className='cursor-pointer'
                              >
                                <div className='flex items-center'>
                                  {
                                    feed.favicon ? <img src={feed.favicon} alt="favicon" className='w-4 h-4 mr-2' /> : <Rss className="mr-2 h-4 w-4" />
                                  }
                                  {feed.title}
                                </div>
                                <span className='text-[12px] opacity-70'>{feed.unreadCount > 0 ? feed.unreadCount : ''}</span>
                              </SidebarMenuSubButton>
                            </SidebarMenuSubItem>
                          ))}
                      </SidebarMenuSub>
                    </CollapsibleContent>
                  </Collapsible>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroup>
        )}
      </SidebarContent>
      <SidebarFooter className='bg-background'>
        <NewFeed open={open} setOpen={setOpen} categories={categories ?? []} onSuccess={onNewFeedSuccess} />
        <Button variant="ghost" className="w-full justify-start">
          <Settings className="mr-2 h-4 w-4" />
          Settings
        </Button>
        {user?.role === 'admin' && (
          <Button variant="ghost" className="w-full justify-start" onClick={() => setUserDialogVisible(true)}>
            <User className="mr-2 h-4 w-4" />
            User
          </Button>
        )}
        <Button variant="ghost" className="w-full justify-start" onClick={logout}>
          <LogOut className="mr-2 h-4 w-4" />
          Logout
        </Button>
      </SidebarFooter>
    </SidebarUI>
  )
}