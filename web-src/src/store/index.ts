import {createStore} from 'vuex'
import products from "@/store/modules/products";

export default createStore({
    modules: {
        products
    }
})
