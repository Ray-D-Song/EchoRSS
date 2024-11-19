import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import UserTable from './user-table'

function UserDialog({ open, onOpenChange }: { open: boolean, onOpenChange: (open: boolean) => void }) {
  return <Dialog open={open} onOpenChange={onOpenChange}>
    <DialogContent className='w-[80vw] max-w-xl'>
      <DialogHeader>
        <DialogTitle>
          Users
        </DialogTitle>
      </DialogHeader>
      <UserTable />
    </DialogContent>
  </Dialog>
}

export default UserDialog