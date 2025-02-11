<script setup lang="ts">
import ChatMessage from '../components/ChatMessage.vue'
import { useModelStore } from '@/stores/model';
import { ref, toRefs, useTemplateRef } from 'vue'

const modelStore = useModelStore();
const { currentModel } = toRefs(modelStore);

import { Textarea } from 'primevue';

// import ollama, { type Message } from 'ollama/browser';
// Use proxy instead of direct ollama communication
import { Ollama, type Message } from 'ollama/browser';
const ollama = new Ollama({
    host: window.location.origin + '/ollama/', // Replace with your actual URL
});

const messages = ref(<Message[]>[]);
const currentOutputMessageContent = ref('');

// const chatInput = ref('Why is the sky blue?')
const chatInput = ref('')
const chatInputBoxRef = useTemplateRef('chatInputBoxRef')

const handleKeyDown = (event: KeyboardEvent) => {
    if (event.key === 'Enter' && event.shiftKey) {
        console.log('Shift+Enter pressed!');
        // Your Shift+Enter logic here (e.g., insert newline)
        // If you don't want the default newline, use: event.preventDefault();
    } else if (event.key === 'Enter') {
        console.log('Enter pressed!');
        event.preventDefault(); // Prevent default newline so submitChat can control it.
        submitChat(); // Call your submitChat function
    }
};

const submitChat = async () => {
    // console.log(imageFile.replace(new RegExp('data:.*/.*;base64,'), ''));

    const content = chatInput.value;
    chatInput.value = '';
    // const inputMessage = { role: 'user', content };
    let inputMessage: Message;
    if (imageFile != '') {
        inputMessage = { role: 'user', content: content, images: [imageFile.replace(new RegExp('data:.*/.*;base64,'), '')] };
    }
    else {
        inputMessage = { role: 'user', content: content };
    }
    imageFile = ''

    messages.value.push(inputMessage);

    // const response = await ollama.chat({ model: currentModel.value, messages: [inputMessage], stream: true });
    const response = await ollama.chat({ model: currentModel.value, messages: messages.value, stream: true });
    for await (const part of response) {
        currentOutputMessageContent.value += part.message.content;
    }
    messages.value.push({ role: 'assistant', content: currentOutputMessageContent.value });
    currentOutputMessageContent.value = '';
};


let imageFile: string = '';

const uploadFile = (event: Event) => {
    const target = event.target as HTMLInputElement
    if (target && target.files) {
        const file = target.files[0];
        const reader = new FileReader();
        reader.onload = (e) => {
            if (e.target && e.target.result) { imageFile = e.target.result.toString() }
        };
        reader.readAsDataURL(file);
    }
    // const file = event.target.files[0];
    // console.log("have file" + file.name)
};

const clearChat = () => {
    // messages.value.pop()
    messages.value = <Message[]>[]
}

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
                    <ChatMessage :message="{ role: 'assistant', content: currentOutputMessageContent }" />
                </div>
            </div>
        </div>
        <!-- <div id="inputArea">
            <input id="chatInput" v-model="chatInput" @keyup.enter="submitChat" />
            <button @click="submitChat" id="submitButton">Submit</button>
        </div> -->
        <div id="inputArea">
            <div id="inputTextArea">
                <Textarea v-model="chatInput" @keydown="handleKeyDown" placeholder="Enter your prompt here ..."
                    id="chatInputBox" ref="chatInputBoxRef" style="box-shadow: none" />
            </div>
            <div id="inputButtonArea">
                <!-- <Button @click="submitChat" icon="pi pi-send" rounded outlined severity="secondary"></Button> -->
                <div>
                    <label>
                        <input type="file" @change="uploadFile" style="display: none" />
                        <i class="pi pi-images" title="Load File"></i>
                    </label>
                    <label>
                        <button @click="submitChat" style="display: none" />
                        <i class="pi pi-send" title="Submit Message"></i>
                    </label>
                </div>
                <div>
                    <label>
                        <button @click="clearChat" style="display: none;" />
                        <i class="pi pi-trash" title="Clear Chat" style="color: var(--custom-red)"></i>
                    </label>
                </div>


                <!-- <FileUpload mode="basic" @select="onFileSelect" icon="pi pi-images" customUpload auto
                    severity="secondary" class="p-button-outlined" /> -->
            </div>
        </div>

    </div>
</template>

<style scoped>
/* .logo {
    align-self: center;
} */

label {
    padding: 10px;
}

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

#inputTextArea {
    padding-top: 10px;
    padding-left: 10px;
    padding-right: 10px;
}

#inputButtonArea {
    display: flex;
    justify-content: space-between;
    padding-bottom: 10px;
    padding-left: 10px;
    padding-right: 10px;
}

#inputArea {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-left: 10px;
    margin-right: 10px;
    background-color: var(--custom-2);
    border-radius: 0.5rem;
    border-width: 0;
    /* align-items: space-between; */
}

#chatInputBox {
    width: 100%;
    height: 60px;
    /* padding: 10px; */
    background-color: var(--custom-2);
    border-width: 0;
}
</style>