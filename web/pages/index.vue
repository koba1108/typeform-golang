<script lang="ts" setup>
import { createWidget } from '@typeform/embed'

const formId = 'mALmrXIz'

const refreshRef = ref(() => {})
const unmountRef = ref(() => {})

onMounted(() => {
  const { refresh, unmount } = createWidget(formId, {
    container: document.getElementById('typeform') as HTMLElement, // typeformを挿入するタグを指定
    transitiveSearchParams: true, // クエリパラメータをトラッキングする設定
    shareGaInstance: '', // Google Analytics ID
    tracking: {
      // utmパラメータを追跡するための設定
      utm_source: 'example.com',
      utm_medium: 'embed',
      utm_campaign: 'campaign-1',
      utm_content: 'content-1',
      utm_term: 'term-1',
    },
    onReady: () => {
      // alert('Typeform is ready')
    },
    onEndingButtonClick: () => {
      // alert('Typeform ending button clicked')
    },
    onSubmit: () => {
      // alert('Typeform submitted')
    },
    onClose: () => {
      // alert('Typeform closed')
    },
  })
  refreshRef.value = refresh
  unmountRef.value = unmount
})

function setName() {
  const router = useRouter()
  router.push({ query: { name: 'yamada' } })
  if (refreshRef.value) {
    refreshRef.value()
  }
}
</script>

<template>
  <main class="page">
    <h1>ページタイトル</h1>
    <button @click="setName">SetName</button>
    <div class="content">
      <label>入力欄その1</label>
      <input type="text" />
    </div>
    <div class="content">
      <label>入力欄その2</label>
      <input type="text" />
    </div>
    <div id="typeform"/>
    <div class="content">
      <label>入力欄その3</label>
      <input type="text" />
    </div>
    <div class="content">
      <label>入力欄その4</label>
      <input type="text" />
    </div>
    <div class="content">
      <button>送信</button>
    </div>
  </main>
</template>

<style lang="scss">
.page {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  max-width: 800px;
  margin: 0 auto;

  .content {
    height: 120px;
    width: 100%;
    display: flex;
    gap: 20px;
    align-items: center;
  }

  #typeform {
    height: 500px;
    width: 100%;

    .tf-v1-widget {
      height: 100%;
      width: 100%;

      iframe {
        height: 100%;
        width: 100%;
      }
    }
  }
}
</style>
