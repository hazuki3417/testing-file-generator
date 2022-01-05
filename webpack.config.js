const TerserPlugin = require("terser-webpack-plugin")
const path = require("path")

/**
 * 参考情報
 * @see https://qiita.com/soarflat/items/28bf799f7e0335b68186
 * webpack公式ドキュメント
 * @see https://webpack.js.org/concepts/
 * Bable公式ドキュメント
 * @see https://babeljs.io/docs/en/
 *
 * webpack = バンドラー
 * Babel   = コンパイラー
 */
module.exports = {
  // 実行環境名を指定（オプション --mode development|production|noneと同じ）
  mode: "development",
  // エントリーポイントの設定
  entry: {
    index: "./src/index.tsx",
  },
  // 出力の設定
  output: {
    // 出力するファイル名
    filename: "[name].bundle.js",
    // 出力先のパス（絶対パスを指定する必要がある）
    path: path.join(__dirname, "./public/js"),
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"],
  },
  // ローダーの設定
  module: {
    rules: [
      // {
      // eslint-loaderの設定
      // enforce: "pre"を指定することによって
      // enforce: "pre"がついていないローダーより早く処理が実行される
      // 今回はbabel-loaderで変換する前にコードを検証したいため、指定が必要
      //   enforce: "pre",
      //   loader: "eslint-loader",
      //   test: /\[.tsx|.ts|.js]$/,
      //   exclude: /node_modules/
      // },
      {
        // babel-loaderの設定
        loader: "babel-loader",
        test: /\[.tsx|.ts|.jsx|.js]$/,
        exclude: /node_modules/,
      }
    ]
  },
  target: ["web", "es5"],
  devtool: "source-map",
  optimization: {
    splitChunks: {
      cacheGroups: {
        // 今回はvendorだが、任意の名前で問題ない
        vendor: {
          // node_modules配下のモジュールをバンドル対象とする
          test: /node_modules/,
          name: "vendor",
          chunks: "initial",
          enforce: true,
        },
      },
    },
    minimizer: [
      new TerserPlugin({
        terserOptions: {
          compress: { drop_console: true },
        },
      }),
    ],
  },
}
