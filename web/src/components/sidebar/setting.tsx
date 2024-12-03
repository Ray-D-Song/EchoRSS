import { Settings } from 'lucide-react'
import { Button } from '../ui/button'
import SettingDialog from '../setting-dialog'
import { useState } from 'react'

function Setting() {
  const [settingDialogVisible, setSettingDialogVisible] = useState(false)

  return (
    <>
      <SettingDialog
        visible={settingDialogVisible}
        onVisibleChange={setSettingDialogVisible}
      />
      <Button
        variant="ghost"
        className="w-full justify-start"
        onClick={() => setSettingDialogVisible(true)}
      >
        <Settings className="mr-2 h-4 w-4" />
        Settings
      </Button>
    </>
  )
}

export default Setting
