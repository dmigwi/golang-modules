# golang-modules

```
├── LICENSE
├── README.md
├── delete
│   └── drop.go
├── go.mod
├── go.sum
├── main.go
├── printer
│   ├── go.mod
│   └── print.go
└── topup
    ├── append.go
    └── go.mod

3 directories, 10 files
```

Output

```
    $ go build . && ./v1
    2019/03/24 19:04:37 Executing : github.com/dmigwi/golang-modules/topup.AddFruits
    2019/03/24 19:04:37 Executing : github.com/dmigwi/golang-modules/printer.Display
    Pineapple is at index 0 
    Tomatoes is at index 1 
    Mango is at index 2 
    Banana is at index 3 
    Watermelon is at index 4 
    2019/03/24 19:04:37 Executing : v1.AddFruits
    2019/03/24 19:04:37 Executing : github.com/dmigwi/golang-modules/printer.Display
    Pineapple is at index 0 
    Mango is at index 1 
    Banana is at index 2 
    Watermelon is at index 3 
```

## Golang modules

`github.com/dmigwi/golang-modules/printer` and `github.com/dmigwi/golang-modules/topup`

## Sub package 

`github.com/dmigwi/golang-modules/v1/delete`