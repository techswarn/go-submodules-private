# # exec git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
# # exec git clone "https://github.com/techswarn/tempconv.git"
# exec git clone "git+https://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:x-oauth-basic@github.com/techswarn/tempconv.git"

#!/bin/env bash

# Do build locally with build.sh file in function directory
# so that we can use the machine's configuration to download
# private Go dependencies.
#
# Because we're building locally, we need to do everything the
# remote build would normally do for us. So after downloading
# private dependencies, we need to use docker with the official
# OpenWhisk Go runtime to compile our application.

# Cleanup
rm -r vendor 2> /dev/null
rm __deployer__.zip 2> /dev/null

# Vendor dependencies so that they're all in the current
# directory and can be included in the ZIP file that is used
# in the next step.
go mod vendor

# Version of Docker image used must match version of Go runtime
# used with DO Functions. Save as a particular filename recognized
# by doctl as "the artifact to be built" so that doctl skips the
# build entirely and uploads it as is.
zip -r - * | docker run openwhisk/action-golang-v1.17 -compile main >__deployer__.zip