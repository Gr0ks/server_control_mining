<template>
<div>  
  <tr v-bind:class="{ 'bg-danger text-white': !worker.online }">
     
        <th scope="row">
            {{name}}
            <button @click="reboot(worker.addr)" type="button" class="btn btn-info btn-sm">REBOOT</button>
        </th>

        <td > <!-- v-bind:class="{ 'd-none': !showIp }"-->
            {{worker.addr}}
        
        </td>
        <td v-if="worker.status">
            {{worker.status.time |formattime}}
        </td>
        <td v-else-if="worker.status" class="bg-danger"></td>
        <td v-if="worker.status">
            
            <table border="1" cellpadding="1" cellspacing="0">
                <tr>
                    <td>
                        ETH
                    </td>
                    <td>
                        {{worker.status.e_hashrate*1000 |formatHashrate}}
                    </td>
                    <ethhr v-for="(ethhr, index) in worker.status.e_gpu_hashrate"  
                    v-bind:key="index" :ethhr="ethhr" :index="index"/>
                </tr>
                <tr> 
                    <td>
                        DCR
                    </td>
                    <td>
                        {{worker.status.d_hashrate*1000 |formatHashrate}}
                    </td>                
                    <dcrhr v-for="(dcrhr, index) in worker.status.d_gpu_hashrate"  
                    v-bind:key="index" :dcrhr="dcrhr" :index="index"/>
                </tr>
                <tr> 
                    <td>Temper<br>
                        Fan
                    </td>
                    <td>
                        {{worker.temp}}<br>
                        ~
                    </td>                
                    <temper v-for="(temper, index) in worker.status.gpus"  
                    v-bind:key="index" :temper="temper" :index="index"/>
                </tr>
                
            </table>
        </td>
        
        
    </tr>
</div>
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
