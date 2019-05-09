const path = require('path');

module.exports = {
    entry: './resources/index.js',
    mode :"development",
    output: {
        path: path.resolve(__dirname, '../dist'),
        filename: 'demo.js'
    },
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'resources/assets/js'),
            '#': path.resolve(__dirname, 'resources/assets/sass')
        }
    },
    module: {
        
        rules : [
            {
                test: /\.css$/,
                use: ['style-loader','css-loader']
            }
        ]
    }
}