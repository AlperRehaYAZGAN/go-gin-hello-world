pipeline {
    agent any
    environment {
        GIT_SOURCE_URL = "${params.GIT_SOURCE_URL}"
        CLONE_FOLDER = "${params.CLONE_FOLDER}"
        IMAGE_NAME = "${params.IMAGE_NAME}"
        IMAGE_TAG = "${params.IMAGE_TAG}"
        DOCKERHUB_CREDENTIAL = credentials("DOCKERHUB_CREDENTIAL")
    }
    stages {
        stage('Clone repository') {
            steps {
                sh "git clone ${GIT_SOURCE_URL} ${CLONE_FOLDER} "
                cd "${CLONE_FOLDER}"
            }
        }
        stage('DockerHub Login with credentials') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'DOCKERHUB_CREDENTIAL', usernameVariable: 'DOCKERHUBUSERNAME', passwordVariable: 'DOCKERHUBPASSWORD')]) {
                        sh "docker login -u ${DOCKERHUBUSERNAME} -p ${DOCKERHUBPASSWORD}"
                }
            }
        }
        stage('Build Docker Container Image from source') {
            steps {
                sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG} ."
            }
        }
        stage('Push Image to Docker Hub') {
            steps {
                sh "docker push ${IMAGE_NAME}:${IMAGE_VERSION}"
            }
        }
        stage('Cleanup') {
            steps {
                sh 'cd ..'
                sh "rm -rf ${CLONE_FOLDER}"
            }
        }
    }
}