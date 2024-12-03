import { LanguagesIcon } from 'lucide-react'
import { Button } from './ui/button'
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from './ui/drawer'
import Markdown from 'react-markdown'
import useTranslate from '@/hooks/use-trans'
import { useEffect } from 'react'
import hljs from 'highlight.js'

interface ArticleDrawerProps {
  visible: boolean
  setVisible: (visible: boolean) => void
  content: string
  url: string
}

function ArticleDrawer({ visible, setVisible, content, url }: ArticleDrawerProps) {
  const { translate, result } = useTranslate({
    content: content,
    url: url,
  })

  useEffect(() => {
    hljs.highlightAll()
  }, [result])
  return (
    <Drawer open={visible} onOpenChange={setVisible}>
    <DrawerContent>
      <div className='h-[90vh] overflow-y-scroll'>
        <DrawerHeader>
          <DrawerTitle>
            <Button 
              size="icon" 
              variant="outline"
              className="fixed bottom-6 right-6 z-10 shadow-lg hover:shadow-xl transition-shadow"
              onClick={() => {
                translate()
              }}
            >
              <LanguagesIcon className='w-5 h-5' />
            </Button>
          </DrawerTitle>
        </DrawerHeader>
        <div className='items-center flex justify-center'>
          <section className='prose dark:prose-invert'>
            <Markdown>{result}</Markdown>
          </section>
        </div>
      </div>
    </DrawerContent>
  </Drawer>
  )
}

export default ArticleDrawer