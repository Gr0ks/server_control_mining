<template>      
  <td v-bind:class="{ 'bg-warning': dcrhr<1, 'bg-success': dcrhr>1 }">
    {{dcrhr*1000 |formatHashrate}}
  </td>
</template>

<script>
export default {
  props: {
    dcrhr: null,
    index: null
  }
};
</script>
