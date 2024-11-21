import { GuessLangWorker } from '@ray-d-song/guesslang-js/worker'

const guessLang = new GuessLangWorker()

async function langDetect(content: string) {
  const result = await guessLang.runModel(content)
  return result[0].languageId
}

export default langDetect