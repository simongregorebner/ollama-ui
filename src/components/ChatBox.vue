<script setup lang="ts">
import ChatMessage from '../components/ChatMessage.vue'
import { useModelStore } from '@/stores/model';

import { ref, toRefs } from 'vue'



const modelStore = useModelStore();
const { currentModel } = toRefs(modelStore);
// const { currentModel } = ref(modelStore);

// const chatInput = ref('Why is the sky blue?')

// const submitChat = () => {
//     console.log('submitting chat', chatInput.value)
//     chatInput.value = ''
// }

import ollama from 'ollama';

const messages = ref([{ role: 'agent', content: 'Hello, I am Lliam. How can I help you?' }]);
const currentOutputMessageContent = ref('');

const chatInput = ref('Why is the sky blue?')

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

</script>

<template>
    <div id="chatBox">
        <div id="chatContainer">
            <!-- <div id="chatArea" ref="chatArea">
                <div v-for="message in messages" :key="message.content">
                    {{ message.content }}
                </div>
                <div v-if="currentOutputMessageContent">
                    {{ currentOutputMessageContent }}
                </div>
            </div> -->
            <div id="chatArea" ref="chatArea">
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
            <Textarea v-model="chatInput" @keyup.enter="submitChat" id="chatInput" />
            <Button @click="submitChat" id="submitButton">Submit</Button>
        </div>
    </div>
</template>

<style scoped>
#chatBox {
    display: flex;
    height: 100%;
    flex-direction: column;
}

#chatContainer {
    position: relative;
    width: 100%;
    height: calc(100% - 100px);
}

#chatArea {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    overflow-y: auto;
}

#inputArea {
    display: flex;
    height: 100px;
    width: 100%;
    padding: 10px;
    align-items: space-between;
}

#chatInput {
    width: calc(100% - 82px);
    height: 100%;
    padding: 10px;
    margin-right: 10px;
}
</style>