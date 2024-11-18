# Clutta-go Quickstart Documentation
This is a client library built in Go to help users communicate with Clutta. It simplifies the process of setting up a clutta client and sending pulses.

This documentation will help you set up and test clutta-go quickly.

## Prerequisites for Quickstart
- Install Go (https://go.dev/doc/install) and Git on your computer.
- Set up an account on clutta.io and obtain an API key {See future doc on how to get API key}.

## Install the library
To install and use clutta-go in your project, kindly follow the steps listed below:

1. Create a **Personal Access Token** on Github by going to your `Settings > Developer Settings > Tokens`.
   Ensure the token has `repo` and `read:packages` scopes selected. Copy the token and store temporarily.

2. Open your terminal/command prompt and `cd` to your home directory. You can also `cd` to any directory of your choosing.
	On Linux/Mac:
	```
	cd
	```
	On Windows:
	```
	cd %HOMEPATH%
	```
	**Note:** You can also use an IDE like VSCode and store your quickstart code in any folder you choose. From there, you can use the terminal or command prompt inside the IDE to run all your commands.

3. Create a directory to store your quickstart code.
	```
	mkdir test_clutta_go
	cd test_clutta_go
	```

4. Create a go module in the directory.
	```
	go mod init example.com/test_clutta
	```

5. Since the clutta-go repository is currently private, you have to set the `GOPRIVATE` environment variable in your terminal/command prompt.
	```
	go env -w GOPRIVATE="github.com/sefastech/*"
	```

6. Set your github username and access token (the one generated above) as environment variables as shown below.
	```
	export GH_ACCESS_USER = <YOUR_GITHUB_USERNAME>
	export GH_ACCESS_TOKEN = <YOUR GENERATED PERSONAL ACCESS TOKEN>

	```

7. Configure Git to Use Access Tokens for GitHub Authentication.
	```
	git config --global url."https://${GH_ACCESS_USER}:${GH_ACCESS_TOKEN}@github.com".insteadOf "https://github.com"
	```

8. Install clutta-go.
	```
	go get github.com/sefastech/clutta/tree/main/libraries/golang/clients/clutta-sync@latest
	```
	You should have the package installed.

## Example Usage of the Library

1. Create a new file `main.go` in the same directory `test_clutta_go` from above. Copy the code below into the file and fill in the details for your `chain_id, correlation_id, signature_id and source_id`.

	```
	package main

	import (
		cluttaSyncClient "github.com/sefastech/clutta/libraries/golang/clients/clutta-sync"
		"fmt"
	)

	func main(){
		const chain_id="******"
		const signature_id="********"
		const correlation_id="******"
		const source_id="********"
		const status="SUCCESS"
		const status_description="User testing of clutta go"
		const customer_id="********"
		new_client,err:= cluttaSyncClient.SetupClient(customer_id,"sync.clutta.io","443")
		if err!=nil{
			fmt.Println(err)
		}
		new_client.SendPulse(chain_id,signature_id,correlation_id,source_id,status,status_description)
	}
	```

2. Run the file in the terminal/command prompt.
	```
	go run main.go
	```

	You should see a `"Pulse sent successfully"` message.