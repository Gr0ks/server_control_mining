<template>
  

<div class="m-1"> 
     
  
    <b-card border-variant="dark"  class="text-center" bg-variant="primary" >
  
      <p class="card-text">
        <span class="badge badge-success text-center">{{name}}</span>
     <!-- <button @click= "showIp=!showIp" type="button" class="btn btn-info btn-sm">SHOW IP</button>-->
        
      </p>
      
    
   
    <div class="d-flex flex-wrap" >
     <worker v-for="(worker, name) in miner" 
        v-bind:key="name" :worker="worker" :name="name" />
    </div>
    </b-card>
  
</div>

 
  
</template>

<script>
import Worker from "~/components/Worker.vue";
export default {
  props: {
    miner: null,
    name: null
  },
  components: {
    Worker
  }
};
</script>
