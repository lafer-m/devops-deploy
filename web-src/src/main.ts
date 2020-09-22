import {createApp} from 'vue'
import {Button, Checkbox, Input, Popconfirm, Select, Table} from 'ant-design-vue'
import App from './App.vue'
import router from './router'
import store from './store'
// 引入style文件
import 'ant-design-vue/dist/antd.css'

createApp(App).use(Select)
    .use(Table)
    .use(Button)
    .use(Input)
    .use(Popconfirm)
    .use(Checkbox)
    .use(store).use(router).mount('#app')
