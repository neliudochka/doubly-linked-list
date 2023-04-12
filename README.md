# doubly-linked-list
Implementation of the doubly-linked-list structure.  

## Variant
Номер заліковки - 1209 =>  
1209/4 = 302 (ост. 1) =>  
Номер варіанту= 1

## Usage
To be able to use this code you need to install Go.  
How to do it you can see [here](https://go.dev/doc/install)  

```bash
$ git clone https://github.com/neliudochka/doubly-linked-list.git
$ cd doubly-linked-list/
```

There is an example of the code with doubly-linked-list struct in the cmd/main/main.go  
To run it type:  
```bash
$ go run cmd/main/main.go
```

## Testing
To run tests type:  
```bash
$ go test
```
## Commit with failed tests
[Link](https://github.com/neliudochka/doubly-linked-list/commit/367bb0fc5e7e7b26e7166312d8afcb908761dd7e)

## Conclusion
Це був мій перший досвід у написанні тестів і, мабуть, через це я витратила на це дуже багато часу. Тим не менш, під час написання тестів я виявила кілька варіянтів подій, про які не подумала, коли прописувала методи. А під час переписування списку з використанням слайсів було дуже зручно не перевіряти все ручками, а запукати тести. Отже, тепер вже з практичного досвіду можу сказати, що тести штука корисна, особливо якщо немає обмеження по часу.
