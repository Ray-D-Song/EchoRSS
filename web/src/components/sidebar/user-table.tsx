import { Table, TableBody, TableCell, TableHeader, TableHead, TableRow } from '@/components/ui/table'
import useFetch from '@/hooks/use-fetch'
import { memo, useContext, useState } from 'react'
import { Skeleton } from '@/components/ui/skeleton'
import { Button } from '@/components/ui/button'
import { Check, Edit, Eye, EyeOff, RefreshCcw, Trash, X } from 'lucide-react'
import { AppCtx } from '@/lib/ctx'
import { Input } from '@/components/ui/input'
import fetcher from '@/lib/fetcher'
import { toast } from 'react-hot-toast'

function UserTable() {

  const { user: currentUser } = useContext(AppCtx)

  const { data, loading, run: fetchUsers } = useFetch<UserListItem[]>('/users', {}, { immediate: true })
  const [showPassword, setShowPassword] = useState<{[key: string]: boolean}>({})

  const togglePassword = (userId: string) => {
    setShowPassword(prev => ({
      ...prev,
      [userId]: !prev[userId]
    }))
  }

  const [newUserForm, setNewUserForm] = useState<{username: string, password: string}>({username: '', password: ''})
  const [newUserBKVisible, setNewUserBKVisible] = useState(false)

  const handleNewUserSubmit = () => {
    fetcher('/users', {
      method: 'POST',
      body: JSON.stringify(newUserForm)
    }).then(() => {
      toast.success('User created successfully')
      setNewUserBKVisible(false)
      setNewUserForm({username: '', password: ''})
      fetchUsers()
    })
  }

  const handleDeleteUser = (userId: string) => {
    fetcher(`/users?id=${userId}`, {
      method: 'DELETE'
    }).then(() => {
      toast.success('User deleted successfully')
      fetchUsers()
    })
  }
  const handleRestoreUser = (userId: string) => {
    fetcher(`/users/restore?id=${userId}`, {
      method: 'PUT'
    }).then(() => {
      toast.success('User restored successfully')
      fetchUsers()
    })
  }

  return ( 
    <section>
      <div className='mb-2 flex justify-end'>
        <Button size="sm" onClick={() => setNewUserBKVisible(true)}>
          New User
        </Button>
      </div>
  <Table className='text-center'>
    <TableHeader>
      <TableRow>
        <TableHead className='text-center'>Username</TableHead>
        <TableHead className='text-center'>Password</TableHead>
        <TableHead className='text-center'>Actions</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableLoadingWrapper loading={loading}>
        {data?.map((user) => (
          <TableRow key={user.id} className={user.username === currentUser?.username ? 'text-yellow-600' : ''}>
            <TableCell className='p-1'>{user.username}</TableCell>
            <TableCell className='p-1'>
              <div className='flex items-center justify-center'>
                <span>{showPassword[user.id] ? user.password : '••••••'}</span>
                <Button 
                  variant="ghost" 
                  className='hover:text-yellow-600'
                  onClick={() => togglePassword(user.id)}
                >
                  {showPassword[user.id] ? 
                  <EyeOff className="h-4 w-4" /> : 
                  <Eye className="h-4 w-4" />
                  }
                </Button>
              </div>
            </TableCell>
            <TableCell className='flex gap-2 justify-center p-1'>
              <Button variant="ghost" className='p-2'>
                <Edit className="h-4 w-4" />
              </Button>
              {
                (user.username !== currentUser?.username && currentUser?.role === 'admin') && (
                  <>
                    <Button variant="ghost" className='p-2' onClick={() => user.deleted === 0 ? handleDeleteUser(user.id) : handleRestoreUser(user.id)}>
                      {
                        user.deleted === 0 ? 
                        <Trash className="h-4 w-4 text-red-900" /> :
                        <RefreshCcw className="h-4 w-4 text-green-800" />
                      }
                    </Button>
                  </>
                )
              }
            </TableCell>
          </TableRow>
        ))}
        {
          newUserBKVisible && (
            <TableRow>
              <TableCell className='p-1'>
                <Input
                  value={newUserForm.username}
                  onChange={(e) => setNewUserForm(prev => ({...prev, username: e.target.value}))}
                  autoFocus
                  className='h-8 w-fit border-b border-x-0 border-t-0 rounded-none focus-visible:ring-0 focus-visible:ring-offset-0 text-center'
                />
              </TableCell>
              <TableCell className='p-1'>
                <Input
                  value={newUserForm.password}
                  onChange={(e) => setNewUserForm(prev => ({...prev, password: e.target.value}))}
                  className='h-8 w-fit border-b border-x-0 border-t-0 rounded-none focus-visible:ring-0 focus-visible:ring-offset-0 text-center'
                />
              </TableCell>
              <TableCell className='flex gap-2 justify-center p-1'>
                <Button variant="ghost" onClick={handleNewUserSubmit} className='p-2'>
                  <Check className="h-4 w-4 text-green-500" />
                </Button>
                <Button variant="ghost" onClick={() => setNewUserBKVisible(false)} className='p-2'>
                  <X className="h-4 w-4 text-red-500" />
                </Button>
              </TableCell>
            </TableRow>
          )
        }
      </TableLoadingWrapper>
    </TableBody>
  </Table>
  </section>
  )
}

const TableSkeleton = memo(() => {
  // 5 rows skeleton
  return Array.from({ length: 5 }).map((_, i) => (
    <TableRow key={i}>
      <Skeleton className='h-4 w-full' />
    </TableRow>
  ))
})

const TableLoadingWrapper = ({ children, loading }: { children: React.ReactNode, loading: boolean }) => {
  return loading ? <TableSkeleton /> : children
}

export default UserTable