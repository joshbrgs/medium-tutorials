#!/bin/bash
echo $STEP_FUNCTION_ARN
EXECUTION_ARN=$(aws stepfunctions start-execution --state-machine-arn $STEP_FUNCTION_ARN --debug | jq -r '.executionArn')
while true
do
    STATUS=$(aws stepfunctions describe-execution --execution-arn $EXECUTION_ARN --debug | jq -r '.status')
    if [ $STATUS = "SUCCEEDED" ]; then
        exit 0 
    elif [ $STATUS = "FAILED" ] || [ $STATUS = "TIMED_OUT" ] || [ $STATUS = "ABORTED" ]; then
        exit 1 
    else 
        sleep 10
        continue
    fi
done
