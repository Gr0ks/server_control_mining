import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = () => new Vuex.Store({

    state: {
        miners: null
    },
    mutations: {
        setMiners(state, miners) {
            state.miners = miners
        }
    },
    actions: {
        async nuxtServerInit({ commit }, { req }) {
            const res = await this.$axios.get("");
            commit('setMiners', res.data)
        },
        async getApi({ commit }) {
            const res = await this.$axios.get("");
            commit('setMiners', res.data)
        }
    }
})

export default store