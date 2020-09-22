// init state
import {GetterTree,MutationTree} from 'vuex'

interface State {
    all: Array<Products>;
    product: string;
    version: string;
}

const state: State = {
    all: [],
    product: '',
    version: ''
}

// getters
const getters: GetterTree<State, State> = {
    getProducts(s: State) {
        return s.all
    },
    getProduct(s: State) {
        return s.product
    },
    getVersion(s: State) {
        return s.version
    }
}


// actions;
const actions = {}

const mutations: MutationTree<State> = {
    setProducts(state: State, products: Products[]) {
        state.all = products;
    },
    setProduct(state: State, product: string) {
        state.product = product;
    },
    setVersion(state: State, version: string) {
        state.version = version;
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}