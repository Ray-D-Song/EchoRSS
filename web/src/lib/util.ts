interface Reactive<T> {
  value: T
}

const effectMap = new WeakMap<object, Array<() => void>>();

function reactive<T>(value: T): Reactive<T> {
  const raw = new Proxy({ value }, {
    set(target, prop, newValue) {
      target[prop as keyof typeof target] = newValue
      const cbs = effectMap.get(raw)
      if (cbs) cbs.forEach(cb => cb())
      return true
    }
  })
  return raw
}

function util<T>(reactive: Reactive<T>, targetValue: unknown, failedValue: unknown): Promise<boolean> {
  return new Promise(resolve => {
    if (reactive.value === targetValue) {
      resolve(true)
      return;
    }
    const cbs = effectMap.get(reactive) || []
    cbs.push(() => {
      if (reactive.value === targetValue) {
        resolve(true)
      } else if (reactive.value === failedValue) {
        resolve(false)
      }
    })
  })
}

async function Test() {
  const a = reactive(1)
  setTimeout(() => {
    a.value = 2
  }, 4000)
  await util(a, 2, 3)
}

Test().then(() => {
  console.log('done')
})

export { reactive, util }