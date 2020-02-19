# gil

Formatting Git log output as [JSON Lines](http://jsonlines.org/).

## Installation

```bash
$ go get -u github.com/ariarijp/gil
```

## Usage

```
$ # Clone any Git repo.
$ git clone https://github.com/ariarijp/gil.git
$ # Move to Git repo directory.
$ cd gil
$ # Run `gil` in the directory, you'll see Git logs as JSON Lines format.
$ gil
{"Hash":{"Long":"038d5a988af7f7654f8aba3f8cc7b5e76eee909c","Short":"038d5a9"},"Tree":{"Long":"c9fff68e6fe832f912c788dd2fb2324448f08cb5","Short":"c9fff68"},"Author":{"Name":"ariarijp","Email":"takuya.arita@gmail.com","Date":"2020-02-12T00:13:53+09:00"},"Committer":{"Name":"ariarijp","Email":"takuya.arita@gmail.com","Date":"2020-02-12T00:13:53+09:00"},"Subject":"Update README","Body":""}
{"Hash":{"Long":"57aa24d593b272a4bb329d904ec1563cf4ed142f","Short":"57aa24d"},"Tree":{"Long":"987b0b7791864e7e4d2be05cb01075f16b1c8c86","Short":"987b0b7"},"Author":{"Name":"ariarijp","Email":"takuya.arita@gmail.com","Date":"2020-02-12T00:12:00+09:00"},"Committer":{"Name":"ariarijp","Email":"takuya.arita@gmail.com","Date":"2020-02-12T00:12:48+09:00"},"Subject":"Bugfix","Body":""}
{"Hash":{"Long":"e4ed46a23d996e41664a1bca0503ab056cfb4003","Short":"e4ed46a"},"Tree":{"Long":"7ff0e47cbe841be6ba661f399ae0842cd178ce92","Short":"7ff0e47"},"Author":{"Name":"ariarijp","Email":"takuya.arita@gmail.com","Date":"2020-02-11T23:37:40+09:00"},"Committer":{"Name":"ariarijp","Email":"takuya.arita@gmail.com","Date":"2020-02-11T23:46:55+09:00"},"Subject":"Initial commit","Body":""}
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
