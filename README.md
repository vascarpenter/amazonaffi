# amazonaffi

- amazon でアフィリエイトツールバーがtextだけになっていて、右側の画像付きアフィリンクが死んでいた
- しょうがないので ASINを入れるとHTMLタグを作ってくれるプログラムを組んだ
- moduleを使いたかったが、IntelliJでJava/kotlinでgradleやらmavenの設定がうまくいかなかったので新しめのAPIを使えるツールがgithubで公開されている golangを使った
- ちなみにこのブログのアフィではこれまで500円くらいしか稼げていません　なんとかしてください！

# build
- `go build`

# run
- 実行前に環境変数を設定すること
  - PA_ACCESS_KEY amazon API access key
  - PA_SECRET_KEY  amazon API secret key
  - PA_ASSOCIATE_TAG  アフィアソシエイトtag (例 gikohadiary-22)

- 実行方法 ASINを渡す
```
./amazonaffi B018WNIBJS
  <a href="https://www.amazon.co.jp/dp/B018WNIBJS/?tag=gikohadiary-22" target="_blank"><img src="https://m.media-amazon.com/images/I/41V7ZyqoirL._SL75_.jpg" alt="B018WNIBJS" border="0" />
<br>エレコム USB ゲーミングマウス 【DUX】 有線 14ボタン 3500dpi ハードウェアマクロ対応 M-DUX50BK</a>
```

# todo
- 複数ASIN渡すと`<td>`とかしてくれるやつ

