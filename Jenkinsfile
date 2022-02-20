pipeline {
    agent any
    environment {
        GIT_SOURCE_URL = ""
        CLONE_FOLDER = ""
        IMAGE_NAME = ""
        IMAGE_TAG = ""
        DOCKERHUB_CREDENTIAL = credentials("DOCKERHUB_CREDENTIAL")
    }
    stages {
        stage('Clone repository') {
            steps {
                sh 'git clone ${GIT_SOURCE_URL} ${CLONE_FOLDER}'
                cd '${CLONE_FOLDER}'
            }
        }
        stage('DockerHub Login with credentials') {
            steps {
                sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
            }
        }
        stage('Build Docker Container Image from source') {
            steps {
                sh 'docker build -t $IMAGE_NAME:$IMAGE_TAG .'
            }
        }
        stage('Push Image to Docker Hub') {
            steps {
                sh 'docker push $IMAGE_NAME:$IMAGE_VERSION'
            }
        }
        stage('Cleanup') {
            steps {
                sh 'cd ..'
                sh 'rm -rf ${CLONE_FOLDER}'
            }
        }
    }
}