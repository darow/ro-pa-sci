const path = require('path');
const HTMLWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = env => {
    const cssLoaders = extra => {
        const loaders = [
            {
                loader: MiniCssExtractPlugin.loader,
            },
            'css-loader',
        ]

        if (extra) {
            loaders.push(extra);
        }

        return loaders;
    }

    const isDev = env === 'development'
    const filename = isDev ? '[name]' : '[name]_[chunkhash]';
    const rootPath = path.resolve(__dirname, 'frontend');

    const config = {
        context: rootPath,
        mode: 'development',
        devtool: false,
        optimization: {
            splitChunks: {
                chunks: 'all'
            },
        },
        entry: './source/index.js',
        output: {
            filename: `${filename}.js`,
            path: path.resolve(rootPath, 'distribute'),
        },
        resolve: {
            extensions: ['js', 'json'],
        },
        plugins: [
            new CleanWebpackPlugin(),
            new HTMLWebpackPlugin({
                template: path.resolve(rootPath, './source/index.html')
            }),
            new MiniCssExtractPlugin({
                filename: `${filename}.css`,
                chunkFilename: '[id].css',
            }),
        ],
        module: {
            rules: [
                {
                    test: /\.css$/,
                    use: cssLoaders(),
                },
                {
                    test: /\.s[ac]ss$/,
                    use: cssLoaders('sass-loader'),
                },
                {
                    test: /\.js$/,
                    exclude: /node_modules/,
                    use: {
                        loader: 'babel-loader',
                        options: {
                            presets: ['@babel/preset-env'],
                        },
                    },
                },
            ],
        },
    };

    return config;
}