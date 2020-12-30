deploy:
	gcloud functions deploy Open \
	# --source=./src \
	--env-vars-file=env.yaml \
	--runtime go111 --trigger-http --allow-unauthenticated