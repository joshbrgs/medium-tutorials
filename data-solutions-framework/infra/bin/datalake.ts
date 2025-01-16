#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { CICDPipelineStack } from '../lib/cicd-stack';

const cicdEnvironment: cdk.Environment = {
   region: process.env.CDK_DEFAULT_REGION,
   account: process.env.CDK_DEFAULT_ACCOUNT,
};

const app = new cdk.App();
new CICDPipelineStack(app, 'CICDPipeline',{ env: cicdEnvironment });

