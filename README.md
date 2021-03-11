# gojsontoenv
Take a Json file and output compatible Shell Env variables

## Install

Grad the version you need from the [Github Release page](https://github.com/prune998/gojsontoenv/releases)

ex:

```shell
export GOJSONTOENV_VERSION=1.0.0
export OS=Linux
export ARCH=x86_64
curl --location --output /tmp/gojsontoenv.tar.gz \
	"https://github.com/prune998/gojsontoenv/releases/download/v${GOJSONTOENV_VERSION}/gojsontoenv_${GOJSONTOENV_VERSION}_${OS}_${ARCH}.tar.gz" \
	&& tar -xzf /tmp/gojsontoenv.tar.gz -C /usr/local/bin gojsontoenv \
	&& chmod +x /usr/local/bin/gojsontoenv \
	&& rm /tmp/gojsontoenv.tar.gz
```

## Usage

```shell
./gojsontoenv -h

Usage of ./gojsontoenv:
  -input="": path to the input JSON file, read from stdin if none provided
  -output="export": output format, one of 'export' (default) or 'vars'
  -version=false: Show version and quit
```
By default, `gojsontoenv` will read `<stdin>` so you can just pipe your json file into it:

```shell
cat config.json | ./gojsontoenv

export MY_ENV_1="foo"
export MY_ENV_2="bar"
export MY_INT="42"
export MY_FLOAT="12.34"
export MY_BOOL="true"
```

Or use the `--input` argument to read a file:

```shell
./gojsontoenv --input config.json
```

The `--output` is used to change the output format:

```shell
./gojsontoenv --input config.json --output vars

MY_ENV_1="foo"
MY_ENV_2="bar"
MY_INT="42"
MY_FLOAT="12.34"
MY_BOOL="true"
```

### Using with Terraform

This tool can be used with Terraform `output` variables.

To do so, define some output variables in your TF scripts:

```hcl
output "k8s_core_vars" {
  value = {
    "CLUSTER_NAME"                = local.cluster_name,
    "CLUSTER_REGION"              = var.region,
    "EXTERNAL_DNS_ROLE_ARN"       = aws_iam_role.k8s_external_dns.arn
    "CLUSTER_AUTOSCALER_ROLE_ARN" = aws_iam_role.cluster_autoscaler.arn
    "ALB_INGRESS_ROLE_ARN"        = aws_iam_role.alb_ingress.arn
  }
}
```

After your TF script is applied, you can use the TF `output` command to dump those values:

```shell
terraform output -json k8s_core_vars | jq '.'

{
  "ALB_INGRESS_ROLE_ARN": "arn:aws:iam::123456789:role/my-alb-ingress-role",
  "AWS_CLOUDWATCH_ROLE_ARN": "arn:aws:iam::123456789:role/my-amazon-cloudwatch",
  "CLUSTER_AUTOSCALER_ROLE_ARN": "arn:aws:iam::123456789:role/my-cluster-autoscaler",
  "CLUSTER_NAME": "my-eks-cluster",
  "CLUSTER_REGION": "us-east-1",
  "EXTERNAL_DNS_ROLE_ARN": "arn:aws:iam::123456789:role/my-external-dns"
}

# dump values into a file
terraform output -json k8s_core_vars > k8s_core_vars.json
```

You can the use this app to transform this JSON file into a file usable in a shell:

```shell
./gojsontoenv --input  k8s_core_vars.json

export ALB_INGRESS_ROLE_ARN="arn:aws:iam::123456789:role/my-alb-ingress-role"
export AWS_CLOUDWATCH_ROLE_ARN="arn:aws:iam::123456789:role/my-amazon-cloudwatch"
export CLUSTER_AUTOSCALER_ROLE_ARN="arn:aws:iam::123456789:role/my-cluster-autoscaler"
export CLUSTER_NAME="my-eks-cluster"
export CLUSTER_REGION="us-east-1"
export EXTERNAL_DNS_ROLE_ARN="arn:aws:iam::123456789:role/my-external-dns"
```

To set those variables inside your current shell (for a CI/CD):

```shell
eval $(./gojsontoenv --input  k8s_core_vars.json)
```

Note that the content of the file may be displayed in your shell if you set something like `set -x`.

If you don't want the content to be displayed in your CI logs, source the file instead:

```shell
./gojsontoenv --input k8s_core_vars.json > k8s_core_vars.sh && source k8s_core_vars.sh
```

### Docker

You can run the app from a docker container. Just mount your `config` file:

```shell
docker run --rm -v $(pwd)/config.json:/config.json  prune/gojsontoenv
docker run --rm -v $(pwd)/config.json:/config.json  prune/gojsontoenv --output=vars
```
