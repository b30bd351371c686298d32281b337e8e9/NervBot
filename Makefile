tarball:
	tar -czf credentials.tar.gz gcp-creds.json .env
	gpg -c credentials.tar.gz
	rm credentials.tar.gz
