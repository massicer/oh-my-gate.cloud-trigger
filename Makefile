deploy:
	gcloud functions deploy Open \
	--env-vars-file=env.yaml \
	--runtime go111 --trigger-http --allow-unauthenticated