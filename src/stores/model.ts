import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useModelStore = defineStore('model', () => {
  const currentModel = ref('llama3.2:latest')

  const setCurrentModel = (newModel: string) => {
    currentModel.value = newModel
  }

  return { currentModel, setCurrentModel }
})
