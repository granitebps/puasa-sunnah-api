pipeline {
    agent any

    parameters {
        string(name: 'BRANCH', defaultValue: 'main', description: 'Branch to deploy')
    }

    environment {
        APP_PATH    = '/var/www/puasa-sunnah-api'
        TMP_ENV_DIR = "/tmp/puasa-sunnah-api-env-${env.BUILD_NUMBER}"
        // Ensure jenkins uses its own SSH key for all git operations
        GIT_SSH_COMMAND = 'ssh -i /var/lib/jenkins/.ssh/id_ed25519 -o StrictHostKeyChecking=no'
    }

    stages {

        stage('Fetch Environment File') {
            steps {
                sh '''
                    mkdir -p ${TMP_ENV_DIR}
                    git clone --depth=1 git@github.com:granitebps/env.git ${TMP_ENV_DIR}
                '''
            }
        }

        stage('Pull') {
            steps {
                sh """
                    cd ${APP_PATH}
                    git fetch origin ${params.BRANCH}
                    git reset --hard origin/${params.BRANCH}
                """
            }
        }

        stage('Build') {
            steps {
                sh """
                    cd ${APP_PATH}
                    make build
                """
            }
        }

        stage('Deploy') {
            steps {
                sh """
                    # Update .env from private env repo
                    cp ${TMP_ENV_DIR}/puasa-sunnah-api/.prod.env ${APP_PATH}/.env

                    # Restart service
                    sudo systemctl restart puasa-sunnah-api
                """
            }
        }
    }

    post {
        always {
            sh 'rm -rf ${TMP_ENV_DIR}'
            cleanWs()

            script {
                def mins     = (currentBuild.duration / 60000) as int
                def secs     = ((currentBuild.duration % 60000) / 1000) as int
                def duration = String.format('%02d:%02d', mins, secs)

                emailext(
                    to: 'granitebagas28@gmail.com',
                    subject: "Jenkins ${currentBuild.currentResult}: ${env.JOB_NAME} #${env.BUILD_NUMBER}",
                    body: """
                        <p><b>Status:</b> ${currentBuild.currentResult}</p>
                        <p><b>Branch:</b> ${params.BRANCH}</p>
                        <p><b>Build:</b> <a href="${env.BUILD_URL}">${currentBuild.fullDisplayName}</a></p>
                        <p><b>Duration:</b> ${duration} (mm:ss)</p>
                    """,
                    mimeType: 'text/html'
                )
            }
        }
    }
}
