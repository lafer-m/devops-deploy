// eslint-disable-next-line @typescript-eslint/no-var-requires
const path = require('path')
// eslint-disable-next-line @typescript-eslint/no-var-requires
const TerserPlugin = require('terser-webpack-plugin')

function resolve(dir) {
    return path.join(__dirname, dir)
}

module.exports = {
    lintOnSave: true,
    // publicPath: '/dist',
    assetsDir: 'static',
    devServer: {
        overlay: {
            warnings: true,
            errors: true
        },
        // port: 8080
    },
    // configureWebpack: (config) => {
    //     if (process.env.NODE_ENV === 'production') {
    //         // 为生产环境修改配置
    //         config.mode = 'production'
    //         config.devtool = false
    //         // 将每个依赖包打包成单独的js文件
    //         const optimization = {
    //             minimizer: [
    //                 // eslint-disable-next-line no-undef
    //                 new TerserPlugin({
    //                     exclude: /\/node_modules/,
    //                     parallel: true,
    //                     terserOptions: {
    //                         warnings: false,
    //                         compress: {
    //                           // eslint-disable-next-line @typescript-eslint/camelcase
    //                             drop_console: true,
    //                           // eslint-disable-next-line @typescript-eslint/camelcase
    //                             drop_debugger: false,
    //                           // eslint-disable-next-line @typescript-eslint/camelcase
    //                             pure_funcs: ['console.log']
    //                         }
    //                     }
    //                 })
    //             ]
    //         }
    //         Object.assign(config, {
    //             optimization
    //         })
    //     } else {
    //         // 为开发环境修改配置
    //         config.mode = 'development'
    //     }
    // },
    chainWebpack: (config) => {
        config.resolve.alias
            .set('@', resolve('src'))
            .set('api', resolve('src/api'))
            .set('assets',resolve('src/assets'))
            .set('components', resolve('src/components'));
    },
    pluginOptions: {
        'style-resources-loader': {
            preProcessor: 'scss',
            patterns: [
                path.resolve(__dirname, './src/style/placeholders/*.scss'),
                path.resolve(__dirname, './src/style/mixins/*.scss'),
                path.resolve(__dirname, './src/style/common/var.scss'),
            ],
        },
    },
    css: {
        loaderOptions: {
            less: {
                lessOptions: {
                    javascriptEnabled: true
                }
            }
        }
    }
}
