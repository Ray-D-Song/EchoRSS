/**
 * type of items rendered in the article card
 */
interface Article {
  title: string
  description: string
  link: string
  pubDate: string

  /**
   * whether the article has been read
   */
  read: boolean

  /**
   * whether the article has been archived locally
   */
  archived: boolean

  /**
   * @todo whether the article has been favorited locally
   */
  favorite: boolean
}
