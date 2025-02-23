#!/usr/bin/env bash

if [ -z "$GITHUB_TOKEN" ]; then
	GITHUB_TOKEN=""
fi

set -o errexit
set -o nounset
set -o pipefail

RELEASE_TAG="$(git describe --tag)"
RELEASE_ID=0

# caminogo root folder
CAMINOGO_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Authentication
AUTH_HEADER=""
# check if GITHUB_TOKEN is set
if [ -n "$GITHUB_TOKEN" ]; then
	AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"
fi

publish () {
	# check if GITHUB_TOKEN is set
	if [ -z "$AUTH_HEADER" ]; then
		return 0
	fi
	FILENAME=$(basename "$1")
	UPLOAD_URL="https://uploads.github.com/repos/${GITHUB_REPOSITORY}/releases/${RELEASE_ID}/assets?name=${FILENAME}"
	echo "upload url: ${UPLOAD_URL}"

	# create a temp file for upload output
	logOut=$(mktemp)
	# Upload the artifact - capturing HTTP response-code in our output file.
	response=$(curl \
		-sSL \
		-XPOST \
		-H "${AUTH_HEADER}" \
		--upload-file "$1" \
		--header "Content-Type:application/octet-stream" \
		--write-out "%{http_code}" \
		--output "$logOut" \
		"${UPLOAD_URL}")

   curl_success=$?
   if [ $curl_success -ne 0 ]; then
        echo "err: curl command failed!!!"
		    rm "$logOut"
        return 1
    fi

	cat "$logOut" && echo ""
	rm "$logOut"

	if [ "$response" -ge 400 ]; then
		echo "err: upload not successful ($response)!!!"
 		return 1
	fi
}

if [ -n "$AUTH_HEADER" ]; then
	GH_API="https://api.github.com/repos/${GITHUB_REPOSITORY}"

	# release the version
	response="$(curl  -sH "${AUTH_HEADER}" --data "{\"tag_name\":\"$RELEASE_TAG\", \"draft\":true, \"generate_release_notes\":true}" "$GH_API/releases")"

	# extract id out of response
	eval "$(echo "$response" | grep -m 1 "id.:" | grep -w id | tr : = | tr -cd '[[:alnum:]]=')"
	[ "$id" ] || { echo "Error: Failed to get release id for tag: $RELEASE_TAG"; printf "%s\n" "$response" >&2; exit 1; }
	RELEASE_ID=$id
fi

echo "Building release OS=linux and ARCH=amd64 using GOAMD64 V2 for caminogo version $RELEASE_ID"
rm -rf "$CAMINOGO_PATH"/build/*

DEST_PATH=$CAMINOGO_PATH/dist/
ARCHIVE_PATH=caminogo-$RELEASE_TAG
# prepare a fresh dist folder
rm -rf "$DEST_PATH" && mkdir -p "$DEST_PATH"

# build executables into build dir
GOOS=linux GOARCH=amd64 GOAMD64=v2 "$CAMINOGO_PATH"/scripts/build.sh
# build tools into build dir
GOOS=linux GOARCH=amd64 GOAMD64=v2 "$CAMINOGO_PATH"/scripts/build_tools.sh
# copy the license file
cp "$CAMINOGO_PATH"/LICENSE "$CAMINOGO_PATH"/build

# create the package
echo "building artifact"
ARTIFACT=$DEST_PATH/caminogo-linux-amd64-$RELEASE_TAG.tar.gz
tar -czf "$ARTIFACT" -C "$CAMINOGO_PATH" build --transform "s,build,$ARCHIVE_PATH,"
# publish the newly generated file
publish "$ARTIFACT"
