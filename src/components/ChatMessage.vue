<script setup lang="ts">
import type { Message } from 'ollama';
import { computed } from 'vue';
import MarkdownRenderer from "./MarkdownRenderer.vue";

interface Props {
    message: Message;
}

const props = defineProps<Props>();

const messageClass = computed(() => {
    return props.message.role === 'user' ? 'message userMessage' : 'message agentMessage';
});

const messageContentClass = computed(() => {
    return props.message.role === 'user' ? 'messageContent userMessageContent' : 'messageContent agentMessageContent';
});

</script>

<template>
    <div :class="messageClass">
        <div :class="messageContentClass">
            {{ props.message.content }}
            <!-- <MarkdownRenderer :source="props.message.content"></MarkdownRenderer> -->
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

.agentMessage {
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

.agentMessageContent {
    border-bottom-right-radius: 0.5rem;
    border-bottom-left-radius: 0;
}
</style>