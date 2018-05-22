<template>
       
  <b-card border-variant="dark"  class="text-center m-1 " >
  
    <p class="card-text " v-if="worker.status">
      <b>{{name}}</b> ip: {{worker.addr}} 
      Time RUN: {{worker.status.time |formattime}} <br>
      ETH: {{worker.status.e_hashrate*1000 |formatHashrate}} 
      DCR: {{worker.status.d_hashrate*1000 |formatHashrate}}
      <table border="1" cellpadding="1" cellspacing="0" v-if="worker.status">
        <tr>
          <td>eth</td>            
            <ethhr v-for="(ethhr, index) in worker.status.e_gpu_hashrate"  
            v-bind:key="index" :ethhr="ethhr" :index="index"/>
        </tr>
        <tr> 
          <td>dcr</td>              
            <dcrhr v-for="(dcrhr, index) in worker.status.d_gpu_hashrate"  
            v-bind:key="index" :dcrhr="dcrhr" :index="index"/>
        </tr>
        <tr> 
            <td>Temper<br>
                Fan
            </td>
                            
            <temper v-for="(temper, index) in worker.status.gpus"  
            v-bind:key="index" :temper="temper" :index="index"/>
        </tr>
        
      </table> 
      
        <button @click="reboot(worker.addr)" type="button" class="m-1 btn btn-info btn-sm">REBOOT</button>
      

    </p>
    <p class="card-text bg-danger" v-if="!worker.status">
        <b>{{name}}</b> ip: {{worker.addr}}
    </p>
  </b-card>
 


</template>

<script>
import ethhr from "~/components/EthHR.vue";
import dcrhr from "~/components/DcrHR.vue";
import temper from "~/components/Temper.vue";

export default {
  props: {
    worker: null,
    name: null
  },
  components: {
    ethhr,
    dcrhr,
    temper
  }
};
</script>
