# Feed Summarizer

Feed Summarizer sends you a daily email notification summarizing all the updates to your feeds within the last 24 hours.

## Development

Create urlList file for local:

```sh
touch urlList.txt
```

To start the development process, run the main program using the following command in your terminal:

```sh
go run cmd/main.go
```

Then, in a separate terminal window, simulate the daily email dispatch by executing the following command:

```sh
curl localhost:8080
```

This command will trigger the program to send an email summarizing the feed updates. Enjoy the development process!
