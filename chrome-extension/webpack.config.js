const path = require('path');
const HtmlWebpackPlugin = require("html-webpack-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {
    entry: {
        popup: './src/popup/popup.tsx',
        options: './src/options/options.tsx',
        background: {
            import: './src/sw/background.ts',
            filename: 'background.js'
        }
    },
    devtool: "inline-source-map",
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname, './dist'),
        clean: true,
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: "./src/popup/popup.html",
            filename: "popup.html",
            chunks: ['popup'],
            scriptLoading: "module",
        }),
        new HtmlWebpackPlugin({
            template: "./src/options/options.html",
            filename: "options.html",
            chunks: ['options'],
            scriptLoading: "module",
        }),
        new CopyWebpackPlugin({
            patterns: [
                {from: './public', to: '.'}
            ]
        }),
    ],
    mode: "development",
    module: {
        rules: [
            {
                test: /.(ts|tsx)$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            }
        ]
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js'],
    }
}
