# メモ書き
チュートリアルを終わった後は、サンプルを見る？

[サンプル](https://revel.github.io/samples/)

# gorpを使う

## 参考
[Revelでgorpを使ってDBを操作する](http://qiita.com/k0kubun/items/538ea0dd57800b8d7ca6)

## モデルの作成
railsっぽい

models/user.go

## gorpのコントローラーを追加
controllers/gorp.go

### 参考
[Go言語向けの ORM、gorp がなかなか良い](http://mattn.kaoriya.net/software/lang/go/20120914222828.htm)

# 実装メモ
---出来た
* dbに追加削除
* 画面から追加
* UserモデルをTodoモデルに変更
* 画面からタスクの完了が出来る

---まだ
* TodoモデルのValidation
* Validation失敗時の処理
* bootstrapを入れる
* DB周りのリファクタリング
  参考:http://yuroyoro.hatenablog.com/entry/2014/06/15/173043
