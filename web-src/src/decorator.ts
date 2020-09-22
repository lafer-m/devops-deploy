import {createDecorator} from "vue-class-component";

export const Watch = (valueKey: string) => {
    return createDecorator((options, handler) => {
        if (!options.watch) {
            options.watch = {}
        }
        options.watch[valueKey] = {
            handler: handler,
            deep: true
        }
    })
}

export const Prop = (cops: {}) => {
    if (!cops) {
        cops = {}
    }
    return createDecorator((options, key) => {
        if (!options.props) {
            options.props = {}
        }
        options.props[key] = cops
    })
}




