import { ModeToggle } from './mode-toggle'

export function Header() {

  return (
    <header className="flex items-center justify-between p-4 border-b">
      <section className="flex items-center gap-2">
        <img src="/logo.svg" alt="EchoRSS" className="w-10 h-10" />
        <div className="font-bold flex flex-col">
          <span className="text-sm mb-[-6px]">Echo</span>
          <span className="text-xl">RSS</span>
        </div>
      </section>
      <ModeToggle />
    </header>
  )
}
