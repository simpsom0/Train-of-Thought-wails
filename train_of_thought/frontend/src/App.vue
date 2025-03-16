<script setup lang="ts">
  import SpeechOutputView from './views/SpeechOutputView.vue';
  import RecordButton from './components/RecordButton.vue';
  import NotificationCard from './components/NotificationCard.vue';
  // import { watch } from 'vue';
  // import { useRecordingState } from './constants/store';
  // import { severityLevel } from './constants/severity';

  // const isRecording = useRecordingState();
  // let notificationProps: severityLevel = severityLevel.info;
  // let notificationText: string = '';

  // watch(
  //   () => isRecording.value,
  //   (value) => {
  //     notificationProps = severityLevel.info;
  //     notificationText = value ? 'recording..' : 'stopped recording';
  //   }
  // );
</script>

<template>
  <body class="w-screen h-full min-h-screen bg-c-white transition-base">
    <div class="w-full h-full">
      <SpeechOutputView />
      <RecordButton />
      <NotificationCard />
      <!-- <NotificationCard v-bind:notificationProps>
        {{ notificationText }}
      </NotificationCard> -->
    </div>
  </body>
</template>
