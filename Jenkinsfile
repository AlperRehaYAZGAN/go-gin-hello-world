pipeline {
    agent any
    environment {
        GIT_SOURCE_URL = ""
        CLONE_FOLDER = ""
        IMAGE_NAME = ""
        IMAGE_TAG = ""
        DOCKER_CREDENTIAL = ""
        CONTAINER_NAME = ""
        APP_PORT = "9090"
        SSH_KEY_PATH = "" 
        SSH_USERNAME = ""
        SSH_HOST = ""
    }
    stages {
        stage('Clone repository') {
            steps {
                // clone GIT_SOURCE_URL into folder CLONE_FOLDER
                sh 'git clone ${GIT_SOURCE_URL} ${CLONE_FOLDER}'
                cd '${CLONE_FOLDER}'
            }
        }
        stage('Build Image With Go from source') {
            steps {
                sh 'docker build -t $IMAGE_NAME:$IMAGE_VERSION .'
            }
        }
        stage('Docker Login') {
            steps {
                sh 'docker login -u $DOCKER_USERNAME'
            }
        }
        stage('Push Image to Docker Hub with credentials') {
            steps {
                sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
                sh 'docker push $IMAGE_NAME:$IMAGE_VERSION'
            }
        }
        stage('connect deployment server with SSH') {
            steps {
                sh 'ssh -i $SSH_KEY_PATH $SSH_USERNAME@$SSH_HOST'
            }
        }
        stage('In deployment server stop and remove old container') {
            steps {
                sh 'docker stop $CONTAINER_NAME'
                sh 'docker rm $CONTAINER_NAME'
            }
        }
        stage('Run Docker New Container') {
            steps {
                sh 'docker run -d -p $APP_PORT:$APP_PORT $IMAGE_NAME:$IMAGE_VERSION'
                // exit ssh
                sh 'exit'
            }
        }
        stage('Cleanup') {
            steps {
                sh 'cd ..'
                sh 'rm -rf go-gin-hello-world'
            }
        }
    }
}