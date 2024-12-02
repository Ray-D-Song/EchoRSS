import { useContext, useState } from 'react'
import { Button } from '@/components/ui/button'
import { Check, Edit, Eye, EyeOff, RefreshCcw, Trash, X } from 'lucide-react'
import { AppCtx } from '@/lib/ctx'
import { Input } from '@/components/ui/input'
import fetcher from '@/lib/fetcher'
import { toast } from 'react-hot-toast'
import useFetch from '@/hooks/use-fetch'

function UserTable() {
  const { user: currentUser } = useContext(AppCtx)
  const { data, loading, run: fetchUsers } = useFetch<UserListItem[]>('/users', {}, { immediate: true })
  const [showPassword, setShowPassword] = useState<{[key: string]: boolean}>({})
  const [editingUser, setEditingUser] = useState<{id: string, username: string, password: string} | null>(null)
  const [newUserForm, setNewUserForm] = useState<{username: string, password: string}>({username: '', password: ''})
  const [newUserBKVisible, setNewUserBKVisible] = useState(false)

  const togglePassword = (userId: string) => {
    setShowPassword(prev => ({
      ...prev,
      [userId]: !prev[userId]
    }))
  }

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

  const handleUpdateUser = (userId: string, username: string, password: string) => {
    fetcher(`/users`, {
      method: 'PUT',
      body: JSON.stringify({id: userId, username, password})
    }).then(() => {
      toast.success('User updated successfully')
      fetchUsers()
    })
  }

  const handleEditClick = (user: UserListItem) => {
    setEditingUser({
      id: user.id,
      username: user.username,
      password: user.password
    })
  }

  const handleEditSubmit = () => {
    if(!editingUser) return
    handleUpdateUser(editingUser.id, editingUser.username, editingUser.password)
    setEditingUser(null)
  }

  return (
    <section>
      <div className='mb-2 flex'>
        <Button size="sm" onClick={() => setNewUserBKVisible(true)}>
          New User
        </Button>
      </div>

        {loading ? (
          <div className='col-span-3 text-center p-4'>Loading...</div>
        ) : (
          <>
            {data?.map((user) => (
              <div key={user.id} className={`flex justify-between md:text-sm items-center ${user.username === currentUser?.username ? 'text-yellow-600' : ''}`}>
                <div>
                  {editingUser?.id === user.id ? (
                    <Input
                      value={editingUser.username}
                      onChange={(e) => setEditingUser(prev => ({...prev!, username: e.target.value}))}
                    />
                  ) : (
                    <span className='px-2'>{user.username}</span>
                  )}
                </div>
                <div>
                  <div className='flex items-center justify-center'>
                    {editingUser?.id === user.id ? (
                      <Input
                        value={editingUser.password}
                        onChange={(e) => setEditingUser(prev => ({...prev!, password: e.target.value}))}
                      />
                    ) : (
                      <div className='flex items-center gap-2'>
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
                    )}
                  </div>
                </div>
                <div className='flex gap-2 justify-center'>
                  {editingUser?.id === user.id ? (
                    <>
                      <Button variant="ghost" onClick={handleEditSubmit} className='p-2'>
                        <Check className="h-4 w-4 text-green-500" />
                      </Button>
                      <Button variant="ghost" onClick={() => setEditingUser(null)} className='p-2'>
                        <X className="h-4 w-4 text-red-500" />
                      </Button>
                    </>
                  ) : (
                    <>
                      <Button variant="ghost" className='p-2' onClick={() => handleEditClick(user)}>
                        <Edit className="h-4 w-4" />
                      </Button>
                      {(user.username !== currentUser?.username && currentUser?.role === 'admin') && (
                        <Button variant="ghost" className='p-2' onClick={() => user.deleted === 0 ? handleDeleteUser(user.id) : handleRestoreUser(user.id)}>
                          {user.deleted === 0 ? 
                            <Trash className="h-4 w-4 text-red-900" /> :
                            <RefreshCcw className="h-4 w-4 text-green-800" />
                          }
                        </Button>
                      )}
                    </>
                  )}
                </div>
              </div>
            ))}

            {newUserBKVisible && (
              <div className='flex justify-between items-center'>
                <div>
                  <Input
                    value={newUserForm.username}
                    onChange={(e) => setNewUserForm(prev => ({...prev, username: e.target.value}))}
                    autoFocus
                  />
                </div>
                <div>
                  <Input
                    value={newUserForm.password}
                    onChange={(e) => setNewUserForm(prev => ({...prev, password: e.target.value}))}
                  />
                </div>
                <div className='flex gap-2 justify-center'>
                  <Button variant="ghost" onClick={handleNewUserSubmit} className='p-2'>
                    <Check className="h-4 w-4 text-green-500" />
                  </Button>
                  <Button variant="ghost" onClick={() => setNewUserBKVisible(false)} className='p-2'>
                    <X className="h-4 w-4 text-red-500" />
                  </Button>
                </div>
              </div>
            )}
          </>
        )}
    </section>
  )
}

export default UserTable