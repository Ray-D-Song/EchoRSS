import { DialogHeader, Dialog, DialogContent, DialogPortal, DialogTitle } from './ui/dialog'
import useFetch from '@/hooks/use-fetch'
import { useState } from 'react'
import { Input } from './ui/input'
import { Button } from './ui/button'

interface UserConfig {
  API_KEY: string
  API_ENDPOINT: string
  TARGET_LANGUAGE: string
}

interface SettingDialogProps {
  visible: boolean
  onVisibleChange: (visible: boolean) => void
}

function SettingDialog({ visible, onVisibleChange }: SettingDialogProps) {
  const [aiSetting, setAiSetting] = useState<UserConfig>({
    API_KEY: '',
    API_ENDPOINT: '',
    TARGET_LANGUAGE: ''
  })

  useFetch<UserConfig>(
    '/user/config',
    {
      method: 'GET',
    },
    {
      immediate: true,
      onSuccess: (data) => {
        setAiSetting({
          API_KEY: data.API_KEY,
          API_ENDPOINT: data.API_ENDPOINT,
          TARGET_LANGUAGE: data.TARGET_LANGUAGE
        })
      },
    },
  )

  const { run: updateAiSetting } = useFetch('/user/config', {
    method: 'PUT',
  })

  const handleSubmit = () => {
    updateAiSetting()
  }

  return (
    <Dialog
      open={visible}
      onOpenChange={onVisibleChange}
    >
      <DialogPortal>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Setting</DialogTitle>
          </DialogHeader>

          <div className="space-y-4 mt-4">
          <div className="space-y-2">
            <label className="text-sm font-medium">OpenAI API Key</label>
            <Input
              value={aiSetting.API_KEY}
              onChange={(e) => setAiSetting(prev => ({
                ...prev,
                API_KEY: e.target.value
              }))}
              placeholder="sk-..."
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">API Endpoint</label>
            <Input
              value={aiSetting.API_ENDPOINT}
              onChange={(e) => setAiSetting(prev => ({
                ...prev,
                API_ENDPOINT: e.target.value
              }))}
              placeholder="https://api.openai.com/v1"
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">Target Language</label>
            <Input
              value={aiSetting.TARGET_LANGUAGE}
              onChange={(e) => setAiSetting(prev => ({
                ...prev,
                TARGET_LANGUAGE: e.target.value
              }))}
            />
          </div>

          <Button className="w-full" onClick={handleSubmit}>
            Save
            </Button>
          </div>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  )
}

export default SettingDialog
