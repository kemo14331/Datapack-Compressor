# Datapack-Compressor
Minecraftのデータパックを圧縮するツール  

## Feature  
* 簡単な操作でデータパックを圧縮することができる  
* コメント行や改行のみの行を削除できる  
* データパックとは関係のないファイルを除外できる  

## Usage　
### 基本的な使用法
  圧縮したいデータパックフォルダを実行ファイルにドラッグ&ドロップするだけ！
　　
### コマンドオプション
|短いオプション | 長いオプション | 説明 |
|:---|:---|:---|
|-o |--output-path [PATH] |ファイルの出力先パス(ファイル名可) |
|-d |--do-not-remove-cmt |ファンクションのコメント行を削除しない |
|-f |--exclude-remove-file [String]|削除から除外するファイルパスを正規表現で指定(複数可)|
|-s | --show-log |圧縮したファイル一覧を表示する|  
  
  #### コマンド例
  ```console
  dpc src/mydatapack -o release/release_v1.zip -s -f update_log\.txt
  ```  
  mydatapackが圧縮され、release_v1.zipとして出力されます。  
  `update_log.txt`は除外されずにzipファイル内に出力されます。
  
## Downloads  
 [Release](https://github.com/kemo14331/Datapack-Compressor/releases)

## Author  
* Kemo431  
* Twitter: [@newkemo431](https://twitter.com/newkemo431)  
 
## License
This app is under the [MIT license](https://en.wikipedia.org/wiki/MIT_License).
