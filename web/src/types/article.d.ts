/**
 * type of items rendered in the article card
 */
interface Article {
  id: string
  title: string
  content: string
  description: string
  link: string
  pubDate: string

  /**
   * whether the article has been read
   */
  read: boolean

  /**
   * @todo whether the article has been favorited locally
   */
  favorite: boolean
}
