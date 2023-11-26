build:
	GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bin/bootstrap cmd/typeform-webhook/main.go

zip: build
	zip -j bin/bootstrap.zip bin/bootstrap

lambda-exists:
	aws lambda get-function --function-name typeform-webhook

lambda-create: zip
	aws lambda create-function \
		--function-name typeform-webhook \
		--runtime provided.al2 \
		--handler bootstrap \
		--architectures arm64 \
		--role "$LAMBDA_ROLE_ARN" \
		--zip-file fileb://bin/bootstrap.zip

lambda-update: zip
	aws lambda update-function-code \
		--function-name typeform-webhook \
		--architectures arm64 \
		--zip-file fileb://bin/bootstrap.zip
