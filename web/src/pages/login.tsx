
import { useContext, useState } from 'react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { UserIcon, LockIcon } from 'lucide-react'
import { useNavigate } from 'react-router-dom'
import { AppCtx } from '@/lib/ctx'
import useFetch from '@/hooks/use-fetch'

function LoginPageComponent() {

  const { setUser } = useContext(AppCtx)

  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  const navigate = useNavigate()

  const { loading, run } = useFetch<User>('/auth/login', {
    method: 'POST',
    body: JSON.stringify({ username, password })
  }, {
    onSuccess: (data) => {
      setUser(data)
      navigate('/')
    }
  })

  return (
    <div className="min-h-screen flex flex-col justify-center items-center p-4">
      <div className="w-full max-w-md">
        <Card>
          <form onSubmit={(e) => {
            e.preventDefault()
            run()
          }}>
            <CardHeader className="flex flex-row items-center gap-2">
              <img src="/logo.svg" alt="logo" className="w-12 h-12" />
              <div>
                <CardTitle>EchoRSS</CardTitle>
                <CardDescription>Please enter your account and password</CardDescription>
              </div>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="account">Account</Label>
                <div className="relative">
                  <UserIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
                  <Input
                    id="account"
                    type="text"
                    placeholder="Please enter your account"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    className="pl-10"
                    required
                  />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <div className="relative">
                  <LockIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
                  <Input
                    id="password"
                    type="password"
                    placeholder="Please enter your password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="pl-10"
                    required
                    autoComplete="current-password"
                  />
                </div>
              </div>
            </CardContent>
            <CardFooter className="flex flex-col items-center gap-2">
              <Button type="submit" className="w-full" disabled={loading}>
                {loading ? 'Loading...' : 'Login'}
              </Button>
              <h6 className="text-center text-[12px] text-gray-500">
                Please contact the administrator,
                <br />
                if you want to register.
              </h6>
            </CardFooter>
          </form>
        </Card>
      </div>
    </div>
  )
}

export default LoginPageComponent
