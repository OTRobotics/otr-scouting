node {
    withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}", "GOROOT=/usr/local/go"]) {
        env.PATH="${GOPATH}/bin:${GOROOT}/bin:$PATH"
            stage('Pre Test'){
                echo 'Pulling Dependencies'
                sh 'go version'
                sh 'go get .'
            }

            stage('Compile Match Uploader') {
                steps {
                    echo 'MatchUploader Build for All OSs'
                    sh 'cd matchUploader && export GOOS=windows && go build main.go'
                    sh 'cd matchUploader && export GOOS=darwin && go build main.go'
                    sh 'cd matchUploader && export GOOS=linux && go build main.go'
                }
            }
            stage('Deploy') {
                environment {
                    GCP_APPENGINE = credentials('otr-scouting-appengine')
                }
                steps {
                    echo 'Deploying to GCP'
                    sh "/var/jenkins_home/google-cloud-sdk/bin/gcloud auth activate-service-account --key-file ${env.GCP_APPENGINE}"
                    sh 'export GOROOT=/usr/local/go && export GOHOME=/var/jenkins_home/go && export PATH=$GOPATH/bin:$GOROOT/bin:$PATH && cd appengine && gcloud app deploy'
                }
            }
            stage('Upload MatchUploader to GCS') {
                environment {
                    GCS_CREDS = credentials('otr-scouting-web-gcloud')
                }
                steps {
                    sh "/var/jenkins_home/google-cloud-sdk/bin/gcloud auth activate-service-account --key-file ${env.GCS_CREDS}"
                    sh '/var/jenkins_home/google-cloud-sdk/bin/gsutil cp *.pdf gs://staging.otr-scouting.appspot.com/matchUploader/'
                }
            }
    }
}