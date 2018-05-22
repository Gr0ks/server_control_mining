<template>      
  <td v-bind:class="{ 'bg-danger text-white': temper[0,0]>80, 'bg-success': temper[0,0]<80  }">
    {{temper[0,0]}}<br>
    {{temper[0,1]}}
  </td>
</template>

<script>
export default {
  props: {
    temper: null,
    index: null
  }
};
</script>

