# learning-transaction
it is a simple transaction test

1. 正常情況下，交易只針對資料的隔離性做處理，也就是讀取的部分, 如果要防止寫入還是要靠 lock 
2. 當 rows 被`select for update` 鎖住的時候，如果要針對這筆資料做寫入的動作, 必須要等 for update 的鎖被解除

Reference: 
http://www.asktheway.org/2020/04/12/261/
https://blog.xuite.net/vexed/tech/22289223-%E7%94%A8+SELECT+...+FOR+UPDATE+%E9%81%BF%E5%85%8D+Race+condition