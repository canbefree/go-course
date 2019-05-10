// var ExtractTextPlugin = require('extract-text-webpack-plugin');
const MinCssExtractPlugin = require('mini-css-extract-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin')
const webpack = require('webpack')
const path = require('path');
const VueLoaderPlugin = require('vue-loader/lib/plugin');

module.exports = {
    entry: {
        // 'vendor': path.resolve(__dirname, 'resources/assets/js/vendor.js'),
        'vendor': ['vue'],
        'main': path.resolve(__dirname, 'resources/assets/js/main.js'),
    },
    mode: "development",
    output: {
        path: path.resolve(__dirname, '../../dist/demo'),
        filename: '[name].[contenthash].js',
        chunkFilename: 'js/[name].[chunkhash].js'
    },
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'resources/assets/js'),
            '#': path.resolve(__dirname, 'resources/assets/sass'),
            'vue$': 'vue/dist/vue.common.js'
        },
        extensions: ['.js', '.vue']
    },
    module: {
        rules: [
            {
                test: /\.js$/,

                exclude: /node_modules/,
                loader: 'babel-loader',
                options: {
                    plugins: ['@babel/plugin-syntax-dynamic-import']
                }
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            },
            {
                test: /\.sass$/,
                // use: ['style-loader', 'css-loader', 'sass-loader']
                // use:ExtractTextPlugin.extract(['css-loader', 'sass-loader'])
                use: [MinCssExtractPlugin.loader, 'css-loader', 'sass-loader']
            },
            {
                test: /\.vue$/,
                use: ['vue-loader']
            }
        ]
    },
    plugins: [
        new MinCssExtractPlugin({
            filename: 'css/[name].[hash].css'
        }),

        new HtmlWebpackPlugin({
            filename: 'index.html',    //生成的html存放路径，相对于 path
            chunckTemplate: 'js',
            template: 'view/index.html',    //html模板路径
            inject: true,    //允许插件修改哪些内容，true/'head'/'body'/false,
            chunks: ['main', 'vendor'],//加载指定模块中的文件，否则页面会加载所有文件
            hash: false,    //为静态资源生成hash值
            minify: {    //压缩HTML文件
                removeComments: true,    //移除HTML中的注释
                collapseWhitespace: true//删除空白符与换行符
            }
        }),

        //Vue-loader在15.*之后的版本都是 vue-loader的使用都是需要伴生 VueLoaderPlugin的
        new VueLoaderPlugin()
    ],
    optimization: {
        splitChunks: {
            cacheGroups: {
                commons: {
                    name: "vendor",
                    chunks: "initial",
                    minChunks: 1
                }
            }
        },
        minimize: true
    },

}

