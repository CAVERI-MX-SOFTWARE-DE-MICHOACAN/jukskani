
<script  lang="ts">
import { defineComponent, ref } from 'vue'
import {API_BASE} from '../constants'

export default defineComponent({
    props:{
        Index:{type:String, required:true}, 
        State:{type:String, required:true}, 
    },
    
    setup(props) {
        const Enabled = ref(parseInt(props.State))

        const Toggle = ()=>{
            console.log(Enabled.value)
            Enabled.value=1-Enabled.value
            setState(props.Index, Enabled.value)
        }
        const setState=(relay_index: String, state: Number)=>{
            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/json");

            var raw = JSON.stringify({
                "State": state?false:true
            });

            var requestOptions:RequestInit = {
                method: 'POST',
                headers: myHeaders,
                body: raw,
                redirect: 'follow'
            };

            fetch(`${API_BASE}/relays/${props.Index}`, requestOptions)
                .then(response => response.text())
                .then(result => console.log(result))
                .catch(error => console.log('error', error));
        }
        return {
            Enabled, Toggle
        }
    },

   
})
</script>
<template>
  <div class="py-16">
    <Switch
      v-model="Enabled"
      @click="Toggle"
      :class="Enabled ? 'bg-teal-900' : 'bg-teal-700'"
      class="relative inline-flex h-[38px] w-[74px] shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75"
    >
      <span class="sr-only">Use setting</span>
      <span
        aria-hidden="true"
        :class="Enabled ? 'translate-x-9' : 'translate-x-0'"
        class="pointer-events-none inline-block h-[34px] w-[34px] transform rounded-full bg-white shadow-lg ring-0 transition duration-200 ease-in-out"
      />
    </Switch>
  </div>
</template>

