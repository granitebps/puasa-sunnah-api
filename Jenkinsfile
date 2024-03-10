pipeline {
    agent any

    tools {
        // Specify the desired Go version defined in the Global Tool Configuration
        go '1.21'
    }

    environment {
        GOPATH = "${WORKSPACE}/go"
        PATH = "${GOPATH}/bin:${PATH}"
        CGO_ENABLED = 0  // Disable CGO for static binary
    }

    options {
        // Skip checkout for existing changes to speed up the build
        skipDefaultCheckout true
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout the source code from your version control system (e.g., Git)
                checkout scm
            }
        }

        stage('Dependencies') {
            steps {
                // Fetch dependencies using a vendoring tool like Go Modules
                sh 'go mod download'
            }
        }

        stage('Build') {
            steps {
                // Set up Go environment
                script {
                    sh 'mkdir -p ${GOPATH}/src'
                    sh 'ln -s ${WORKSPACE} ${GOPATH}/src/puasa-sunnah-api'
                }

                // Build the Golang application with optimizations and set the output filename
                sh 'go build -o puasa-sunnah-api -ldflags="-s -w" ./...'
            }
        }

        stage('Test') {
            steps {
                // Run tests
                sh 'go test ./...'
            }
        }
    }

    post {
        success {
            // Notify success and trigger other downstream jobs or processes
            emailext body: "Build successful! Version ${env.BUILD_NUMBER}",
                     subject: "Success: Build ${env.BUILD_NUMBER}",
                     to: "granitebagas28@gmail.com"
        }

        failure {
            // Notify failure and take appropriate actions
            emailext body: "Build failed for version ${env.BUILD_NUMBER}",
                     subject: "Failure: Build ${env.BUILD_NUMBER}",
                     to: "granitebagas28@gmail.com"
        }
    }
}
