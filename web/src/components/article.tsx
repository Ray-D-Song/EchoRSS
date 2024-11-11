function Article(props: Article) {
  return <div className='prose dark:prose-invert'>
    <section dangerouslySetInnerHTML={{ __html: props.content.length > 0 ? props.content : props.description }} />
  </div>
}

export default Article