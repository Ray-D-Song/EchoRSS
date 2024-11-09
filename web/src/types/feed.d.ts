/**
 * type of items rendered in the sidebar
 */
interface Feed {
  id: number
  title: string

  /**
   * url of the feed
   */
  url: string

  /**
   * base64 encoded favicon
   */
  favicon: string

  /**
   * number of new articles
   */
  new: number

  /**
   * number of unread articles
   */
  unread: number
}
