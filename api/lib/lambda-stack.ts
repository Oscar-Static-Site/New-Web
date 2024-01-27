import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class ApiStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const Function = new lambda.Function(this, 'VisitorFunction', {
      code: lambda.Code.fromAsset('lambda'),
      handler: 'main',
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });

    const gateway = new RestApi(this, 'VisitorAPI', {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*.oscarconer.com"],
        allowMethods: ["GET"],
              },
    });

    const ingtegraion = new LambdaIntegration(Function);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'ApiQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });
  }
}
