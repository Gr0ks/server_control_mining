<template>
  <section class="p-5">
    <div>
      <miner v-for="(miner, name) in $store.state.miners.miner" 
      v-bind:key="name" :miner="miner" :name="name"/>
     <br>
    </div>
  </section>
</template>

<script>
import Miner from "~/components/Miner.vue";

let interval = 0;

export default {
  components: {
    Miner
  },
  created() {
    if (!this.$isServer) {
      interval = setInterval(() => {
        this.$store.dispatch("getApi");
      }, 10000);
    }
  },
  beforeDestroy() {
    clearInterval(interval);
  },
  methods: {}
};
</script>
