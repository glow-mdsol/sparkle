const path = require('path');
const precss = require('precss');
const webpack = require('webpack');
const autoprefixer = require('autoprefixer');
const functions = require('postcss-functions');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const postCssLoader = [
    'css-loader?modules',
    '&localIdentName=[name]__[local]___[hash:base64:5]',
    '&disableStructuralMinification',
    '!postcss-loader'
];

const miniCssExtractPlugin = new MiniCssExtractPlugin({
    // Options similar to the same options in webpackOptions.output
    // both options are optional
    filename: "[name].css",
    chunkFilename: "[id].css"
});

var plugins = [
    new webpack.NoEmitOnErrorsPlugin(),
    // miniCssExtractPlugin
];

if (process.env.NODE_ENV === 'production') {
    plugins = plugins.concat([
        new webpack.optimize.UglifyJsPlugin({
            output: {comments: false},
            test: /bundle\.js?$/
        }),
        new webpack.DefinePlugin({
            'process.env': {NODE_ENV: JSON.stringify('production')}
        })
    ]);

    postCssLoader.splice(1, 1); // drop human readable names
}

const config = {
    mode: "development",
    entry: {
        bundle: path.join(__dirname, 'ui/index.js')
    },
    output: {
        path: path.join(__dirname, 'server', 'data', 'static', 'build'),
        publicPath: '/static/build/',
        filename: '[name].js'
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader",
                    query: {
                        presets: ['env', 'react']
                    }
                }
            },
            {
                test: /\.css$/,
                use: [
                    {
                        loader: "css-loader",
                        options: {
                            modules: true,
                            importLoaders: 1,
                            localIdentName: "[name]_[local]_[hash:base64]",
                            sourceMap: true,
                            minimize: true
                        }
                    }]
            }
        ]
    },
    resolve: {
        extensions: ['.js', '.jsx', '.css'],
        alias: {
            '#app': path.join(__dirname, 'ui'),
            '#c': path.join(__dirname, 'ui/components'),
            '#css': path.join(__dirname, 'ui/css')
        }
    },
    plugins: plugins
};

module.exports = config;
