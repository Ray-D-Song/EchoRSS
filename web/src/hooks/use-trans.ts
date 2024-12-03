import { useState, useEffect } from 'react';

import useFetch from './use-fetch';

interface UseTranslateReturn {
  translate: () => void;
  result: string;
  toggle: () => void;
}

interface Options {
  // if translate feed item, url will be item id
  url: string
  content: string
  onSuccess?: (translated: string) => void
  onError?: (original: string) => void
  onFinally?: () => void
}

function useTranslate(options: Options): UseTranslateReturn {
  const [contentMap, setContentMap] = useState<Map<string, string>>(new Map([
    ['raw', options.content],
    ['translated', '']
  ]))
  const [current, setCurrent] = useState<'raw' | 'translated'>('raw')
  
  useEffect(() => {
    setContentMap(new Map([
      ['raw', options.content],
      ['translated', contentMap.get('translated') || '']
    ]))
  }, [options.content])

  const toggle = () => setCurrent(current === 'raw' ? 'translated' : 'raw')

  const formData = new URLSearchParams()
  formData.append('url', options.url)
  formData.append('content', options.content)
  
  const { run, loading } = useFetch('/translate', {
    method: 'POST',
    body: formData,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
  }, {
    immediate: false,
    onSuccess: (data: string) => {
      setContentMap(prev => new Map([
        ['raw', prev.get('raw') || ''],
        ['translated', data]
      ]))
      setCurrent('translated')
    },
    onError() {
      options.onError?.(options.content)
    },
    onFinally: options.onFinally
  })

  const translate = () => {
    const translated = contentMap.get('translated')
    if (translated && typeof translated === 'string' && translated.length > 0) {
      setCurrent(current === 'raw' ? 'translated' : 'raw')
    } else {
      run()
    }
  }

  return {
    translate,
    result: loading ? 'loading...' : contentMap.get(current) ?? '',
    toggle
  }
}

export default useTranslate;