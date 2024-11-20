// @ts-expect-error
const guessLang = new GuessLang()

async function langDetect(content: string) {
  console.log(content)
  const result = await guessLang.runModel(content)
  console.log(result)
  return result[0].languageId
}

export default langDetect