const path = require('path')

const NODE_ENV = process.env.NODE_ENV

let mode = 'development'
if (NODE_ENV && NODE_ENV.trim() === 'production') {
  mode = 'production'
}

console.log('ENV: ', mode)

module.exports = {
  mode,
  entry: path.resolve(__dirname, 'client/index.jsx'),
  output: {
    path: path.resolve(__dirname, 'assets/js'),
    publicPath: '/js/',
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.js(x?)$/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              presets: ['env', 'react']
            }
          }
        ],
        exclude: /node_modules/
      }
    ]
  },
  resolve: {
    modules: ['node_modules'],
    extensions: ['.js', '.jsx']
  },
  plugins: [],
  cache: true
}
