<script setup lang="ts">
import type { Message } from 'ollama/browser';
import { computed } from 'vue';

// import MarkdownIt from 'markdown-it';
// const markdown = new MarkdownIt();


interface Props {
    message: Message;
}

const props = defineProps<Props>();

const messageClass = computed(() => {
    return props.message.role === 'user' ? 'message userMessage' : 'message assistantMessage';
});

const messageContentClass = computed(() => {
    return props.message.role === 'user' ? 'messageContent userMessageContent' : 'messageContent assistantMessageContent';
});

</script>

<template>
    <div :class="messageClass">
        <!-- <div :class="messageContentClass" v-html="markdown.render(props.message.content)">
        </div> -->
        <div :class="messageContentClass" v-html="props.message.content.replace(/\n/g, '<br/>')">
        </div>
    </div>
</template>

<style scoped>
.message {
    width: 100%;
    display: flex;
    padding-bottom: 10px;
    /* margin: 10px; */
}

.userMessage {
    justify-content: flex-end;
    padding-left: 10%;
}

.assistantMessage {
    justify-content: flex-start;
    padding-right: 10%;
}

.messageContent {
    /* background-color: var(--vt-c-divider-dark-2); */
    border-style: solid;
    border-width: 1px;
    border-color: var(--vt-c-divider-dark-2);
    border-top-right-radius: 0.5rem;
    border-top-left-radius: 0.5rem;
    padding: 15px;
}

.userMessageContent {
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0.5rem;
}

.assistantMessageContent {
    border-bottom-right-radius: 0.5rem;
    border-bottom-left-radius: 0;
}
</style>