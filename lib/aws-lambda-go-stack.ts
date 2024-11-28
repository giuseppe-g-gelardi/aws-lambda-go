import { Stack, type StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { LambdaIntegration, RestApi } from "aws-cdk-lib/aws-apigateway";
import * as path from "path";

export class AwsLambdaGoStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // define api gateway with cors
    const api = new RestApi(this, "myGateway", {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*"],
        allowMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
      },
    });

    // define the lambda function
    const myFunction = new lambda.Function(this, "MyLambda", {
      code: lambda.Code.fromAsset(path.join(__dirname, "../lambdas/Get")),
      handler: "main",
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });
    // define the integration
    const integration = new LambdaIntegration(myFunction);

    const helloLambda = new lambda.Function(this, "HelloLambda", {
      code: lambda.Code.fromAsset(path.join(__dirname, "../lambdas/Hello/")),
      handler: "main",
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });
    const helloIntegration = new LambdaIntegration(helloLambda);

    const goodbyeLambda = new lambda.Function(this, "GoodbyeLambda", {
      code: lambda.Code.fromAsset(path.join(__dirname, "../lambdas/Goodbye/")),
      handler: "main",
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });
    const goodbyeIntegration = new LambdaIntegration(goodbyeLambda);

    const numsDouble = new lambda.Function(this, "NumsdoubleLambda", {
      code: lambda.Code.fromAsset(
        path.join(__dirname, "../lambdas/Numsdouble/"),
      ),
      handler: "main",
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });
    const numsDoubleIntegration = new LambdaIntegration(numsDouble);

    // define the resource and method associated with the lambda
    const helloResource = api.root.addResource("hello");
    const goodbyeResource = api.root.addResource("goodbye");
    const numsDoubleResource = api.root.addResource("numsdouble");

    // add the method to the resource
    helloResource.addMethod("POST", integration, {
      apiKeyRequired: true,
    });

    helloResource.addMethod("GET", helloIntegration, {
      apiKeyRequired: true,
    });

    goodbyeResource.addMethod("POST", goodbyeIntegration, {
      apiKeyRequired: true,
    });

    numsDoubleResource.addMethod("POST", numsDoubleIntegration, {
      apiKeyRequired: true,
    });

    const apiKey = api.addApiKey("ApiKey", {
      apiKeyName: "helloApiKey",
    });

    // you def want to set this up so you cant get ddos'd
    const usagePlan = api.addUsagePlan("HelloUsagePlan", {
      name: "HelloUsagePlan",
      throttle: {
        burstLimit: 5,
        rateLimit: 10,
      },
    });

    usagePlan.addApiKey(apiKey);
  }
}
