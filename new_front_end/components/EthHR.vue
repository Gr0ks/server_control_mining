<template>


  
      
      <td v-bind:class="{ 'bg-danger text-white': ethhr<1, 'bg-success': ethhr>1 }">
        {{ethhr*1000 |formatHashrate}}
      </td>

  
  

</template>

<script>
export default {
  props: {
    ethhr: null,
    index: null
  }
};
</script>
