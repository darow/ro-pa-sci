/* eslint-disable */
const path = require('path');
const HTMLWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
/* eslint-enable */

module.exports = env => {
    const isDev = env === 'development';
    const filename = isDev ? '[name]' : '[name]_[chunkhash]';
    const rootPath = path.resolve(__dirname, 'frontend');

    const cssLoaders = extra => {
        const loaders = [
            {
                loader: MiniCssExtractPlugin.loader
            },
            'css-loader'
        ];

        if (extra) {
            loaders.push(extra);
        }

        return loaders;
    };

    const jsLoaders = () => {
        const loaders = [
            {
                loader: 'babel-loader',
                options: {
                    presets: ['@babel/preset-env']
                }
            }
        ];

        if (isDev) {
            loaders.push('eslint-loader');
        }

        return loaders;
    };

    const config = {
        context: rootPath,
        mode: 'development',
        devtool: false,
        optimization: {
            splitChunks: {
                chunks: 'all'
            }
        },
        entry: './source/index.js',
        output: {
            filename: `${filename}.js`,
            path: path.resolve(rootPath, 'distribute')
        },
        resolve: {
            extensions: ['js', 'json'],
            alias: {
                frontend: path.resolve(__dirname, './frontend'),
                source: path.resolve(__dirname, './frontend/source'),
                '@': path.resolve(__dirname, './frontend/source/blocks')
            }
        },
        plugins: [
            new CleanWebpackPlugin(),
            new HTMLWebpackPlugin({
                template: path.resolve(rootPath, './source/templates/game.html'),
                filename: 'game.html'
            }),
            new MiniCssExtractPlugin({
                filename: `${filename}.css`,
                chunkFilename: '[id].css'
            })
        ],
        module: {
            rules: [
                {
                    test: /\.css$/,
                    use: cssLoaders()
                },
                {
                    test: /\.s[ac]ss$/,
                    use: cssLoaders('sass-loader')
                },
                {
                    test: /\.js$/,
                    exclude: /node_modules/,
                    use: jsLoaders()
                }
            ]
        }
    };

    return config;
};
