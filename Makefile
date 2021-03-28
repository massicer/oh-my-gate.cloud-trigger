REGION = us-central1
GO_RUNTIME = go113
FUNCTION_BASE_SOURCE = .
ENVFILE = .env.yaml

	
deploy:
	gcloud functions deploy Open \
	--source ${FUNCTION_BASE_SOURCE} \
	--runtime ${GO_RUNTIME} \
	--trigger-http \
	--env-vars-file=${ENVFILE} \
	--allow-unauthenticated \
	--region ${REGION}