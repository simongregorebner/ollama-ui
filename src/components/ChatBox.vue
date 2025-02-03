<script setup lang="ts">
import ChatMessage from '../components/ChatMessage.vue'
import { useModelStore } from '@/stores/model';

import { ref, toRefs, useTemplateRef } from 'vue'



const modelStore = useModelStore();
const { currentModel } = toRefs(modelStore);
// const { currentModel } = ref(modelStore);

// const chatInput = ref('Why is the sky blue?')

// const submitChat = () => {
//     console.log('submitting chat', chatInput.value)
//     chatInput.value = ''
// }

import ollama from 'ollama';
import type { RefSymbol } from '@vue/reactivity';
import { Textarea } from 'primevue';

// const messages = ref([{ role: 'agent', content: 'Lets start ...' }]);
const messages = ref([{ role: 'agent', content: '...' }]);
messages.value.pop(); // remove the first message again so that is does not show up
const currentOutputMessageContent = ref('');

// const chatInput = ref('Why is the sky blue?')
const chatInput = ref()
const chatInputBoxRef = useTemplateRef('chatInputBoxRef')

// const submitChat = async () => {
//     const content = chatInput.value;
//     chatInput.value = '';
//     const response = await ollama.chat({
//         model: 'llama3.2',
//         messages: [{ role: 'user', content }],
//     });
//     console.log(response.message.content);
// }

const submitChat = async () => {
    const content = chatInput.value;
    chatInput.value = '';
    const inputMessage = { role: 'user', content };
    messages.value.push(inputMessage);
    // const response = await ollama.chat({ model: 'llama3.2', messages: [inputMessage], stream: true });
    const response = await ollama.chat({ model: currentModel.value, messages: [inputMessage], stream: true });
    for await (const part of response) {
        currentOutputMessageContent.value += part.message.content;
    }
    messages.value.push({ role: 'agent', content: currentOutputMessageContent.value });
    currentOutputMessageContent.value = '';
};

const focusInput = () => {
    // if (chatInputBoxRef.value) { // Check if the ref is set
    //     console.log(chatInputBoxRef.value)
    // }
};

</script>

<template>
    <div id="chatBox">

        <div id="chatContainer">
            <div id="chatArea" ref="chatArea">
                <div class="logolama">
                    <div>
                        <img class="logo" src="@/assets/ollama.svg" width="80" />
                    </div>
                    <div>
                        <!-- <h2>Chat with large language models locally.</h2> -->
                        <h2>Ollama UI - Chat locally.</h2>
                    </div>
                    <div>
                        Chat and with any Ollama
                        supported model locally.
                    </div>
                    <div>
                        <Button label="Start" @click="focusInput" />
                    </div>
                </div>

                <div v-for="message in messages" :key="message.content">
                    <ChatMessage :message="message" />
                </div>
                <div v-if="currentOutputMessageContent">
                    <ChatMessage :message="{ role: 'agent', content: currentOutputMessageContent }" />
                </div>
            </div>
        </div>
        <!-- <div id="inputArea">
            <input id="chatInput" v-model="chatInput" @keyup.enter="submitChat" />
            <button @click="submitChat" id="submitButton">Submit</button>
        </div> -->
        <div id="inputArea">
            <Textarea v-model="chatInput" @keyup.enter="submitChat" placeholder="Enter your prompt here ..."
                id="chatInputBox" ref="chatInputBoxRef" />
            <Button @click="submitChat" id="submitButton"> <i class="pi pi-send" style="font-size: 1rem"></i></Button>
        </div>

    </div>
</template>

<style scoped>
/* .logo {
    align-self: center;
} */

.logolama {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-bottom: 100px;
}

#chatBox {
    display: flex;
    height: 100%;
    flex-direction: column;
}

#chatContainer {
    position: relative;
    width: 100%;
    height: 100%;
}

#chatArea {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    overflow-y: auto;
    align-content: flex-end;
    padding: 10px;
}

#inputArea {
    display: flex;
    height: 100px;
    width: 100%;
    padding: 10px;
    align-items: space-between;
}

#chatInputBox {
    width: 100%;
    height: 100%;
    padding: 10px;
}
</style>