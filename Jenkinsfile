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
        stage('Clean Workspace') {
            steps {
                clearWs()
            }
        }
        stage('Clone repository') {
            steps {
                sh '''#!/bin/bash -e
                echo "Cloning repository"
                git clone ${GIT_SOURCE_URL} ${CLONE_FOLDER}
                echo "Cloning repository done"
                '''
            }
        }
        stage('DockerHub Login with credentials') {
            steps {
                sh '''#!/bin/bash -e
                echo "DockerHub Login with credentials"
                docker login -u ${DOCKERHUB_CREDENTIAL.username} -p ${DOCKERHUB_CREDENTIAL.password}
                echo "DockerHub Login with credentials done"
                '''
            }
        }
        stage('Build Docker Container Image from source') {
            steps {
                sh '''#!/bin/bash -e
                echo "Build Docker Container Image from source"
                cd ${CLONE_FOLDER}
                docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
                echo "Build Docker Container Image from source done"
                '''
            }
        }
        stage('Push Image to Docker Hub') {
            steps {
                sh '''#!/bin/bash -e
                echo "Push Image to Docker Hub"
                docker push ${IMAGE_NAME}:${IMAGE_TAG}
                echo "Push Image to Docker Hub done"
                '''
            }
        }
        stage('Cleanup') {
            steps {
                sh '''#!/bin/bash -e
                echo "Cleanup"
                cd ..
                rm -rf ${CLONE_FOLDER}
                echo "Cleanup done"
                '''
            }
        }
    }
}