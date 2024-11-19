/**
 * type of items rendered in the sidebar
 */
interface Feed {
  id: string
  title: string
  link: string
  favicon: string
  description: string
  lastBuildDate: string
  categoryID: number
  unreadCount: number
  totalCount: number
  recentUpdateCount: number
  createdAt: string
}
