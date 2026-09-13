# ACME DNS Providers

When an account in the ACME secret backend is configured to use the DNS-01
challenge, credentials to access and update the DNS records must be configured.
Set the account's `provider` to the code name shown in parentheses after each
provider, and pass the credentials in `provider_configuration`.

~> **WARNING:** The DNS-01 challenge for a given domain requires to update the
   TXT record at `_acme-challenge.<YOUR_DOMAIN>`. If your DNS provider supports it, use
   credentials with permissions restricted to these records to improve security.


[//]: # (The rest of this file was produced from the lego v5.4.1 documentation under MIT license)
[//]: # (with the following program, run in lego's checkout:)
[//]: # (for provider in sorted(glob('./providers/dns/**/*.toml', recursive=True)):)
[//]: # (    with open(provider, 'rb') as f:)
[//]: # (        data = tomllib.load(f))
[//]: # (    print(f"## {data['Name']} (`{data['Code']}`)"))
[//]: # (    if data.get('Description', '').strip():)
[//]: # (        print(data['Description'].strip()))
[//]: # (    if 'Additional' in data:)
[//]: # (        print(data['Additional'].strip().replace('##', '###')))
[//]: # (    conf = data.get('Configuration', {}))
[//]: # (    if 'Credentials' in conf:)
[//]: # (        print('### Credentials'))
[//]: # (        for k, v in conf['Credentials'].items():)
[//]: # (            print(f"  - `{k}`: {v}"))
[//]: # (    if 'Additional' in conf:)
[//]: # (        print('\n### Additional configuration'))
[//]: # (        for k, v in conf['Additional'].items():)
[//]: # (            print(f"  - `{k}`: {v}"))
[//]: # (    print())
[//]: # (    print())

## Abion (`abion`)
### Credentials
  - `ABION_API_KEY`: API key

### Additional configuration
  - `ABION_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ABION_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ABION_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ABION_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Joohoi's ACME-DNS (`acmedns`)
### Credentials
  - `ACME_DNS_API_BASE`: The ACME-DNS API address
  - `ACME_DNS_STORAGE_PATH`: The ACME-DNS JSON account data file. A per-domain account will be registered/persisted to this file and used for TXT updates.
  - `ACME_DNS_STORAGE_BASE_URL`: The ACME-DNS JSON account data server.

### Additional configuration
  - `ACME_DNS_ALLOWLIST`: Source networks using CIDR notation (multiple values should be separated with a comma).


## Active24 (`active24`)
### Credentials
  - `ACTIVE24_API_KEY`: API key
  - `ACTIVE24_SECRET`: Secret

### Additional configuration
  - `ACTIVE24_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ACTIVE24_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ACTIVE24_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ACTIVE24_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Alibaba Cloud DNS (`alidns`)
### Credentials
  - `ALICLOUD_RAM_ROLE`: Your instance RAM role (https://www.alibabacloud.com/help/en/ecs/user-guide/attach-an-instance-ram-role-to-an-ecs-instance)
  - `ALICLOUD_ACCESS_KEY`: Access key ID
  - `ALICLOUD_SECRET_KEY`: Access Key secret
  - `ALICLOUD_SECURITY_TOKEN`: STS Security Token (optional)

### Additional configuration
  - `ALICLOUD_REGION_ID`: Region ID (Default: cn-hangzhou)
  - `ALICLOUD_LINE`: Line (Default: default)
  - `ALICLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ALICLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ALICLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `ALICLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## AlibabaCloud ESA (`aliesa`)
### Credentials
  - `ALIESA_RAM_ROLE`: Your instance RAM role (https://www.alibabacloud.com/help/en/ecs/user-guide/attach-an-instance-ram-role-to-an-ecs-instance)
  - `ALIESA_ACCESS_KEY`: Access key ID
  - `ALIESA_SECRET_KEY`: Access Key secret
  - `ALIESA_SECURITY_TOKEN`: STS Security Token (optional)

### Additional configuration
  - `ALIESA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ALIESA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ALIESA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ALIESA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## all-inkl (`allinkl`)
### Credentials
  - `ALL_INKL_LOGIN`: KAS login
  - `ALL_INKL_PASSWORD`: KAS password

### Additional configuration
  - `ALL_INKL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ALL_INKL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ALL_INKL_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Alwaysdata (`alwaysdata`)
### Credentials
  - `ALWAYSDATA_API_KEY`: API Key

### Additional configuration
  - `ALWAYSDATA_ACCOUNT`: Account name
  - `ALWAYSDATA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ALWAYSDATA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ALWAYSDATA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ALWAYSDATA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Anexia CloudDNS (`anexia`)
### Description

You need to create an API token in the [Anexia Engine](https://engine.anexia-it.com/).

The token must have permissions to manage DNS zones and records.
### Credentials
  - `ANEXIA_TOKEN`: API token for Anexia Engine

### Additional configuration
  - `ANEXIA_API_URL`: API endpoint URL (default: https://engine.anexia-it.com)
  - `ANEXIA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ANEXIA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `ANEXIA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `ANEXIA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ArtFiles (`artfiles`)
### Credentials
  - `ARTFILES_USERNAME`: API username
  - `ARTFILES_PASSWORD`: API password

### Additional configuration
  - `ARTFILES_SERVER_NAME`: Your server name (Default: dcp)
  - `ARTFILES_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ARTFILES_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 360)
  - `ARTFILES_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ArvanCloud (`arvancloud`)
### Credentials
  - `ARVANCLOUD_API_KEY`: API key

### Additional configuration
  - `ARVANCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ARVANCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `ARVANCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `ARVANCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Aurora DNS (`auroradns`)
### Credentials
  - `AURORA_API_KEY`: API key or username to used
  - `AURORA_SECRET`: Secret password to be used

### Additional configuration
  - `AURORA_ENDPOINT`: API endpoint URL
  - `AURORA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `AURORA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `AURORA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)


## Autodns (`autodns`)
### Credentials
  - `AUTODNS_API_USER`: Username
  - `AUTODNS_API_PASSWORD`: User Password

### Additional configuration
  - `AUTODNS_ENDPOINT`: API endpoint URL, defaults to https://api.autodns.com/v1/
  - `AUTODNS_CONTEXT`: API context (4 for production, 1 for testing. Defaults to 4)
  - `AUTODNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `AUTODNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `AUTODNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `AUTODNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Axelname (`axelname`)
### Credentials
  - `AXELNAME_NICKNAME`: Account nickname
  - `AXELNAME_TOKEN`: API token

### Additional configuration
  - `AXELNAME_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `AXELNAME_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `AXELNAME_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `AXELNAME_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Azion (`azion`)
### Credentials
  - `AZION_PERSONAL_TOKEN`: Your Azion personal token.

### Additional configuration
  - `AZION_PAGE_SIZE`: The page size for the API request (Default: 50)
  - `AZION_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `AZION_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `AZION_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `AZION_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Azure DNS (`azuredns`)
### Description

Several authentication methods can be used to authenticate against Azure DNS API.

#### Default Azure Credentials (default option)

Default Azure Credentials automatically detects in the following locations and prioritized in the following order:

1. Environment variables for client secret: `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_SECRET`
2. Environment variables for client certificate: `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_CERTIFICATE_PATH`
3. Workload identity for resources hosted in Azure environment (see below)
4. Shared credentials (defaults to `~/.azure` folder), used by Azure CLI

Link:
- [Azure Authentication](https://learn.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication)

#### Environment variables

###### Service Discovery

Lego automatically finds all visible Azure (private) DNS zones using [Azure ResourceGraph query](https://learn.microsoft.com/en-us/azure/governance/resource-graph/).
This can be limited by specifying environment variable `AZURE_SUBSCRIPTION_ID` and/or `AZURE_RESOURCE_GROUP` which limits the
DNS zones to only a subscription or to one resourceGroup.

Additionally environment variable `AZURE_SERVICEDISCOVERY_FILTER` can be used to filter DNS zones with an addition Kusto filter eg:

```
resources
| where type =~ "microsoft.network/dnszones"
| ${AZURE_SERVICEDISCOVERY_FILTER}
| project subscriptionId, resourceGroup, name
```


###### Client secret

The Azure Credentials can be configured using the following environment variables:
* AZURE_CLIENT_ID = "Client ID"
* AZURE_CLIENT_SECRET = "Client secret"
* AZURE_TENANT_ID = "Tenant ID"

This authentication method can be specifically used by setting the `AZURE_AUTH_METHOD` environment variable to `env`.

###### Client certificate

The Azure Credentials can be configured using the following environment variables:
* AZURE_CLIENT_ID = "Client ID"
* AZURE_CLIENT_CERTIFICATE_PATH = "Client certificate path"
* AZURE_TENANT_ID = "Tenant ID"

This authentication method can be specifically used by setting the `AZURE_AUTH_METHOD` environment variable to `env`.

#### Workload identity

Workload identity allows workloads running Azure Kubernetes Services (AKS) clusters to authenticate as an Azure AD application identity using federated credentials.

This must be configured in kubernetes workload deployment in one hand and on the Azure AD application registration in the other hand.

Here is a summary of the steps to follow to use it :
* create a `ServiceAccount` resource, add following annotations to reference the targeted Azure AD application registration : `azure.workload.identity/client-id` and `azure.workload.identity/tenant-id`.
* on the `Deployment` resource you must reference the previous `ServiceAccount` and add the following label : `azure.workload.identity/use: "true"`.
* create a federated credentials of type `Kubernetes accessing Azure resources`, add the cluster issuer URL  and add the namespace and name of your kubernetes service account.

Link :
- [Azure AD Workload identity](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html)

This authentication method can be specifically used by setting the `AZURE_AUTH_METHOD` environment variable to `wli`.

#### Azure Managed Identity

###### Azure Managed Identity (with Azure workload)

The Azure Managed Identity service allows linking Azure AD identities to Azure resources, without needing to manually manage client IDs and secrets.

Workloads with a Managed Identity can manage their own certificates, with permissions on specific domain names set using IAM assignments.
For this to work, the Managed Identity requires the **Reader** role on the target DNS Zone,
and the **DNS Zone Contributor** on the relevant `_acme-challenge` TXT records.

For example, to allow a Managed Identity to create a certificate for "fw01.lab.example.com", using Azure CLI:

```bash
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
export AZURE_RESOURCE_GROUP="rg1"
export SERVICE_PRINCIPAL_ID="00000000-0000-0000-0000-000000000000"

export AZURE_DNS_ZONE="lab.example.com"
export AZ_HOSTNAME="fw01"
export AZ_RECORD_SET="_acme-challenge.${AZ_HOSTNAME}"

az role assignment create \
--assignee "${SERVICE_PRINCIPAL_ID}" \
--role "Reader" \
--scope "/subscriptions/${AZURE_SUBSCRIPTION_ID}/resourceGroups/${AZURE_RESOURCE_GROUP}/providers/Microsoft.Network/dnszones/${AZURE_DNS_ZONE}"

az role assignment create \
--assignee "${SERVICE_PRINCIPAL_ID}" \
--role "DNS Zone Contributor" \
--scope "/subscriptions/${AZURE_SUBSCRIPTION_ID}/resourceGroups/${AZURE_RESOURCE_GROUP}/providers/Microsoft.Network/dnszones/${AZURE_DNS_ZONE}/TXT/${AZ_RECORD_SET}"
```

A timeout wrapper is configured for this authentication method.
The duration can be configured by setting the `AZURE_AUTH_MSI_TIMEOUT`.
The default timeout is 2 seconds.
This authentication method can be specifically used by setting the `AZURE_AUTH_METHOD` environment variable to `msi`.

###### Azure Managed Identity (with Azure Arc)

The Azure Arc agent provides the ability to use a Managed Identity on resources hosted outside of Azure
(such as on-prem virtual machines, or VMs in another cloud provider).

While the upstream `azidentity` SDK will try to automatically identify and use the Azure Arc metadata service,
if you get `azuredns: DefaultAzureCredential: failed to acquire a token.` error messages,
you may need to set the environment variables:
* `IMDS_ENDPOINT=http://localhost:40342`
* `IDENTITY_ENDPOINT=http://localhost:40342/metadata/identity/oauth2/token`

A timeout wrapper is configured for this authentication method.
The duration can be configured by setting the `AZURE_AUTH_MSI_TIMEOUT`.
The default timeout is 2 seconds.
This authentication method can be specifically used by setting the `AZURE_AUTH_METHOD` environment variable to `msi`.

#### Azure CLI

The Azure CLI is a command-line tool provided by Microsoft to interact with Azure resources.
It provides an easy way to authenticate by simply running `az login` command.
The generated token will be cached by default in the `~/.azure` folder.

This authentication method can be specifically used by setting the `AZURE_AUTH_METHOD` environment variable to `cli`.

#### Open ID Connect

Open ID Connect is a mechanism that establish a trust relationship between a running environment and the Azure AD identity provider.
It can be enabled by setting the `AZURE_AUTH_METHOD` environment variable to `oidc`.

#### Azure DevOps Pipelines

It can be enabled by setting the `AZURE_AUTH_METHOD` environment variable to `pipeline`.
### Credentials
  - `AZURE_CLIENT_ID`: Client ID
  - `AZURE_CLIENT_SECRET`: Client secret
  - `AZURE_TENANT_ID`: Tenant ID
  - `AZURE_CLIENT_CERTIFICATE_PATH`: Client certificate path

### Additional configuration
  - `AZURE_ENVIRONMENT`: Azure environment, one of: public, usgovernment, and china
  - `AZURE_SUBSCRIPTION_ID`: DNS zone subscription ID
  - `AZURE_RESOURCE_GROUP`: DNS zone resource group
  - `AZURE_SERVICEDISCOVERY_FILTER`: Advanced ServiceDiscovery filter using Kusto query condition
  - `AZURE_PRIVATE_ZONE`: Set to true to use Azure Private DNS Zones and not public
  - `AZURE_ZONE_NAME`: Zone name to use inside Azure DNS service to add the TXT record in
  - `AZURE_AUTH_METHOD`: Specify which authentication method to use
  - `AZURE_AUTH_MSI_TIMEOUT`: Managed Identity timeout duration
  - `AZURE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `AZURE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `AZURE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)


## Baidu Cloud (`baiducloud`)
### Credentials
  - `BAIDUCLOUD_ACCESS_KEY_ID`: Access key
  - `BAIDUCLOUD_SECRET_ACCESS_KEY`: Secret access key

### Additional configuration
  - `BAIDUCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BAIDUCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `BAIDUCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)


## Beget.com (`beget`)
### Credentials
  - `BEGET_USERNAME`: API username
  - `BEGET_PASSWORD`: API password

### Additional configuration
  - `BEGET_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 30)
  - `BEGET_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `BEGET_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `BEGET_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Binary Lane (`binarylane`)
### Credentials
  - `BINARYLANE_API_TOKEN`: API token

### Additional configuration
  - `BINARYLANE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BINARYLANE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `BINARYLANE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `BINARYLANE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Bindman (`bindman`)
### Credentials
  - `BINDMAN_MANAGER_ADDRESS`: The server URL, should have scheme, hostname, and port (if required) of the Bindman-DNS Manager server

### Additional configuration
  - `BINDMAN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BINDMAN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `BINDMAN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)


## Bluecat (`bluecat`)
### Credentials
  - `BLUECAT_SERVER_URL`: The server URL, should have scheme, hostname, and port (if required) of the authoritative Bluecat BAM serve
  - `BLUECAT_USER_NAME`: API username
  - `BLUECAT_PASSWORD`: API password
  - `BLUECAT_CONFIG_NAME`: Configuration name
  - `BLUECAT_DNS_VIEW`: External DNS View Name

### Additional configuration
  - `BLUECAT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BLUECAT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `BLUECAT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `BLUECAT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `BLUECAT_SKIP_DEPLOY`: Skip deployments


## Bluecat v2 (`bluecatv2`)
### Credentials
  - `BLUECATV2_SERVER_URL`: The server URL: it should have a scheme, hostname, and port (if required) of the authoritative Bluecat BAM serve
  - `BLUECATV2_USERNAME`: API username
  - `BLUECATV2_PASSWORD`: API password
  - `BLUECATV2_CONFIG_NAME`: Configuration name
  - `BLUECATV2_VIEW_NAME`: DNS View Name

### Additional configuration
  - `BLUECATV2_SKIP_DEPLOY`: Skip quick deployments
  - `BLUECATV2_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BLUECATV2_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `BLUECATV2_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `BLUECATV2_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## BookMyName (`bookmyname`)
### Credentials
  - `BOOKMYNAME_USERNAME`: Username
  - `BOOKMYNAME_PASSWORD`: Password

### Additional configuration
  - `BOOKMYNAME_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BOOKMYNAME_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `BOOKMYNAME_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `BOOKMYNAME_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Bunny (`bunny`)
### Credentials
  - `BUNNY_API_KEY`: API key

### Additional configuration
  - `BUNNY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `BUNNY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `BUNNY_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `BUNNY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Checkdomain (`checkdomain`)
### Credentials
  - `CHECKDOMAIN_TOKEN`: API token

### Additional configuration
  - `CHECKDOMAIN_ENDPOINT`: API endpoint URL, defaults to https://api.checkdomain.de
  - `CHECKDOMAIN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `CHECKDOMAIN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 300)
  - `CHECKDOMAIN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 7)
  - `CHECKDOMAIN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Civo (`civo`)
### Credentials
  - `CIVO_TOKEN`: Authentication token

### Additional configuration
  - `CIVO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 30)
  - `CIVO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `CIVO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `CIVO_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## CloudDNS (`clouddns`)
### Credentials
  - `CLOUDDNS_CLIENT_ID`: Client ID
  - `CLOUDDNS_EMAIL`: Account email
  - `CLOUDDNS_PASSWORD`: Account password

### Additional configuration
  - `CLOUDDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `CLOUDDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `CLOUDDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `CLOUDDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Cloudflare (`cloudflare`)
### Description

You may use `CF_API_EMAIL` and `CF_API_KEY` to authenticate, or `CF_DNS_API_TOKEN`, or `CF_DNS_API_TOKEN` and `CF_ZONE_API_TOKEN`.

#### API keys

If using API keys (`CF_API_EMAIL` and `CF_API_KEY`), the Global API Key needs to be used, not the Origin CA Key.

Please be aware, that this in principle allows Lego to read and change *everything* related to this account.

#### API tokens

With API tokens (`CF_DNS_API_TOKEN`, and optionally `CF_ZONE_API_TOKEN`),
very specific access can be granted to your resources at Cloudflare.
See this [Cloudflare announcement](https://blog.cloudflare.com/api-tokens-general-availability/) for details.

The main resources Lego cares for are the DNS entries for your Zones.
It also needs to resolve a domain name to an internal Zone ID in order to manipulate DNS entries.

Hence, you should create an API token with the following permissions:

* Zone / Zone / Read
* Zone / DNS / Edit

You also need to scope the access to all your domains for this to work.
Then pass the API token as `CF_DNS_API_TOKEN` to Lego.

**Alternatively,** if you prefer a more strict set of privileges,
you can split the access tokens:

* Create one with *Zone / Zone / Read* permissions and scope it to all your zones or just the individual zone you need to edit.
  This is needed to resolve domain names to Zone IDs and can be shared among multiple Lego installations.
  Pass this API token as `CF_ZONE_API_TOKEN` to Lego.
* Create another API token with *Zone / DNS / Edit* permissions and set the scope to the domains you want to manage with a single Lego installation.
  Pass this token as `CF_DNS_API_TOKEN` to Lego.
* Repeat the previous step for each host you want to run Lego on.
* It is possible to use the same api token for both variables if it is given `Zone:Read` and `DNS:Edit` permission for the zone.

This "paranoid" setup is mainly interesting for users who manage many zones/domains with a single Cloudflare account.
It follows the principle of least privilege and limits the possible damage, should one of the hosts become compromised.
### Credentials
  - `CF_API_EMAIL`: Account email
  - `CF_API_KEY`: API key
  - `CF_DNS_API_TOKEN`: API token with DNS:Edit permission (since v3.1.0)
  - `CF_ZONE_API_TOKEN`: API token with Zone:Read permission (since v3.1.0)
  - `CLOUDFLARE_EMAIL`: Alias to CF_API_EMAIL
  - `CLOUDFLARE_API_KEY`: Alias to CF_API_KEY
  - `CLOUDFLARE_DNS_API_TOKEN`: Alias to CF_DNS_API_TOKEN
  - `CLOUDFLARE_ZONE_API_TOKEN`: Alias to CF_ZONE_API_TOKEN

### Additional configuration
  - `CLOUDFLARE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CLOUDFLARE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `CLOUDFLARE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `CLOUDFLARE_HTTP_TIMEOUT`: API request timeout in seconds (Default: )
  - `CLOUDFLARE_BASE_URL`: API base URL (Default: https://api.cloudflare.com/client/v4)


## ClouDNS (`cloudns`)
### Credentials
  - `CLOUDNS_AUTH_ID`: The API user ID
  - `CLOUDNS_AUTH_PASSWORD`: The password for API user ID

### Additional configuration
  - `CLOUDNS_SUB_AUTH_ID`: The API sub user ID
  - `CLOUDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `CLOUDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 180)
  - `CLOUDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `CLOUDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Cloud.ru (`cloudru`)
### Credentials
  - `CLOUDRU_SERVICE_INSTANCE_ID`: Service Instance ID (parentId)
  - `CLOUDRU_KEY_ID`: Key ID (login)
  - `CLOUDRU_SECRET`: Key Secret

### Additional configuration
  - `CLOUDRU_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `CLOUDRU_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `CLOUDRU_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `CLOUDRU_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `CLOUDRU_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 120)


## 35.com/三五互联 (`com35`)
### Credentials
  - `COM35_USERNAME`: Username
  - `COM35_PASSWORD`: API password

### Additional configuration
  - `COM35_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `COM35_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `COM35_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `COM35_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Connbyte (`connbyte`)
### Credentials
  - `CONNBYTE_TOKEN`: Token

### Additional configuration
  - `CONNBYTE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CONNBYTE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CONNBYTE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `CONNBYTE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ConoHa v2 (`conoha`)
### Credentials
  - `CONOHA_TENANT_ID`: Tenant ID
  - `CONOHA_API_USERNAME`: The API username
  - `CONOHA_API_PASSWORD`: The API password

### Additional configuration
  - `CONOHA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CONOHA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CONOHA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `CONOHA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `CONOHA_REGION`: The region (Default: tyo1)


## ConoHa v3 (`conohav3`)
### Credentials
  - `CONOHAV3_TENANT_ID`: Tenant ID
  - `CONOHAV3_API_USER_ID`: The API user ID
  - `CONOHAV3_API_PASSWORD`: The API password

### Additional configuration
  - `CONOHAV3_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CONOHAV3_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CONOHAV3_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `CONOHAV3_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `CONOHAV3_REGION`: The region (Default: c3j1)


## Constellix (`constellix`)
### Credentials
  - `CONSTELLIX_API_KEY`: User API key
  - `CONSTELLIX_SECRET_KEY`: User secret key

### Additional configuration
  - `CONSTELLIX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `CONSTELLIX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CONSTELLIX_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `CONSTELLIX_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Core-Networks (`corenetworks`)
### Credentials
  - `CORENETWORKS_LOGIN`: The username of the API account
  - `CORENETWORKS_PASSWORD`: The password

### Additional configuration
  - `CORENETWORKS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CORENETWORKS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CORENETWORKS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `CORENETWORKS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `CORENETWORKS_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)


## CPanel/WHM (`cpanel`)
### Credentials
  - `CPANEL_USERNAME`: username
  - `CPANEL_TOKEN`: API token
  - `CPANEL_BASE_URL`: API server URL

### Additional configuration
  - `CPANEL_MODE`: use cpanel API or WHM API (Default: cpanel)
  - `CPANEL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CPANEL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `CPANEL_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `CPANEL_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Curanet (`curanet`)
### Credentials
  - `CURANET_API_KEY`: API key

### Additional configuration
  - `CURANET_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CURANET_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CURANET_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `CURANET_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Czechia (`czechia`)
### Credentials
  - `CZECHIA_TOKEN`: Authorization token

### Additional configuration
  - `CZECHIA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `CZECHIA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `CZECHIA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `CZECHIA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DanDomain (`dandomain`)
### Credentials
  - `DANDOMAIN_API_KEY`: API key

### Additional configuration
  - `DANDOMAIN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DANDOMAIN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DANDOMAIN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DANDOMAIN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DDnss (DynDNS Service) (`ddnss`)
### Credentials
  - `DDNSS_KEY`: Update key

### Additional configuration
  - `DDNSS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DDNSS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DDNSS_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `DDNSS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DDNSS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Derak Cloud (`derak`)
### Credentials
  - `DERAK_API_KEY`: The API key

### Additional configuration
  - `DERAK_WEBSITE_ID`: Force the zone/website ID
  - `DERAK_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `DERAK_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `DERAK_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DERAK_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## deSEC.io (`desec`)
### Credentials
  - `DESEC_TOKEN`: Domain token

### Additional configuration
  - `DESEC_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 4)
  - `DESEC_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `DESEC_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `DESEC_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Designate DNSaaS for Openstack (`designate`)
### Description

There are three main ways of authenticating with Designate:

1. The first one is by using the `OS_CLOUD` environment variable and a `clouds.yaml` file.
2. The second one is using your username and password, via the `OS_USERNAME`, `OS_PASSWORD` and `OS_PROJECT_NAME` environment variables.
3. The third one is by using an application credential, via the `OS_APPLICATION_CREDENTIAL_*` and `OS_USER_ID` environment variables.

For the username/password and application methods, the `OS_AUTH_URL` and `OS_REGION_NAME` environment variables are required.

For more information, you can read about the different methods of authentication with OpenStack in the Keystone's documentation and the gophercloud documentation:

- [Keystone username/password](https://docs.openstack.org/keystone/latest/user/supported_clients.html)
- [Keystone application credentials](https://docs.openstack.org/keystone/latest/user/application_credentials.html)

Public cloud providers with support for Designate:

- [Fuga Cloud](https://fuga.cloud/)
### Credentials
  - `OS_AUTH_URL`: Identity endpoint URL
  - `OS_USERNAME`: Username
  - `OS_PASSWORD`: Password
  - `OS_USER_ID`: User ID
  - `OS_APPLICATION_CREDENTIAL_ID`: Application credential ID
  - `OS_APPLICATION_CREDENTIAL_NAME`: Application credential name
  - `OS_APPLICATION_CREDENTIAL_SECRET`: Application credential secret
  - `OS_PROJECT_NAME`: Project name
  - `OS_REGION_NAME`: Region name

### Additional configuration
  - `OS_PROJECT_ID`: Project ID
  - `OS_TENANT_NAME`: Tenant name (deprecated see OS_PROJECT_NAME and OS_PROJECT_ID)
  - `DESIGNATE_ZONE_NAME`: The zone name to use in the OpenStack Project to manage TXT records.
  - `DESIGNATE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `DESIGNATE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `DESIGNATE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 10)


## Digital Ocean (`digitalocean`)
### Credentials
  - `DO_AUTH_TOKEN`: Authentication token

### Additional configuration
  - `DO_API_URL`: The URL of the API
  - `DO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `DO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 30)
  - `DO_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Dinahosting (`dinahosting`)
### Credentials
  - `DINAHOSTING_USERNAME`: Username
  - `DINAHOSTING_PASSWORD`: Password

### Additional configuration
  - `DINAHOSTING_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DINAHOSTING_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DINAHOSTING_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DINAHOSTING_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DirectAdmin (`directadmin`)
### Credentials
  - `DIRECTADMIN_API_URL`: URL of the API
  - `DIRECTADMIN_USERNAME`: API username
  - `DIRECTADMIN_PASSWORD`: API password

### Additional configuration
  - `DIRECTADMIN_ZONE_NAME`: Zone name used to add the TXT record
  - `DIRECTADMIN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `DIRECTADMIN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DIRECTADMIN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 30)
  - `DIRECTADMIN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## 51DNS (`dns51`)
### Credentials
  - `DNS51_API_KEY`: API key
  - `DNS51_API_SECRET`: API secret

### Additional configuration
  - `DNS51_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNS51_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNS51_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNS51_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DNScale (`dnscale`)
### Credentials
  - `DNSCALE_API_TOKEN`: API token

### Additional configuration
  - `DNSCALE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNSCALE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNSCALE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNSCALE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DNSExit (`dnsexit`)
### Credentials
  - `DNSEXIT_API_KEY`: API key

### Additional configuration
  - `DNSEXIT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `DNSEXIT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `DNSEXIT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNSEXIT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## dnsHome.de (`dnshomede`)
### Credentials
  - `DNSHOMEDE_CREDENTIALS`: Comma-separated list of domain:password credential pairs

### Additional configuration
  - `DNSHOMEDE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 1200)
  - `DNSHOMEDE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 2)
  - `DNSHOMEDE_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 120)
  - `DNSHOMEDE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DNSimple (`dnsimple`)
### Description

`DNSIMPLE_BASE_URL` is optional and must be set to production (https://api.dnsimple.com).
if `DNSIMPLE_BASE_URL` is not defined or empty, the production URL is used by default.

While you can manage DNS records in the [DNSimple Sandbox environment](https://developer.dnsimple.com/sandbox/),
DNS records will not resolve, and you will not be able to satisfy the ACME DNS challenge.

To authenticate you need to provide a valid API token.
HTTP Basic Authentication is intentionally not supported.

#### API tokens

You can [generate a new API token](https://support.dnsimple.com/articles/api-access-token/) from your account page.
Only Account API tokens are supported, if you try to use a User API token you will receive an error message.
### Credentials
  - `DNSIMPLE_OAUTH_TOKEN`: OAuth token

### Additional configuration
  - `DNSIMPLE_BASE_URL`: API endpoint URL
  - `DNSIMPLE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNSIMPLE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNSIMPLE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)


## dns.la (`dnsla`)
### Credentials
  - `DNSLA_API_ID`: API ID
  - `DNSLA_API_SECRET`: API secret

### Additional configuration
  - `DNSLA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNSLA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNSLA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNSLA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DNS Made Easy (`dnsmadeeasy`)
### Credentials
  - `DNSMADEEASY_API_KEY`: The API key
  - `DNSMADEEASY_API_SECRET`: The API Secret key

### Additional configuration
  - `DNSMADEEASY_SANDBOX`: Activate the sandbox (boolean)
  - `DNSMADEEASY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNSMADEEASY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNSMADEEASY_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNSMADEEASY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## DNS.services (`dnsservices`)
### Credentials
  - `DNSSERVICES_USERNAME`: Username
  - `DNSSERVICES_PASSWORD`: Password

### Additional configuration
  - `DNSSERVICES_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNSSERVICES_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNSSERVICES_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNSSERVICES_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DNS Update (RFC2136) (`dnsupdate`)
### TSIG-GSS / RFC3645 / Kerberos

To ease the usage of DNS Update in some environments, lego provides some aliases for RFC3645.

- `DNSUPDATE_RFC3645_REALM` is an alias on `DNSUPDATE_TSIG_GSS_REALM`
- `DNSUPDATE_RFC3645_USERNAME` is an alias on `DNSUPDATE_TSIG_GSS_USERNAME`
- `DNSUPDATE_RFC3645_PASSWORD` is an alias on `DNSUPDATE_TSIG_GSS_PASSWORD`
- `DNSUPDATE_RFC3645_KEYTAB_FILE` is an alias on `DNSUPDATE_TSIG_GSS_KEYTAB_FILE`

#### Examples

```bash
# Using password

DNSUPDATE_NAMESERVER=127.0.0.1 \
DNSUPDATE_TSIG_ALGORITHM=gss-tsig. \
DNSUPDATE_RFC3645_REALM=realm.example
DNSUPDATE_RFC3645_USERNAME='xxx'
DNSUPDATE_RFC3645_PASSWORD='yyy'
lego run --dns dnsupdate -d '*.example.com' -d example.com
```

```bash
# Using a keytab file.

DNSUPDATE_NAMESERVER="127.0.0.1" \
DNSUPDATE_TSIG_ALGORITHM=gss-tsig. \
DNSUPDATE_RFC3645_REALM=realm.example \
DNSUPDATE_RFC3645_USERNAME='xxx' \
DNSUPDATE_RFC3645_KEYTAB_FILE="/path/to/my.keytab" \
lego run --dns dnsupdate -d '*.example.com' -d example.com
```
### Credentials
  - `DNSUPDATE_NAMESERVER`: Network address in the form "host" or "host:port"

### Additional configuration
  - `DNSUPDATE_TSIG_ALGORITHM`: TSIG algorithm. See [miekg/dns#tsig.go](https://github.com/miekg/dns/blob/master/tsig.go) for supported values. To disable TSIG authentication, leave the `DNSUPDATE_TSIG_KEY` or `DNSUPDATE_TSIG_SECRET` variables unset.
  - `DNSUPDATE_TSIG_KEY`: Name of the secret key as defined in DNS server configuration. To disable TSIG authentication, leave the `DNSUPDATE_TSIG_KEY` variable unset.
  - `DNSUPDATE_TSIG_SECRET`: Secret key payload. To disable TSIG authentication, leave the `DNSUPDATE_TSIG_SECRET` variable unset.
  - `DNSUPDATE_TSIG_FILE`: Path to a key file generated by tsig-keygen
  - `DNSUPDATE_TSIG_GSS_REALM`: Kerberos realm. The TSIG algorithm must be `gss-tsig.`.
  - `DNSUPDATE_TSIG_GSS_USERNAME`: Kerberos username. The TSIG algorithm must be `gss-tsig.`.
  - `DNSUPDATE_TSIG_GSS_PASSWORD`: Kerberos password. The TSIG algorithm must be `gss-tsig.`.
  - `DNSUPDATE_TSIG_GSS_KEYTAB_FILE`: Path to Kerberos keytab file. The TSIG algorithm must be `gss-tsig.`.
  - `DNSUPDATE_ZONES`: List of potential zones (separated by commas)
  - `DNSUPDATE_DNS_TIMEOUT`: API request timeout in seconds (Default: 10)
  - `DNSUPDATE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DNSUPDATE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DNSUPDATE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DNSUPDATE_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)


## Domain Offensive (do.de) (`dode`)
### Credentials
  - `DODE_TOKEN`: API token

### Additional configuration
  - `DODE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DODE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DODE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `DODE_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)


## Domeneshop (`domeneshop`)
#### API credentials

Visit the following page for information on how to create API credentials with Domeneshop:

  https://api.domeneshop.no/docs/#section/Authentication
### Credentials
  - `DOMENESHOP_API_TOKEN`: API token
  - `DOMENESHOP_API_SECRET`: API secret

### Additional configuration
  - `DOMENESHOP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `DOMENESHOP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `DOMENESHOP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DreamHost (`dreamhost`)
### Credentials
  - `DREAMHOST_API_KEY`: The API key

### Additional configuration
  - `DREAMHOST_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 60)
  - `DREAMHOST_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 3600)
  - `DREAMHOST_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Duck DNS (`duckdns`)
### Credentials
  - `DUCKDNS_TOKEN`: Account token

### Additional configuration
  - `DUCKDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DUCKDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DUCKDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `DUCKDNS_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)


## Dyn (`dyn`)
### Credentials
  - `DYN_CUSTOMER_NAME`: Customer name
  - `DYN_USER_NAME`: User name
  - `DYN_PASSWORD`: Password

### Additional configuration
  - `DYN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DYN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DYN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DYN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Dynadot (`dynadot`)
### Credentials
  - `DYNADOT_API_KEY`: API key
  - `DYNADOT_API_SECRET`: API secret

### Additional configuration
  - `DYNADOT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DYNADOT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DYNADOT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `DYNADOT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## DynDnsFree.de (`dyndnsfree`)
### Credentials
  - `DYNDNSFREE_USERNAME`: Username
  - `DYNDNSFREE_PASSWORD`: Password

### Additional configuration
  - `DYNDNSFREE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `DYNDNSFREE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `DYNDNSFREE_HTTP_TIMEOUT`: Request timeout in seconds (Default: 30)


## Dynu (`dynu`)
### Credentials
  - `DYNU_API_KEY`: API key

### Additional configuration
  - `DYNU_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `DYNU_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 180)
  - `DYNU_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `DYNU_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## EasyDNS (`easydns`)
To test with the sandbox environment set ```EASYDNS_ENDPOINT=https://sandbox.rest.easydns.net```
### Credentials
  - `EASYDNS_TOKEN`: API Token
  - `EASYDNS_KEY`: API Key

### Additional configuration
  - `EASYDNS_ENDPOINT`: The endpoint URL of the API Server
  - `EASYDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `EASYDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `EASYDNS_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `EASYDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `EASYDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## EdgeCenter (`edgecenter`)
### Credentials
  - `EDGECENTER_PERMANENT_API_TOKEN`: Permanent API token (https://edgecenter.ru/blog/permanent-api-token-explained/)

### Additional configuration
  - `EDGECENTER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `EDGECENTER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 360)
  - `EDGECENTER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `EDGECENTER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Akamai EdgeDNS (`edgedns`)
Akamai edgedns supersedes FastDNS; implementing a DNS provider for solving the DNS-01 challenge using Akamai EdgeDNS
Akamai's credentials are automatically detected in the following locations and prioritized in the following order:

1. Section-specific environment variables (where `{SECTION}` is specified using `AKAMAI_EDGERC_SECTION`):
  - `AKAMAI_{SECTION}_HOST`
  - `AKAMAI_{SECTION}_ACCESS_TOKEN`
  - `AKAMAI_{SECTION}_CLIENT_TOKEN`
  - `AKAMAI_{SECTION}_CLIENT_SECRET`
2. If `AKAMAI_EDGERC_SECTION` is not defined or is set to `default`, environment variables:
  - `AKAMAI_HOST`
  - `AKAMAI_ACCESS_TOKEN`
  - `AKAMAI_CLIENT_TOKEN`
  - `AKAMAI_CLIENT_SECRET`
3. `.edgerc` file located at `AKAMAI_EDGERC`
  - defaults to `~/.edgerc`, sections can be specified using `AKAMAI_EDGERC_SECTION`
4. Default environment variables:
  - `AKAMAI_HOST`
  - `AKAMAI_ACCESS_TOKEN`
  - `AKAMAI_CLIENT_TOKEN`
  - `AKAMAI_CLIENT_SECRET`

See also:

- [Setting up Akamai credentials](https://developer.akamai.com/api/getting-started)
- [.edgerc Format](https://developer.akamai.com/legacy/introduction/Conf_Client.html#edgercformat)
- [API Client Authentication](https://developer.akamai.com/legacy/introduction/Client_Auth.html)
- [Config from Env](https://github.com/akamai/AkamaiOPEN-edgegrid-golang/blob/master/pkg/edgegrid/config.go#L118)
- [Manage many accounts](https://techdocs.akamai.com/developer/docs/manage-many-accounts-with-one-api-client)
### Credentials
  - `AKAMAI_HOST`: API host, managed by the Akamai EdgeGrid client
  - `AKAMAI_CLIENT_TOKEN`: Client token, managed by the Akamai EdgeGrid client
  - `AKAMAI_CLIENT_SECRET`: Client secret, managed by the Akamai EdgeGrid client
  - `AKAMAI_ACCESS_TOKEN`: Access token, managed by the Akamai EdgeGrid client
  - `AKAMAI_EDGERC`: Path to the .edgerc file, managed by the Akamai EdgeGrid client
  - `AKAMAI_EDGERC_SECTION`: Configuration section, managed by the Akamai EdgeGrid client

### Additional configuration
  - `AKAMAI_ACCOUNT_SWITCH_KEY`: Target account ID when the DNS zone and credentials belong to different accounts
  - `AKAMAI_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 15)
  - `AKAMAI_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 180)
  - `AKAMAI_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)


## Tencent EdgeOne (`edgeone`)
### Credentials
  - `EDGEONE_SECRET_ID`: Access key ID
  - `EDGEONE_SECRET_KEY`: Access Key secret

### Additional configuration
  - `EDGEONE_SESSION_TOKEN`: Access Key token
  - `EDGEONE_REGION`: Region
  - `EDGEONE_ZONES_MAPPING`: Mapping between DNS zones and site IDs. (ex: 'example.org:id1,example.com:id2')
  - `EDGEONE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 30)
  - `EDGEONE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 1200)
  - `EDGEONE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `EDGEONE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Efficient IP (`efficientip`)
### Credentials
  - `EFFICIENTIP_USERNAME`: Username
  - `EFFICIENTIP_PASSWORD`: Password
  - `EFFICIENTIP_HOSTNAME`: Hostname (ex: foo.example.com)
  - `EFFICIENTIP_DNS_NAME`: DNS name (ex: dns.smart)

### Additional configuration
  - `EFFICIENTIP_INSECURE_SKIP_VERIFY`: Whether or not to verify EfficientIP API certificate
  - `EFFICIENTIP_VIEW_NAME`: View name (ex: external)
  - `EFFICIENTIP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `EFFICIENTIP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `EFFICIENTIP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Epik (`epik`)
### Credentials
  - `EPIK_SIGNATURE`: Epik API signature (https://registrar.epik.com/account/api-settings/)

### Additional configuration
  - `EPIK_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `EPIK_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `EPIK_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `EPIK_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## EuroDNS (`eurodns`)
### Credentials
  - `EURODNS_APP_ID`: Application ID
  - `EURODNS_API_KEY`: API key

### Additional configuration
  - `EURODNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `EURODNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `EURODNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `EURODNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## EUserv (`euserv`)
### Credentials
  - `EUSERV_EMAIL`: The customer email address. You can also use the customer id instead.
  - `EUSERV_PASSWORD`: The customer account password.
  - `EUSERV_ORDER_ID`: The order ID of the API contract that you want to use for this login session.

### Additional configuration
  - `EUSERV_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `EUSERV_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 360)
  - `EUSERV_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `EUSERV_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Excedo (`excedo`)
### Credentials
  - `EXCEDO_API_KEY`: API key
  - `EXCEDO_API_URL`: API base URL

### Additional configuration
  - `EXCEDO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `EXCEDO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `EXCEDO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `EXCEDO_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## External program (`exec`)
Solving the DNS-01 challenge using an external program.
### Base Configuration

| Environment Variable Name | Description                           |
|---------------------------|---------------------------------------|
| `EXEC_MODE`               | `RAW`, none                           |
| `EXEC_PATH`               | The path of the the external program. |


### Additional Configuration

| Environment Variable Name  | Description                                                        |
|----------------------------|--------------------------------------------------------------------|
| `EXEC_POLLING_INTERVAL`    | Time between DNS propagation check in seconds (Default: 3).        |
| `EXEC_PROPAGATION_TIMEOUT` | Maximum waiting time for DNS propagation in seconds (Default: 60). |
| `EXEC_SEQUENCE_INTERVAL`   | Time between sequential requests in seconds (Default: 60).         |


### Description

The file name of the external program is specified in the environment variable `EXEC_PATH`.

When it is run by lego, three command-line parameters are passed to it:
The action ("present" or "cleanup"), the fully-qualified domain name and the value for the record.

For example, requesting a certificate for the domain 'my.example.org' can be achieved by calling lego as follows:

```bash
EXEC_PATH=./update-dns.sh \
lego run --dns exec --d my.example.org
```

It will then call the program './update-dns.sh' with like this:

```bash
./update-dns.sh "present" "_acme-challenge.my.example.org." "MsijOYZxqyjGnFGwhjrhfg-Xgbl5r68WPda0J9EgqqI"
```

The program then needs to make sure the record is inserted.
When it returns an error via a non-zero exit code, lego aborts.

When the record is to be removed again,
the program is called with the first command-line parameter set to `cleanup` instead of `present`.

If you want to use the raw domain, token, and keyAuth values with your program, you can set `EXEC_MODE=RAW`:

```bash
EXEC_MODE=RAW \
EXEC_PATH=./update-dns.sh \
lego run --dns exec -d my.example.org
```

It will then call the program `./update-dns.sh` like this:

```bash
./update-dns.sh "present" "--" "my.example.org." "some-token" "KxAy-J3NwUmg9ZQuM-gP_Mq1nStaYSaP9tYQs5_-YsE.ksT-qywTd8058G-SHHWA3RAN72Pr0yWtPYmmY5UBpQ8"
```

### Commands

{{% notice note %}}
The `--` is because the token MAY start with a `-`, and the called program may try and interpret a `-` as indicating a flag.
In the case of urfave, which is commonly used,
you can use the `--` delimiter to specify the start of positional arguments, and handle such a string safely.
{{% /notice %}}

#### Present

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram present <FQDN> <record>`                |
| `RAW`   | `myprogram present -- <domain> <token> <key_auth>` |

#### Cleanup

| Mode    | Command                                            |
|---------|----------------------------------------------------|
| default | `myprogram cleanup <FQDN> <record>`                |
| `RAW`   | `myprogram cleanup -- <domain> <token> <key_auth>` |


## Exoscale (`exoscale`)
### Credentials
  - `EXOSCALE_API_KEY`: API key
  - `EXOSCALE_API_SECRET`: API secret

### Additional configuration
  - `EXOSCALE_ENDPOINT`: API endpoint URL
  - `EXOSCALE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `EXOSCALE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `EXOSCALE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `EXOSCALE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)


## F5 XC (`f5xc`)
### Credentials
  - `F5XC_API_TOKEN`: API token
  - `F5XC_TENANT_NAME`: XC Tenant shortname
  - `F5XC_GROUP_NAME`: Group name

### Additional configuration
  - `F5XC_SERVER`: Server domain (Default: console.ves.volterra.io)
  - `F5XC_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `F5XC_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `F5XC_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `F5XC_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Fornex (`fornex`)
### Credentials
  - `FORNEX_API_KEY`: API key

### Additional configuration
  - `FORNEX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `FORNEX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `FORNEX_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `FORNEX_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## freemyip.com (`freemyip`)
### Credentials
  - `FREEMYIP_TOKEN`: Account token

### Additional configuration
  - `FREEMYIP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `FREEMYIP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `FREEMYIP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `FREEMYIP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `FREEMYIP_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)


## Gandi (`gandi`)
### Credentials
  - `GANDI_API_KEY`: API key

### Additional configuration
  - `GANDI_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 60)
  - `GANDI_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 2400)
  - `GANDI_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `GANDI_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)


## Gandi Live DNS (v5) (`gandiv5`)
### Credentials
  - `GANDIV5_PERSONAL_ACCESS_TOKEN`: Personal Access Token
  - `GANDIV5_API_KEY`: API key (Deprecated)

### Additional configuration
  - `GANDIV5_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `GANDIV5_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 1200)
  - `GANDIV5_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `GANDIV5_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Google Cloud (`gcloud`)
Supports service account impersonation to access Google Cloud DNS resources across different projects or with restricted permissions.

When using impersonation, the source service account must have:
1. The "Service Account Token Creator" role on the source service account
2. The "https://www.googleapis.com/auth/cloud-platform" scope
### Credentials
  - `GCE_PROJECT`: Project name (by default, the project name is auto-detected by using the metadata service)
  - `Application Default Credentials`: [Documentation](https://cloud.google.com/docs/authentication/production#providing_credentials_to_your_application)
  - `GCE_SERVICE_ACCOUNT_FILE`: Account file path
  - `GCE_SERVICE_ACCOUNT`: Account

### Additional configuration
  - `GCE_ALLOW_PRIVATE_ZONE`: Allows requested domain to be in private DNS zone, works only with a private ACME server (by default: false)
  - `GCE_ZONE_ID`: Allows to skip the automatic detection of the zone
  - `GCE_IMPERSONATE_SERVICE_ACCOUNT`: Service account email to impersonate
  - `GCE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `GCE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 180)
  - `GCE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)


## G-Core (`gcore`)
### Credentials
  - `GCORE_PERMANENT_API_TOKEN`: Permanent API token (https://gcore.com/blog/permanent-api-token-explained/)

### Additional configuration
  - `GCORE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `GCORE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 360)
  - `GCORE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `GCORE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Gehirn (`gehirn`)
### Credentials
  - `GEHIRN_TOKEN_ID`: Token ID
  - `GEHIRN_TOKEN_SECRET`: Token secret

### Additional configuration
  - `GEHIRN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `GEHIRN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `GEHIRN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `GEHIRN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Gigahost.no (`gigahostno`)
### Credentials
  - `GIGAHOSTNO_USERNAME`: Username (optional if GIGAHOSTNO_API_KEY is defined)
  - `GIGAHOSTNO_PASSWORD`: Password (optional if GIGAHOSTNO_API_KEY is defined)
  - `GIGAHOSTNO_API_KEY`: API key (optionnal of GIGAHOSTNO_USERNAME and GIGAHOSTNO_PASSWORD are defined)

### Additional configuration
  - `GIGAHOSTNO_SECRET`: TOTP secret (Only usable with username/password)
  - `GIGAHOSTNO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `GIGAHOSTNO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `GIGAHOSTNO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `GIGAHOSTNO_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Glesys (`glesys`)
### Credentials
  - `GLESYS_API_USER`: API user
  - `GLESYS_API_KEY`: API key

### Additional configuration
  - `GLESYS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `GLESYS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 1200)
  - `GLESYS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `GLESYS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Gname (`gname`)
### Credentials
  - `GNAME_APP_ID`: App ID
  - `GNAME_APP_KEY`: App key

### Additional configuration
  - `GNAME_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `GNAME_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `GNAME_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `GNAME_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Go Daddy (`godaddy`)
GoDaddy has recently (2024-04) updated the account requirements to access parts of their production Domains API:

- Availability API: Limited to accounts with 50 or more domains.
- Management and DNS APIs: Limited to accounts with 10 or more domains and/or an active Discount Domain Club plan.

https://community.letsencrypt.org/t/getting-unauthorized-url-error-while-trying-to-get-cert-for-subdomains/217329/12
### Credentials
  - `GODADDY_API_KEY`: API key
  - `GODADDY_API_SECRET`: API secret

### Additional configuration
  - `GODADDY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `GODADDY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `GODADDY_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `GODADDY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Gravity (`gravity`)
### Credentials
  - `GRAVITY_SERVER_URL`: URL of the server
  - `GRAVITY_USERNAME`: Username
  - `GRAVITY_PASSWORD`: Password

### Additional configuration
  - `GRAVITY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `GRAVITY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `GRAVITY_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 1)
  - `GRAVITY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Hetzner (`hetzner`)
### Credentials
  - `HETZNER_API_TOKEN`: API token

### Additional configuration
  - `HETZNER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HETZNER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `HETZNER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HETZNER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Hosting.de (`hostingde`)
### Credentials
  - `HOSTINGDE_API_KEY`: API key

### Additional configuration
  - `HOSTINGDE_ZONE_NAME`: Zone name in ACE format
  - `HOSTINGDE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HOSTINGDE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `HOSTINGDE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HOSTINGDE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Hostinger (`hostinger`)
### Credentials
  - `HOSTINGER_API_TOKEN`: API Token

### Additional configuration
  - `HOSTINGER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HOSTINGER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `HOSTINGER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HOSTINGER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Hosting.nl (`hostingnl`)
### Credentials
  - `HOSTINGNL_API_KEY`: The API key

### Additional configuration
  - `HOSTINGNL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HOSTINGNL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `HOSTINGNL_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HOSTINGNL_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Hosttech (`hosttech`)
### Credentials
  - `HOSTTECH_API_KEY`: API login
  - `HOSTTECH_PASSWORD`: API password

### Additional configuration
  - `HOSTTECH_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HOSTTECH_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `HOSTTECH_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `HOSTTECH_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## HostUp (`hostup`)
### Credentials
  - `HOSTUP_API_KEY`: API token (required scopes: read:dns, write:dns, read:domains)

### Additional configuration
  - `HOSTUP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HOSTUP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `HOSTUP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HOSTUP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## http.net (`httpnet`)
### Credentials
  - `HTTPNET_API_KEY`: API key

### Additional configuration
  - `HTTPNET_ZONE_NAME`: Zone name in ACE format
  - `HTTPNET_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HTTPNET_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `HTTPNET_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HTTPNET_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## HTTP request (`httpreq`)
### Description

The server must provide:

- `POST` `/present`
- `POST` `/cleanup`

The URL of the server must be defined by `HTTPREQ_ENDPOINT`.

#### Mode

There are 2 modes (`HTTPREQ_MODE`):

- default mode:
```json
{
  "fqdn": "_acme-challenge.domain.",
  "value": "LHDhK3oGRvkiefQnx7OOczTY5Tic_xZ6HcMOc_gmtoM"
}
```

- `RAW`
```json
{
  "domain": "domain",
  "token": "token",
  "keyAuth": "key"
}
```

#### Authentication

Basic authentication (optional) can be set with some environment variables:

- `HTTPREQ_USERNAME` and `HTTPREQ_PASSWORD`
- both values must be set, otherwise basic authentication is not defined.
### Credentials
  - `HTTPREQ_MODE`: `RAW`, none
  - `HTTPREQ_ENDPOINT`: The URL of the server

### Additional configuration
  - `HTTPREQ_USERNAME`: Basic authentication username
  - `HTTPREQ_PASSWORD`: Basic authentication password
  - `HTTPREQ_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HTTPREQ_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `HTTPREQ_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Huawei Cloud (`huaweicloud`)
### Credentials
  - `HUAWEICLOUD_ACCESS_KEY_ID`: Access key ID
  - `HUAWEICLOUD_SECRET_ACCESS_KEY`: Access Key secret
  - `HUAWEICLOUD_REGION`: Region

### Additional configuration
  - `HUAWEICLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HUAWEICLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `HUAWEICLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `HUAWEICLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Hurricane Electric DNS (`hurricane`)
Before using lego to request a certificate for a given domain or wildcard (such as `my.example.org` or `*.my.example.org`),
create a TXT record named `_acme-challenge.my.example.org`, and enable dynamic updates on it.
Generate a token for each URL with Hurricane Electric's UI, and copy it down.
Stick to alphanumeric tokens for greatest reliability.

To authenticate with the Hurricane Electric API,
add each record name/token pair you want to update to the `HURRICANE_TOKENS` environment variable, as shown in the examples.
Record names (without the `_acme-challenge.` component) and their tokens are separated with colons,
while the credential pairs are concatenated into a comma-separated list, like so:

```
HURRICANE_TOKENS=my.example.org:token1,demo.example.org:token2
```

If you are issuing both a wildcard certificate and a standard certificate for a given subdomain,
you should not have repeat entries for that name, as both will use the same credential.

```
HURRICANE_TOKENS=example.org:token
```
### Credentials
  - `HURRICANE_TOKENS`: TXT record names and tokens

### Additional configuration
  - `HURRICANE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `HURRICANE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation (Default: 300)
  - `HURRICANE_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `HURRICANE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## HyperOne (`hyperone`)
### Description

Default configuration does not require any additional environment variables,
just a passport file in `~/.h1/passport.json` location.

#### Generating passport file using H1 CLI

To use this application you have to generate passport file for `sa`:

```
h1 iam project sa credential generate --name my-passport --project <project ID> --sa <sa ID> --passport-output-file ~/.h1/passport.json
```

#### Required permissions

The application requires following permissions:
-  `dns/zone/list`
-  `dns/zone.recordset/list`
-  `dns/zone.recordset/create`
-  `dns/zone.recordset/delete`
-  `dns/zone.record/create`
-  `dns/zone.record/list`
-  `dns/zone.record/delete`

All required permissions are available via platform role `tool.lego`.

### Additional configuration
  - `HYPERONE_PASSPORT_LOCATION`: Allows to pass custom passport file location (default ~/.h1/passport.json)
  - `HYPERONE_API_URL`: Allows to pass custom API Endpoint to be used in the challenge (default https://api.hyperone.com/v2)
  - `HYPERONE_LOCATION_ID`: Specifies location (region) to be used in API calls. (default pl-waw-1)
  - `HYPERONE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `HYPERONE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 2)
  - `HYPERONE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 60)
  - `HYPERONE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## IBM Cloud (SoftLayer) (`ibmcloud`)
### Credentials
  - `SOFTLAYER_USERNAME`: Username (IBM Cloud is {accountID}_{emailAddress})
  - `SOFTLAYER_API_KEY`: Classic Infrastructure API key

### Additional configuration
  - `SOFTLAYER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SOFTLAYER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SOFTLAYER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SOFTLAYER_TIMEOUT`: API request timeout in seconds (Default: 30)


## IIJ DNS Platform Service (`iijdpf`)
### Credentials
  - `IIJ_DPF_API_TOKEN`: API token
  - `IIJ_DPF_DPM_SERVICE_CODE`: IIJ Managed DNS Service's service code

### Additional configuration
  - `IIJ_DPF_API_ENDPOINT`: API endpoint URL, defaults to https://api.dns-platform.jp/dpf/v1
  - `IIJ_DPF_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `IIJ_DPF_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 660)
  - `IIJ_DPF_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)


## Infoblox (`infoblox`)
When creating an API's user ensure it has the proper permissions for the view you are working with.
### Credentials
  - `INFOBLOX_USERNAME`: Account Username
  - `INFOBLOX_PASSWORD`: Account Password
  - `INFOBLOX_HOST`: Host URI

### Additional configuration
  - `INFOBLOX_DNS_VIEW`: The view for the TXT records (Default: External)
  - `INFOBLOX_WAPI_VERSION`: The version of WAPI being used  (Default: 2.11)
  - `INFOBLOX_PORT`: The port for the infoblox grid manager  (Default: 443)
  - `INFOBLOX_SSL_VERIFY`: Whether or not to verify the TLS certificate  (Default: true)
  - `INFOBLOX_CA_CERTIFICATE`: The path to the CA certificate (PEM encoded)
  - `INFOBLOX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `INFOBLOX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `INFOBLOX_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `INFOBLOX_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Infomaniak (`infomaniak`)
### Access token

Access token can be created at the url https://manager.infomaniak.com/v3/ng/accounts/token/list.
You will need `dns:read` and `dns:write` permissions.
### Credentials
  - `INFOMANIAK_ACCESS_TOKEN`: Access token

### Additional configuration
  - `INFOMANIAK_ENDPOINT`: API endpoint (default: https://api.infomaniak.com)
  - `INFOMANIAK_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `INFOMANIAK_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `INFOMANIAK_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `INFOMANIAK_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Internet.bs (`internetbs`)
### Credentials
  - `INTERNET_BS_API_KEY`: API key
  - `INTERNET_BS_PASSWORD`: API password

### Additional configuration
  - `INTERNET_BS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `INTERNET_BS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `INTERNET_BS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `INTERNET_BS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## INWX (`inwx`)
### Credentials
  - `INWX_USERNAME`: Username
  - `INWX_PASSWORD`: Password

### Additional configuration
  - `INWX_SHARED_SECRET`: shared secret related to 2FA
  - `INWX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `INWX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 360)
  - `INWX_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `INWX_SANDBOX`: Activate the sandbox (boolean)


## Ionos (`ionos`)
### Credentials
  - `IONOS_API_KEY`: API key `<prefix>.<secret>` https://developer.hosting.ionos.com/docs/getstarted

### Additional configuration
  - `IONOS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `IONOS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 900)
  - `IONOS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `IONOS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Ionos Cloud (`ionoscloud`)
### Credentials
  - `IONOSCLOUD_API_TOKEN`: API token

### Additional configuration
  - `IONOSCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `IONOSCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `IONOSCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `IONOSCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## IPv64 (`ipv64`)
### Credentials
  - `IPV64_API_KEY`: Account API Key

### Additional configuration
  - `IPV64_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `IPV64_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `IPV64_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ISPConfig 3 (`ispconfig`)
### Credentials
  - `ISPCONFIG_SERVER_URL`: Server URL
  - `ISPCONFIG_USERNAME`: Username
  - `ISPCONFIG_PASSWORD`: Password

### Additional configuration
  - `ISPCONFIG_INSECURE_SKIP_VERIFY`: Whether to verify the API certificate
  - `ISPCONFIG_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ISPCONFIG_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ISPCONFIG_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ISPCONFIG_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ISPConfig 3 - Dynamic DNS (DDNS) Module (`ispconfigddns`)
ISPConfig DNS provider supports leveraging the [ISPConfig 3 Dynamic DNS (DDNS) Module](https://github.com/mhofer117/ispconfig-ddns-module).

Requires the DDNS module described at https://www.ispconfig.org/ispconfig/download/

See https://www.howtoforge.com/community/threads/ispconfig-3-danymic-dns-ddns-module.87967/ for additional details.
### Credentials
  - `ISPCONFIG_DDNS_SERVER_URL`: API server URL (ex: https://panel.example.com:8080)
  - `ISPCONFIG_DDNS_TOKEN`: DDNS API token

### Additional configuration
  - `ISPCONFIG_DDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ISPCONFIG_DDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ISPCONFIG_DDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `ISPCONFIG_DDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## JD Cloud (`jdcloud`)
### Credentials
  - `JDCLOUD_ACCESS_KEY_ID`: Access key ID
  - `JDCLOUD_ACCESS_KEY_SECRET`: Access key secret

### Additional configuration
  - `JDCLOUD_REGION_ID`: Region ID (Default: cn-north-1)
  - `JDCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `JDCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `JDCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `JDCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Joker (`joker`)
### SVC mode

In the SVC mode, username and password are not your email and account passwords, but those displayed in Joker.com domain dashboard when enabling Dynamic DNS.

As per [Joker.com documentation](https://joker.com/faq/content/6/496/en/let_s-encrypt-support.html):

> 1. please log in at Joker.com, visit 'My Domains',
>    find the domain you want to add  Let's Encrypt certificate for, and chose "DNS" in the menu
>
> 2. on the top right, you will find the setting for 'Dynamic DNS'.
>    If not already active, please activate it.
>    It will not affect any other already existing DNS records of this domain.
>
> 3. please take a note of the credentials which are now shown as 'Dynamic DNS Authentication', consisting of a 'username' and a 'password'.
>
> 4. this is all you have to do here - and only once per domain.
### Credentials
  - `JOKER_API_MODE`: 'DMAPI' or 'SVC'. DMAPI is for resellers accounts. (Default: DMAPI)
  - `JOKER_USERNAME`: Joker.com username
  - `JOKER_PASSWORD`: Joker.com password
  - `JOKER_API_KEY`: API key (only with DMAPI mode)

### Additional configuration
  - `JOKER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `JOKER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `JOKER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `JOKER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)
  - `JOKER_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60), only with 'SVC' mode


## Katapult (`katapult`)
### Credentials
  - `KATAPULT_API_KEY`: API key

### Additional configuration
  - `KATAPULT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `KATAPULT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `KATAPULT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `KATAPULT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## KeyHelp (`keyhelp`)
### Credentials
  - `KEYHELP_BASE_URL`: Server URL
  - `KEYHELP_API_KEY`: API key

### Additional configuration
  - `KEYHELP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `KEYHELP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `KEYHELP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `KEYHELP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Leaseweb (`leaseweb`)
### Credentials
  - `LEASEWEB_API_KEY`: API key

### Additional configuration
  - `LEASEWEB_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `LEASEWEB_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `LEASEWEB_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `LEASEWEB_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Liara (`liara`)
### Credentials
  - `LIARA_API_KEY`: The API key

### Additional configuration
  - `LIARA_TEAM_ID`: The team ID to access services in a team
  - `LIARA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `LIARA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `LIARA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `LIARA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Amazon Lightsail (`lightsail`)
### Description

AWS Credentials are automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, [`AWS_SESSION_TOKEN`]
2. Shared credentials file (defaults to `~/.aws/credentials`, profiles can be specified using `AWS_PROFILE`)
3. Amazon EC2 IAM role

AWS region is not required to set as the Lightsail DNS zone is in global (us-east-1) region.

### Policy

The following AWS IAM policy document describes the minimum permissions required for lego to complete the DNS challenge.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "lightsail:DeleteDomainEntry",
        "lightsail:CreateDomainEntry"
      ],
      "Resource": "<Lightsail DNS zone ARN>"
    }
  ]
}
```

Replace the `Resource` value with your Lightsail DNS zone ARN.
You can retrieve the ARN using aws cli by running `aws lightsail get-domains --region us-east-1` (Lightsail web console does not show the ARN, unfortunately).
It should be in the format of `arn:aws:lightsail:global:<ACCOUNT ID>:Domain/<DOMAIN ID>`.
You also need to replace the region in the ARN to `us-east-1` (instead of `global`).

Alternatively, you can also set the `Resource` to `*` (wildcard), which allow to access all domain, but this is not recommended.
### Credentials
  - `AWS_ACCESS_KEY_ID`: Managed by the AWS client. Access key ID (`AWS_ACCESS_KEY_ID_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead)
  - `AWS_SECRET_ACCESS_KEY`: Managed by the AWS client. Secret access key (`AWS_SECRET_ACCESS_KEY_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead)
  - `DNS_ZONE`: Domain name of the DNS zone

### Additional configuration
  - `AWS_SHARED_CREDENTIALS_FILE`: Managed by the AWS client. Shared credentials file.
  - `LIGHTSAIL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `LIGHTSAIL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)


## Lima-City (`limacity`)
### Credentials
  - `LIMACITY_API_KEY`: The API key

### Additional configuration
  - `LIMACITY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 80)
  - `LIMACITY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 480)
  - `LIMACITY_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 90)
  - `LIMACITY_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `LIMACITY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Linode (v4) (`linode`)
### Credentials
  - `LINODE_TOKEN`: API token

### Additional configuration
  - `LINODE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 15)
  - `LINODE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `LINODE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `LINODE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Liquid Web (`liquidweb`)
### Credentials
  - `LWAPI_USERNAME`: Liquid Web API Username
  - `LWAPI_PASSWORD`: Liquid Web API Password

### Additional configuration
  - `LWAPI_ZONE`: DNS Zone
  - `LWAPI_URL`: Liquid Web API endpoint
  - `LWAPI_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `LWAPI_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `LWAPI_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `LWAPI_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)


## Loopia (`loopia`)
#### API user

You can [generate a new API user](https://customerzone.loopia.com/api/) from your account page.

It needs to have the following permissions:

* addZoneRecord
* getZoneRecords
* removeZoneRecord
* removeSubdomain
### Credentials
  - `LOOPIA_API_USER`: API username
  - `LOOPIA_API_PASSWORD`: API password

### Additional configuration
  - `LOOPIA_API_URL`: API endpoint. Ex: https://api.loopia.se/RPCSERV or https://api.loopia.rs/RPCSERV
  - `LOOPIA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2400)
  - `LOOPIA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `LOOPIA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `LOOPIA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)


## LuaDNS (`luadns`)
### Credentials
  - `LUADNS_API_USERNAME`: Username (your email)
  - `LUADNS_API_TOKEN`: API token

### Additional configuration
  - `LUADNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `LUADNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `LUADNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `LUADNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Mail-in-a-Box (`mailinabox`)
### Credentials
  - `MAILINABOX_EMAIL`: User email
  - `MAILINABOX_PASSWORD`: User password
  - `MAILINABOX_BASE_URL`: Base API URL (ex: https://box.example.com)

### Additional configuration
  - `MAILINABOX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 4)
  - `MAILINABOX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `MAILINABOX_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ManageEngine CloudDNS (`manageengine`)
### Credentials
  - `MANAGEENGINE_CLIENT_ID`: Client ID
  - `MANAGEENGINE_CLIENT_SECRET`: Client Secret

### Additional configuration
  - `MANAGEENGINE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `MANAGEENGINE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `MANAGEENGINE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)


## Manual (`manual`)
Solving the DNS-01 challenge using CLI prompt.
### Example

To start using the CLI prompt "provider", start lego with `--dns manual`:

```console
$ lego run --dns manual -d example.com
```

What follows are a few log print-outs, interspersed with some prompts, asking for you to do perform some actions:

```txt
No key found for account you@example.com. Generating a P256 key.
Saved key to ./.lego/accounts/acme-v02.api.letsencrypt.org/you@example.com/keys/you@example.com.key
Please review the TOS at https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf
Do you accept the TOS? Y/n
```

If you accept the linked Terms of Service, hit `Enter`.

```txt
[INFO] acme: Registering account for you@example.com
!!!! HEADS UP !!!!

Your account credentials have been saved in your
configuration directory at "./.lego/accounts".

You should make a secure backup of this folder now. This
configuration directory will also contain private keys
generated by lego and certificates obtained from the ACME
server. Making regular backups of this folder is ideal.
[INFO] [example.com] acme: Obtaining bundled SAN certificate
[INFO] [example.com] AuthURL: https://acme-v02.api.letsencrypt.org/acme/authz-v3/2345678901
[INFO] [example.com] acme: Could not find solver for: tls-alpn-01
[INFO] [example.com] acme: Could not find solver for: http-01
[INFO] [example.com] acme: use dns-01 solver
[INFO] [example.com] acme: Preparing to solve DNS-01
lego: Please create the following TXT record in your example.com. zone:
_acme-challenge.example.com. 120 IN TXT "hX0dPkG6Gfs9hUvBAchQclkyyoEKbShbpvJ9mY5q2JQ"
lego: Press 'Enter' when you are done
```

Do as instructed, and create the TXT records, and hit `Enter`.

```txt
[INFO] [example.com] acme: Trying to solve DNS-01
[INFO] [example.com] acme: Checking DNS record propagation using [192.168.8.1:53]
[INFO] Wait for propagation [timeout: 1m0s, interval: 2s]
[INFO] [example.com] acme: Waiting for DNS record propagation.
[INFO] [example.com] The server validated our request
[INFO] [example.com] acme: Cleaning DNS-01 challenge
lego: You can now remove this TXT record from your example.com. zone:
_acme-challenge.example.com. 120 IN TXT "hX0dPkG6Gfs9hUvBAchQclkyyoEKbShbpvJ9mY5q2JQ"
[INFO] [example.com] acme: Validations succeeded; requesting certificates
[INFO] [example.com] Server responded with a certificate.
```

As mentioned, you can now remove the TXT record again.

### Additional configuration
  - `MANUAL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `MANUAL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)


## Metaname (`metaname`)
### Credentials
  - `METANAME_ACCOUNT_REFERENCE`: The four-digit reference of a Metaname account
  - `METANAME_API_KEY`: API Key

### Additional configuration
  - `METANAME_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `METANAME_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `METANAME_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)


## Metaregistrar (`metaregistrar`)
### Credentials
  - `METAREGISTRAR_API_TOKEN`: The API token

### Additional configuration
  - `METAREGISTRAR_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `METAREGISTRAR_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `METAREGISTRAR_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `METAREGISTRAR_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## mijn.host (`mijnhost`)
### Credentials
  - `MIJNHOST_API_KEY`: The API key

### Additional configuration
  - `MIJNHOST_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `MIJNHOST_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `MIJNHOST_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `MIJNHOST_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `MIJNHOST_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Mittwald (`mittwald`)
### Credentials
  - `MITTWALD_TOKEN`: API token

### Additional configuration
  - `MITTWALD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `MITTWALD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `MITTWALD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `MITTWALD_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 120)
  - `MITTWALD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## myaddr.{tools,dev,io} (`myaddr`)
### Credentials
  - `MYADDR_PRIVATE_KEYS_MAPPING`: Mapping between subdomains and private keys. The format is: `<subdomain1>:<private_key1>,<subdomain2>:<private_key2>,<subdomain3>:<private_key3>`

### Additional configuration
  - `MYADDR_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `MYADDR_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `MYADDR_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 2)
  - `MYADDR_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `MYADDR_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## MyDNS.jp (`mydnsjp`)
### Credentials
  - `MYDNSJP_MASTER_ID`: Master ID
  - `MYDNSJP_PASSWORD`: Password

### Additional configuration
  - `MYDNSJP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `MYDNSJP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `MYDNSJP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## MythicBeasts (`mythicbeasts`)
If you are using specific API keys, then the username is the API ID for your API key, and the password is the API secret.

Your API key name is not needed to operate lego.
### Credentials
  - `MYTHICBEASTS_USERNAME`: User name
  - `MYTHICBEASTS_PASSWORD`: Password

### Additional configuration
  - `MYTHICBEASTS_API_ENDPOINT`: The endpoint for the API (must implement v2)
  - `MYTHICBEASTS_AUTH_API_ENDPOINT`: The endpoint for Mythic Beasts' Authentication
  - `MYTHICBEASTS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `MYTHICBEASTS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `MYTHICBEASTS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `MYTHICBEASTS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Namecheap (`namecheap`)
Configuration for [Namecheap](https://www.namecheap.com).

**To enable API access on the Namecheap production environment, some opaque requirements must be met.**
More information in the section [Enabling API Access](https://www.namecheap.com/support/api/intro/) of the Namecheap documentation.
(2020-08: Account balance of $50+, 20+ domains in your account, or purchases totaling $50+ within the last 2 years.)
### Credentials
  - `NAMECHEAP_API_USER`: API user
  - `NAMECHEAP_API_KEY`: API key

### Additional configuration
  - `NAMECHEAP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 15)
  - `NAMECHEAP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 3600)
  - `NAMECHEAP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NAMECHEAP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)
  - `NAMECHEAP_SANDBOX`: Activate the sandbox (boolean)


## Name.com (`namedotcom`)
### Credentials
  - `NAMECOM_USERNAME`: Username
  - `NAMECOM_API_TOKEN`: API token

### Additional configuration
  - `NAMECOM_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `NAMECOM_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 900)
  - `NAMECOM_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `NAMECOM_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Namesilo (`namesilo`)
### Credentials
  - `NAMESILO_API_KEY`: Client ID

### Additional configuration
  - `NAMESILO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NAMESILO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60), it is better to set larger than 15 minutes
  - `NAMESILO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600), should be in [3600, 2592000]


## FusionLayer NameSurfer (`namesurfer`)
### Credentials
  - `NAMESURFER_BASE_URL`: The base URL of NameSurfer API (jsonrpc10) endpoint URL (e.g., https://foo.example.com:8443/API/NSService_10)
  - `NAMESURFER_API_KEY`: API key name
  - `NAMESURFER_API_SECRET`: API secret

### Additional configuration
  - `NAMESURFER_VIEW`: DNS view name (optional, default: empty string)
  - `NAMESURFER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NAMESURFER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `NAMESURFER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `NAMESURFER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `NAMESURFER_INSECURE_SKIP_VERIFY`: Whether to verify the API certificate


## NearlyFreeSpeech.NET (`nearlyfreespeech`)
### Credentials
  - `NEARLYFREESPEECH_API_KEY`: API Key for API requests
  - `NEARLYFREESPEECH_LOGIN`: Username for API requests

### Additional configuration
  - `NEARLYFREESPEECH_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NEARLYFREESPEECH_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NEARLYFREESPEECH_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `NEARLYFREESPEECH_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `NEARLYFREESPEECH_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## NederHost (`nederhost`)
### Credentials
  - `NEDERHOST_API_KEY`: API key

### Additional configuration
  - `NEDERHOST_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NEDERHOST_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NEDERHOST_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NEDERHOST_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Neodigit (`neodigit`)
### Credentials
  - `NEODIGIT_TOKEN`: API token

### Additional configuration
  - `NEODIGIT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `NEODIGIT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `NEODIGIT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NEODIGIT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Netcup (`netcup`)
### Credentials
  - `NETCUP_CUSTOMER_NUMBER`: Customer number
  - `NETCUP_API_KEY`: API key
  - `NETCUP_API_PASSWORD`: API password

### Additional configuration
  - `NETCUP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 30)
  - `NETCUP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 900)
  - `NETCUP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Netlify (`netlify`)
### Credentials
  - `NETLIFY_TOKEN`: Token

### Additional configuration
  - `NETLIFY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NETLIFY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NETLIFY_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `NETLIFY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Netnod (`netnod`)
### Credentials
  - `NETNOD_TOKEN`: API token

### Additional configuration
  - `NETNOD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NETNOD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NETNOD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NETNOD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## NexDNS (`nexdns`)
The API token requires the `zones.read` and `records.write` scopes.
### Credentials
  - `NEXDNS_API_TOKEN`: API token

### Additional configuration
  - `NEXDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NEXDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NEXDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NEXDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Ngenix (`ngenix`)
### Credentials
  - `NGENIX_USERNAME`: Username
  - `NGENIX_TOKEN`: API token
  - `NGENIX_CUSTOMER_ID`: Customer ID

### Additional configuration
  - `NGENIX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 20)
  - `NGENIX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `NGENIX_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Nicmanager (`nicmanager`)
### Description

You can log in using your account name + username or using your email address.
Optionally, if TOTP is configured for your account, set `NICMANAGER_API_OTP`.
### Credentials
  - `NICMANAGER_API_LOGIN`: Login, used for Username-based login
  - `NICMANAGER_API_USERNAME`: Username, used for Username-based login
  - `NICMANAGER_API_EMAIL`: Email-based login
  - `NICMANAGER_API_PASSWORD`: Password, always required

### Additional configuration
  - `NICMANAGER_API_OTP`: TOTP Secret (optional)
  - `NICMANAGER_API_MODE`: mode: 'anycast' or 'zones' (for FreeDNS) (default: 'anycast')
  - `NICMANAGER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NICMANAGER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `NICMANAGER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 900)
  - `NICMANAGER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## RU CENTER (`nicru`)
### Credential information

You can find information about service ID and secret https://www.nic.ru/manager/oauth.cgi?step=oauth.app_list

| ENV Variable        | Parameter from page            | Example           |
|---------------------|--------------------------------|-------------------|
| NICRU_USER          | Username (Number of agreement) | NNNNNNN/NIC-D     |
| NICRU_PASSWORD      | Password account               |                   |
| NICRU_SERVICE_ID    | Application ID                 | hex-based, len 32 |
| NICRU_SECRET        | Identity endpoint              | string len 91     |
### Credentials
  - `NICRU_USER`: Agreement for an account in RU CENTER
  - `NICRU_PASSWORD`: Password for an account in RU CENTER
  - `NICRU_SERVICE_ID`: Service ID for application in DNS-hosting RU CENTER
  - `NICRU_SECRET`: Secret for application in DNS-hosting RU CENTER

### Additional configuration
  - `NICRU_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 60)
  - `NICRU_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `NICRU_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 30)


## NIFCloud (`nifcloud`)
### Credentials
  - `NIFCLOUD_ACCESS_KEY_ID`: Access key
  - `NIFCLOUD_SECRET_ACCESS_KEY`: Secret access key

### Additional configuration
  - `NIFCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NIFCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NIFCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NIFCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Njalla (`njalla`)
### Credentials
  - `NJALLA_TOKEN`: API token

### Additional configuration
  - `NJALLA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NJALLA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NJALLA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `NJALLA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Nodion (`nodion`)
### Credentials
  - `NODION_API_TOKEN`: The API token

### Additional configuration
  - `NODION_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NODION_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `NODION_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NODION_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## NS1 (`ns1`)
### Credentials
  - `NS1_API_KEY`: API key

### Additional configuration
  - `NS1_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `NS1_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `NS1_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `NS1_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Octenium (`octenium`)
### Credentials
  - `OCTENIUM_API_KEY`: API key

### Additional configuration
  - `OCTENIUM_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `OCTENIUM_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `OCTENIUM_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `OCTENIUM_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## omg.lol (`omglol`)
### Credentials
  - `OMGLOL_API_KEY`: API key

### Additional configuration
  - `OMGLOL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `OMGLOL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `OMGLOL_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `OMGLOL_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## 1cloud.ru (`onecloudru`)
### Credentials
  - `ONECLOUDRU_TOKEN`: API token

### Additional configuration
  - `ONECLOUDRU_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ONECLOUDRU_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ONECLOUDRU_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `ONECLOUDRU_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Online.net (`onlinenet`)
### Credentials
  - `ONLINENET_API_TOKEN`: API token

### Additional configuration
  - `ONLINENET_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 15)
  - `ONLINENET_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 240)
  - `ONLINENET_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ONLINENET_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
  - `ONLINENET_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)


## Openprovider (`openprovider`)
{{% notice style="warning" %}}

The provider is only available for resellers because the Openprovider API is only available for resellers.

{{% /notice %}}
### Credentials
  - `OPENPROVIDER_USERNAME`: The user's name
  - `OPENPROVIDER_PASSWORD`: The user's password

### Additional configuration
  - `OPENPROVIDER_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `OPENPROVIDER_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `OPENPROVIDER_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `OPENPROVIDER_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## OpusDNS (`opusdns`)
### Credentials
  - `OPUSDNS_API_KEY`: API key (format: opk_...)

### Additional configuration
  - `OPUSDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 4)
  - `OPUSDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `OPUSDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `OPUSDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Oracle Cloud (`oraclecloud`)
### Credentials
  - `OCI_COMPARTMENT_OCID`: Compartment OCID
  - `OCI_REGION`: Region (it can be empty if `OCI_AUTH_TYPE` is not empty).
  - `OCI_PRIVATE_KEY_PATH`: Private key file (ignored if `OCI_AUTH_TYPE` is not empty)
  - `OCI_PRIVATE_KEY_PASSWORD`: Private key password (ignored if `OCI_AUTH_TYPE` is not empty)
  - `OCI_TENANCY_OCID`: Tenancy OCID (ignored if `OCI_AUTH_TYPE` is not empty)
  - `OCI_USER_OCID`: User OCID (ignored if `OCI_AUTH_TYPE` is not empty)
  - `OCI_FINGERPRINT`: Public key fingerprint (ignored if `OCI_AUTH_TYPE` is not empty)

### Additional configuration
  - `OCI_AUTH_TYPE`: Authorization type. Possible values: 'instance_principal', 'user_principal', ''. (Default: '')
  - `TF_VAR_region`: Alias on `OCI_REGION`
  - `TF_VAR_fingerprint`: Alias on `OCI_FINGERPRINT`
  - `TF_VAR_user_ocid`: Alias on `OCI_USER_OCID`
  - `TF_VAR_tenancy_ocid`: Alias on `OCI_TENANCY_OCID`
  - `TF_VAR_private_key_path`: Alias on `OCI_PRIVATE_KEY_PATH`
  - `OCI_CONFIG_FILE`: Path to the configuration file. (only for `OCI_AUTH_TYPE=user_principal`)
  - `OCI_PROFILE`: Profile name. (only for `OCI_AUTH_TYPE=user_principal`)
  - `OCI_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `OCI_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `OCI_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `OCI_HTTP_TIMEOUT`: API request timeout in seconds (Default: 60)


## Open Telekom Cloud (`otc`)
### Credentials
  - `OTC_USER_NAME`: User name
  - `OTC_PASSWORD`: Password
  - `OTC_PROJECT_NAME`: Project name
  - `OTC_DOMAIN_NAME`: Domain name

### Additional configuration
  - `OTC_IDENTITY_ENDPOINT`: Identity endpoint URL (default: https://iam.eu-de.otc.t-systems.com:443/v3/auth/tokens)
  - `OTC_PRIVATE_ZONE`: Set to true to use private zones only (default: use public zones only)
  - `OTC_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `OTC_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `OTC_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `OTC_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `OTC_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## OVH (`ovh`)
### Application Key and Secret

Application key and secret can be created by following the [OVH guide](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/).

When requesting the consumer key, the following configuration can be used to define access rights:

```json
{
  "accessRules": [
    {
      "method": "POST",
      "path": "/domain/zone/*"
    },
    {
      "method": "DELETE",
      "path": "/domain/zone/*"
    }
  ]
}
```

### OAuth2 Client Credentials

Another method for authentication is by using OAuth2 client credentials.

An IAM policy and service account can be created by following the [OVH guide](https://help.ovhcloud.com/csm/en-manage-service-account?id=kb_article_view&sysparm_article=KB0059343).

Following IAM policies need to be authorized for the affected domain:

* dnsZone:apiovh:record/create
* dnsZone:apiovh:record/delete
* dnsZone:apiovh:refresh

### Important Note

Both authentication methods cannot be used at the same time.
### Credentials
  - `OVH_ENDPOINT`: Endpoint URL (ovh-eu or ovh-ca)
  - `OVH_APPLICATION_KEY`: Application key (Application Key authentication)
  - `OVH_APPLICATION_SECRET`: Application secret (Application Key authentication)
  - `OVH_CONSUMER_KEY`: Consumer key (Application Key authentication)
  - `OVH_CLIENT_ID`: Client ID (OAuth2)
  - `OVH_CLIENT_SECRET`: Client secret (OAuth2)
  - `OVH_ACCESS_TOKEN`: Access token

### Additional configuration
  - `OVH_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `OVH_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `OVH_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `OVH_HTTP_TIMEOUT`: API request timeout in seconds (Default: 180)


## PowerDNS (`pdns`)
### Information

Tested and confirmed to work with PowerDNS authoritative server 3.4.8 and 4.0.1. Refer to [PowerDNS documentation](https://doc.powerdns.com/md/httpapi/README/) instructions on how to enable the built-in API interface.

PowerDNS Notes:
- PowerDNS API does not currently support SSL, therefore you should take care to ensure that traffic between lego and the PowerDNS API is over a trusted network, VPN etc.
- In order to have the SOA serial automatically increment each time the `_acme-challenge` record is added/modified via the API, set `SOA-EDIT-API` to `INCEPTION-INCREMENT` for the zone in the `domainmetadata` table
- Some PowerDNS servers doesn't have root API endpoints enabled and API version autodetection will not work. In that case version number can be defined using `PDNS_API_VERSION`.
### Credentials
  - `PDNS_API_KEY`: API key
  - `PDNS_API_URL`: API URL

### Additional configuration
  - `PDNS_SERVER_NAME`: Name of the server in the URL, 'localhost' by default
  - `PDNS_API_VERSION`: Skip API version autodetection and use the provided version number.
  - `PDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `PDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `PDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `PDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## plesk.com (`plesk`)
### Credentials
  - `PLESK_SERVER_BASE_URL`: Base URL of the server (ex: https://plesk.myserver.com:8443)
  - `PLESK_USERNAME`: API username
  - `PLESK_PASSWORD`: API password

### Additional configuration
  - `PLESK_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `PLESK_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `PLESK_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `PLESK_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## PointDNS/PointHQ (`pointdns`)
### Credentials
  - `POINTDNS_USERNAME`: Username
  - `POINTDNS_PASSWORD`: Password

### Additional configuration
  - `POINTDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `POINTDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `POINTDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `POINTDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Porkbun (`porkbun`)
### Credentials
  - `PORKBUN_SECRET_API_KEY`: secret API key
  - `PORKBUN_API_KEY`: API key

### Additional configuration
  - `PORKBUN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `PORKBUN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `PORKBUN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `PORKBUN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Poweradmin (`poweradmin`)
### Credentials
  - `POWERADMIN_BASE_URL`: Base URL
  - `POWERADMIN_API_KEY`: API key

### Additional configuration
  - `POWERADMIN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `POWERADMIN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `POWERADMIN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `POWERADMIN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Rackspace (`rackspace`)
### Credentials
  - `RACKSPACE_USER`: API user
  - `RACKSPACE_API_KEY`: API key

### Additional configuration
  - `RACKSPACE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 3)
  - `RACKSPACE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `RACKSPACE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `RACKSPACE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Rage4 (`rage4`)
### Credentials
  - `RAGE4_USERNAME`: Username
  - `RAGE4_PASSWORD`: Password

### Additional configuration
  - `RAGE4_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `RAGE4_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `RAGE4_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `RAGE4_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Rain Yun/雨云 (`rainyun`)
### Credentials
  - `RAINYUN_API_KEY`: API key

### Additional configuration
  - `RAINYUN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `RAINYUN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `RAINYUN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `RAINYUN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## RcodeZero (`rcodezero`)
### Description

Generate your API Token via https://my.rcodezero.at with the `ACME` permissions.
These are special tokens with limited access for ACME requests only.

RcodeZero is an Anycast Network so the distribution of the DNS01-Challenge can take up to 2 minutes.
### Credentials
  - `RCODEZERO_API_TOKEN`: API token

### Additional configuration
  - `RCODEZERO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `RCODEZERO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 240)
  - `RCODEZERO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `RCODEZERO_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Regfish (`regfish`)
### Credentials
  - `REGFISH_API_KEY`: API key

### Additional configuration
  - `REGFISH_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `REGFISH_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `REGFISH_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `REGFISH_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## reg.ru (`regru`)
### Credentials
  - `REGRU_USERNAME`: API username
  - `REGRU_PASSWORD`: API password

### Additional configuration
  - `REGRU_TLS_CERT`: authentication certificate
  - `REGRU_TLS_KEY`: authentication private key
  - `REGRU_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `REGRU_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `REGRU_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `REGRU_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## RimuHosting (`rimuhosting`)
### Credentials
  - `RIMUHOSTING_API_KEY`: User API key

### Additional configuration
  - `RIMUHOSTING_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `RIMUHOSTING_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `RIMUHOSTING_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `RIMUHOSTING_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Amazon Route 53 (`route53`)
### Description

AWS Credentials are automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, [`AWS_SESSION_TOKEN`]
2. Shared credentials file (defaults to `~/.aws/credentials`, profiles can be specified using `AWS_PROFILE`)
3. Amazon EC2 IAM role

The AWS Region is automatically detected in the following locations and prioritized in the following order:

1. Environment variables: `AWS_REGION`
2. Shared configuration file if `AWS_SDK_LOAD_CONFIG` is set (defaults to `~/.aws/config`, profiles can be specified using `AWS_PROFILE`)

If `AWS_HOSTED_ZONE_ID` is not set, Lego tries to determine the correct public hosted zone via the FQDN.

See also:

- [sessions](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sessions.html)
- [Setting AWS Credentials](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials)
- [Setting AWS Region](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-the-region)

### IAM Policy Examples

#### Broad privileges for testing purposes

The following [IAM policy](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html) document grants access to the required APIs needed by lego to complete the DNS challenge.
A word of caution:
These permissions grant write access to any DNS record in any hosted zone,
so it is recommended to narrow them down as much as possible if you are using this policy in production.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "route53:GetChange",
        "route53:ChangeResourceRecordSets",
        "route53:ListResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/*",
        "arn:aws:route53:::change/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": "route53:ListHostedZonesByName",
      "Resource": "*"
    }
  ]
}
```

#### Least privilege policy for production purposes

The following AWS IAM policy document describes the least privilege permissions required for lego to complete the DNS challenge.
Write access is limited to a specified hosted zone's DNS TXT records with a key of `_acme-challenge.example.com`.
Replace `Z11111112222222333333` with your hosted zone ID and `example.com` with your domain name to use this policy.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "route53:GetChange",
      "Resource": "arn:aws:route53:::change/*"
    },
    {
      "Effect": "Allow",
      "Action": "route53:ListHostedZonesByName",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ListResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/Z11111112222222333333"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ChangeResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/Z11111112222222333333"
      ],
      "Condition": {
        "ForAllValues:StringEquals": {
          "route53:ChangeResourceRecordSetsNormalizedRecordNames": [
            "_acme-challenge.example.com"
          ],
          "route53:ChangeResourceRecordSetsRecordTypes": [
            "TXT"
          ]
        }
      }
    }
  ]
}
```
### Credentials
  - `AWS_ACCESS_KEY_ID`: Managed by the AWS client. Access key ID (`AWS_ACCESS_KEY_ID_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead)
  - `AWS_SECRET_ACCESS_KEY`: Managed by the AWS client. Secret access key (`AWS_SECRET_ACCESS_KEY_FILE` is not supported, use `AWS_SHARED_CREDENTIALS_FILE` instead)
  - `AWS_REGION`: Managed by the AWS client (`AWS_REGION_FILE` is not supported)
  - `AWS_HOSTED_ZONE_ID`: Override the hosted zone ID.
  - `AWS_PROFILE`: Managed by the AWS client (`AWS_PROFILE_FILE` is not supported)
  - `AWS_SDK_LOAD_CONFIG`: Managed by the AWS client. Retrieve the region from the CLI config file (`AWS_SDK_LOAD_CONFIG_FILE` is not supported)
  - `AWS_ASSUME_ROLE_ARN`: Managed by the AWS Role ARN (`AWS_ASSUME_ROLE_ARN_FILE` is not supported)
  - `AWS_EXTERNAL_ID`: Managed by STS AssumeRole API operation (`AWS_EXTERNAL_ID_FILE` is not supported)
  - `AWS_WAIT_FOR_RECORD_SETS_CHANGED`: Wait for changes to be INSYNC (it can be unstable)

### Additional configuration
  - `AWS_PRIVATE_ZONE`: Set to true to use private zones only (default: use public zones only)
  - `AWS_SHARED_CREDENTIALS_FILE`: Managed by the AWS client. Shared credentials file.
  - `AWS_MAX_RETRIES`: The number of maximum returns the service will use to make an individual API request
  - `AWS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 4)
  - `AWS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `AWS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 10)


## ANS SafeDNS (`safedns`)
### Credentials
  - `SAFEDNS_AUTH_TOKEN`: Authentication token

### Additional configuration
  - `SAFEDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SAFEDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SAFEDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SAFEDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Sakura Cloud (`sakuracloud`)
### Credentials
  - `SAKURACLOUD_ACCESS_TOKEN`: Access token
  - `SAKURACLOUD_ACCESS_TOKEN_SECRET`: Access token secret

### Additional configuration
  - `SAKURACLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SAKURACLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SAKURACLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SAKURACLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Scaleway (`scaleway`)
### Credentials
  - `SCW_SECRET_KEY`: Secret key
  - `SCW_PROJECT_ID`: Project to use (optional)

### Additional configuration
  - `SCW_ACCESS_KEY`: Access key
  - `SCW_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `SCW_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `SCW_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `SCW_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ScanNet (`scannet`)
### Credentials
  - `SCANNET_API_KEY`: API key

### Additional configuration
  - `SCANNET_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SCANNET_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SCANNET_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SCANNET_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Selectel (`selectel`)
### Credentials
  - `SELECTEL_API_TOKEN`: API token

### Additional configuration
  - `SELECTEL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SELECTEL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `SELECTEL_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `SELECTEL_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Selectel v2 (`selectelv2`)
### Credentials
  - `SELECTELV2_USERNAME`: Openstack username
  - `SELECTELV2_PASSWORD`: Openstack username's password
  - `SELECTELV2_ACCOUNT_ID`: Selectel account ID (INT)
  - `SELECTELV2_PROJECT_ID`: Cloud project ID (UUID)

### Additional configuration
  - `SELECTELV2_BASE_URL`: API endpoint URL
  - `SELECTELV2_AUTH_REGION`: Location for auth endpoint like ResellAPI or Keystone (default: 'ru-1')
  - `SELECTELV2_AUTH_URL`: Identity endpoint (default: 'https://cloud.api.selcloud.ru/identity/v3/')
  - `SELECTELV2_USER_DOMAIN_NAME`: To specify the domain name (account ID) where the user is located. (default: SELECTELV2_ACCOUNT_ID)
  - `SELECTELV2_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `SELECTELV2_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `SELECTELV2_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `SELECTELV2_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## SelfHost.(de|eu) (`selfhostde`)
SelfHost.de doesn't have an API to create or delete TXT records,
there is only an "unofficial" and undocumented endpoint to update an existing TXT record.

So, before using lego to request a certificate for a given domain or wildcard (such as `my.example.org` or `*.my.example.org`),
you must create:

- one TXT record named `_acme-challenge.my.example.org` if you are **not** using wildcard for this domain.
- two TXT records named `_acme-challenge.my.example.org` if you are using wildcard for this domain.

After that you must edit the TXT record(s) to get the ID(s).

You then must prepare the `SELFHOSTDE_RECORDS_MAPPING` environment variable with the following format:

```
<domain_A>:<record_id_A1>:<record_id_A2>,<domain_B>:<record_id_B1>:<record_id_B2>,<domain_C>:<record_id_C1>:<record_id_C2>
```

where each group of domain + record ID(s) is separated with a comma (`,`),
and the domain and record ID(s) are separated with a colon (`:`).

For example, if you want to create or renew a certificate for `my.example.org`, `*.my.example.org`, and `other.example.org`,
you would need:

- two separate records for `_acme-challenge.my.example.org`
- and another separate record for `_acme-challenge.other.example.org`

The resulting environment variable would then be: `SELFHOSTDE_RECORDS_MAPPING=my.example.com:123:456,other.example.com:789`
### Credentials
  - `SELFHOSTDE_USERNAME`: Username
  - `SELFHOSTDE_PASSWORD`: Password
  - `SELFHOSTDE_RECORDS_MAPPING`: Record IDs mapping with domains (ex: example.com:123:456,example.org:789,foo.example.com:147)

### Additional configuration
  - `SELFHOSTDE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 30)
  - `SELFHOSTDE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 240)
  - `SELFHOSTDE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SELFHOSTDE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Servercow (`servercow`)
### Credentials
  - `SERVERCOW_USERNAME`: API username
  - `SERVERCOW_PASSWORD`: API password

### Additional configuration
  - `SERVERCOW_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SERVERCOW_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SERVERCOW_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SERVERCOW_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Shellrent (`shellrent`)
### Credentials
  - `SHELLRENT_USERNAME`: Username
  - `SHELLRENT_TOKEN`: Token

### Additional configuration
  - `SHELLRENT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `SHELLRENT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `SHELLRENT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `SHELLRENT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Simply.com (`simply`)
### Credentials
  - `SIMPLY_ACCOUNT_NAME`: Account name
  - `SIMPLY_API_KEY`: API key

### Additional configuration
  - `SIMPLY_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `SIMPLY_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `SIMPLY_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SIMPLY_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Sonic (`sonic`)
### API keys

The API keys must be generated by calling the `dyndns/api_key` endpoint.

Example:

```bash
$ curl -X POST -H "Content-Type: application/json" --data '{"username":"notarealuser","password":"notarealpassword","hostname":"example.com"}' https://public-api.sonic.net/dyndns/api_key
{"userid":"12345","apikey":"4d6fbf2f9ab0fa11697470918d37625851fc0c51","result":200,"message":"OK"}
```

See https://public-api.sonic.net/dyndns/#requesting_an_api_key for additional details.

This `userid` and `apikey` combo allow modifications to any DNS entries connected to the managed domain (hostname).

Hostname should be the toplevel domain managed e.g. `example.com` not `www.example.com`.
### Credentials
  - `SONIC_USER_ID`: User ID
  - `SONIC_API_KEY`: API Key

### Additional configuration
  - `SONIC_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SONIC_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SONIC_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SONIC_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `SONIC_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## Spaceship (`spaceship`)
### Credentials
  - `SPACESHIP_API_KEY`: API key
  - `SPACESHIP_API_SECRET`: API secret

### Additional configuration
  - `SPACESHIP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `SPACESHIP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `SPACESHIP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SPACESHIP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Stackpath (Deprecated) (`stackpath`)
The Stackpath DNS provider shut down in 2024.
### Credentials
  - `STACKPATH_CLIENT_ID`: Client ID
  - `STACKPATH_CLIENT_SECRET`: Client secret
  - `STACKPATH_STACK_ID`: Stack ID

### Additional configuration
  - `STACKPATH_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `STACKPATH_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `STACKPATH_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)


## Syse (`syse`)
### Credentials
  - `SYSE_CREDENTIALS`: Comma-separated list of `zone:password` credential pairs

### Additional configuration
  - `SYSE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `SYSE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 1200)
  - `SYSE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `SYSE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Technitium (`technitium`)
Technitium DNS Server supports Dynamic Updates (RFC2136) for primary zones,
so you can also use the [RFC2136 provider](https://go-acme.github.io/lego/dns/rfc2136/index.html).

[RFC2136 provider](https://go-acme.github.io/lego/dns/rfc2136/index.html) is much better compared to the HTTP API option from security perspective.
Technitium recommends to use it in production over the HTTP API.
### Credentials
  - `TECHNITIUM_SERVER_BASE_URL`: Server base URL
  - `TECHNITIUM_API_TOKEN`: API token

### Additional configuration
  - `TECHNITIUM_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `TECHNITIUM_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `TECHNITIUM_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `TECHNITIUM_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Tele3 (`tele3`)
### Credentials
  - `TELE3_KEY`: Key
  - `TELE3_SECRET`: Secret

### Additional configuration
  - `TELE3_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `TELE3_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `TELE3_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `TELE3_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Tencent Cloud DNS (`tencentcloud`)
### Credentials
  - `TENCENTCLOUD_SECRET_ID`: Access key ID
  - `TENCENTCLOUD_SECRET_KEY`: Access Key secret

### Additional configuration
  - `TENCENTCLOUD_SESSION_TOKEN`: Access Key token
  - `TENCENTCLOUD_REGION`: Region
  - `TENCENTCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `TENCENTCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `TENCENTCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `TENCENTCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Timeweb Cloud (`timewebcloud`)
### Credentials
  - `TIMEWEBCLOUD_AUTH_TOKEN`: Authentication token

### Additional configuration
  - `TIMEWEBCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `TIMEWEBCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `TIMEWEBCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 10)


## TodayNIC/时代互联 (`todaynic`)
### Credentials
  - `TODAYNIC_AUTH_USER_ID`: account ID
  - `TODAYNIC_API_KEY`: API key

### Additional configuration
  - `TODAYNIC_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `TODAYNIC_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `TODAYNIC_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `TODAYNIC_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## TransIP (`transip`)
### Credentials
  - `TRANSIP_ACCOUNT_NAME`: Account name
  - `TRANSIP_PRIVATE_KEY_PATH`: Private key path

### Additional configuration
  - `TRANSIP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `TRANSIP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `TRANSIP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 10)
  - `TRANSIP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## UCloud (`ucloud`)
### Credentials
  - `UCLOUD_PUBLIC_KEY`: Public key
  - `UCLOUD_PRIVATE_KEY`: Private key

### Additional configuration
  - `UCLOUD_REGION`: Region
  - `UCLOUD_PROJECT_ID`: Project ID
  - `UCLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `UCLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `UCLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `UCLOUD_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Ultradns (`ultradns`)
### Credentials
  - `ULTRADNS_USERNAME`: API Username
  - `ULTRADNS_PASSWORD`: API Password

### Additional configuration
  - `ULTRADNS_ENDPOINT`: API endpoint URL, defaults to https://api.ultradns.com/
  - `ULTRADNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ULTRADNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 4)
  - `ULTRADNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)


## United-Domains (`uniteddomains`)
### Credentials
  - `UNITEDDOMAINS_API_KEY`: API key `<prefix>.<secret>` https://www.united-domains.de/help/faq-article/getting-started-with-the-united-domains-dns-api/

### Additional configuration
  - `UNITEDDOMAINS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `UNITEDDOMAINS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 900)
  - `UNITEDDOMAINS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `UNITEDDOMAINS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Variomedia (`variomedia`)
### Credentials
  - `VARIOMEDIA_API_TOKEN`: API token

### Additional configuration
  - `VARIOMEDIA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `VARIOMEDIA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `VARIOMEDIA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `VARIOMEDIA_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `VARIOMEDIA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Veesp (`veesp`)
### Credentials
  - `VEESP_USERNAME`: Username
  - `VEESP_PASSWORD`: Password

### Additional configuration
  - `VEESP_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `VEESP_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `VEESP_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `VEESP_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## VegaDNS (`vegadns`)
### Credentials
  - `SECRET_VEGADNS_KEY`: API key
  - `SECRET_VEGADNS_SECRET`: API secret
  - `VEGADNS_URL`: API endpoint URL

### Additional configuration
  - `VEGADNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 60)
  - `VEGADNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 720)
  - `VEGADNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 10)


## Vercel (`vercel`)
### Credentials
  - `VERCEL_API_TOKEN`: Authentication token

### Additional configuration
  - `VERCEL_TEAM_ID`: Team ID (ex: team_xxxxxxxxxxxxxxxxxxxxxxxx)
  - `VERCEL_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `VERCEL_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `VERCEL_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `VERCEL_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Versio.[nl|eu|uk] (`versio`)
To test with the sandbox environment set ```VERSIO_ENDPOINT=https://www.versio.nl/testapi/v1/```
### Credentials
  - `VERSIO_USERNAME`: Basic authentication username
  - `VERSIO_PASSWORD`: Basic authentication password

### Additional configuration
  - `VERSIO_ENDPOINT`: The endpoint URL of the API Server
  - `VERSIO_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `VERSIO_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `VERSIO_SEQUENCE_INTERVAL`: Time between sequential requests in seconds (Default: 60)
  - `VERSIO_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `VERSIO_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## VinylDNS (`vinyldns`)
The vinyldns integration makes use of dotted hostnames to ease permission management.
Users are required to have DELETE ACL level or zone admin permissions on the VinylDNS zone containing the target host.
### Credentials
  - `VINYLDNS_ACCESS_KEY`: The VinylDNS API key
  - `VINYLDNS_SECRET_KEY`: The VinylDNS API Secret key
  - `VINYLDNS_HOST`: The VinylDNS API URL

### Additional configuration
  - `VINYLDNS_QUOTE_VALUE`: Adds quotes around the TXT record value (Default: false)
  - `VINYLDNS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 4)
  - `VINYLDNS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `VINYLDNS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 30)
  - `VINYLDNS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Virtualname (`virtualname`)
### Credentials
  - `VIRTUALNAME_TOKEN`: API token

### Additional configuration
  - `VIRTUALNAME_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `VIRTUALNAME_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `VIRTUALNAME_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `VIRTUALNAME_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## VK Cloud (`vkcloud`)
### Credential information

You can find all required and additional information on ["Project/Keys" page](https://mcs.mail.ru/app/en/project/keys) of your cloud.

| ENV Variable               | Parameter from page |
|----------------------------|---------------------|
| VK_CLOUD_PROJECT_ID        | Project ID          |
| VK_CLOUD_USERNAME          | Username            |
| VK_CLOUD_DOMAIN_NAME       | User Domain Name    |
| VK_CLOUD_IDENTITY_ENDPOINT | Identity endpoint   |
### Credentials
  - `VK_CLOUD_PROJECT_ID`: String ID of project in VK Cloud
  - `VK_CLOUD_USERNAME`: Email of VK Cloud account
  - `VK_CLOUD_PASSWORD`: Password for VK Cloud account

### Additional configuration
  - `VK_CLOUD_DNS_ENDPOINT`: URL of DNS API. Defaults to https://mcs.mail.ru/public-dns but can be changed for usage with private clouds
  - `VK_CLOUD_IDENTITY_ENDPOINT`: URL of OpenStack Auth API, Defaults to https://infra.mail.ru:35357/v3/ but can be changed for usage with private clouds
  - `VK_CLOUD_DOMAIN_NAME`: Openstack users domain name. Defaults to `users` but can be changed for usage with private clouds
  - `VK_CLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `VK_CLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `VK_CLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)


## Volcano Engine/火山引擎 (`volcengine`)
### Credentials
  - `VOLC_ACCESSKEY`: Access Key ID (AK)
  - `VOLC_SECRETKEY`: Secret Access Key (SK)

### Additional configuration
  - `VOLC_REGION`: Region
  - `VOLC_HOST`: API host
  - `VOLC_SCHEME`: API scheme
  - `VOLC_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `VOLC_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 240)
  - `VOLC_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `VOLC_HTTP_TIMEOUT`: API request timeout in seconds (Default: 15)


## Vscale (`vscale`)
### Credentials
  - `VSCALE_API_TOKEN`: API token

### Additional configuration
  - `VSCALE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `VSCALE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `VSCALE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `VSCALE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Vultr (`vultr`)
### Credentials
  - `VULTR_API_KEY`: API key

### Additional configuration
  - `VULTR_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `VULTR_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `VULTR_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `VULTR_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Wannafind (`wannafind`)
### Credentials
  - `WANNAFIND_API_KEY`: API key

### Additional configuration
  - `WANNAFIND_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `WANNAFIND_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `WANNAFIND_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `WANNAFIND_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## webnames.ca (`webnamesca`)
### Credentials
  - `WEBNAMESCA_API_USER`: API username
  - `WEBNAMESCA_API_KEY`: API key

### Additional configuration
  - `WEBNAMESCA_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `WEBNAMESCA_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `WEBNAMESCA_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `WEBNAMESCA_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## webnames.ru (`webnamesru`)
### API Key

To obtain the key, you need to change the DNS server to `*.nameself.com`: Personal account / My domains and services / Select the required domain / DNS servers

The API key can be found: Personal account / My domains and services / Select the required domain / Zone management / acme.sh or certbot settings
### Credentials
  - `WEBNAMESRU_API_KEY`: Domain API key

### Additional configuration
  - `WEBNAMESRU_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `WEBNAMESRU_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `WEBNAMESRU_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Websupport (`websupport`)
### Credentials
  - `WEBSUPPORT_API_KEY`: API key
  - `WEBSUPPORT_SECRET`: API secret

### Additional configuration
  - `WEBSUPPORT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `WEBSUPPORT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `WEBSUPPORT_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 600)
  - `WEBSUPPORT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## WEDOS (`wedos`)
### Credentials
  - `WEDOS_USERNAME`: Username is the same as for the admin account
  - `WEDOS_WAPI_PASSWORD`: Password needs to be generated and IP allowed in the admin interface

### Additional configuration
  - `WEDOS_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `WEDOS_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 600)
  - `WEDOS_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 300)
  - `WEDOS_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## West.cn/西部数码 (`westcn`)
### Credentials
  - `WESTCN_USERNAME`: Username
  - `WESTCN_PASSWORD`: API password

### Additional configuration
  - `WESTCN_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 10)
  - `WESTCN_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 120)
  - `WESTCN_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)
  - `WESTCN_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Xinnet (`xinnet`)
### Credentials
  - `XINNET_SECRET`: Application secret
  - `XINNET_AGENT_ID`: Agent ID

### Additional configuration
  - `XINNET_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `XINNET_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `XINNET_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `XINNET_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Yandex PDD (`yandex`)
### Credentials
  - `YANDEX_PDD_TOKEN`: Basic authentication username

### Additional configuration
  - `YANDEX_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `YANDEX_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `YANDEX_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 21600)
  - `YANDEX_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Yandex 360 (`yandex360`)
### Credentials
  - `YANDEX360_OAUTH_TOKEN`: The OAuth Token
  - `YANDEX360_ORG_ID`: The organization ID

### Additional configuration
  - `YANDEX360_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `YANDEX360_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `YANDEX360_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 21600)
  - `YANDEX360_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Yandex Cloud (`yandexcloud`)
### IAM Token

The simplest way to retrieve IAM access token is usage of yc-cli,
follow [docs](https://cloud.yandex.ru/docs/iam/operations/iam-token/create-for-sa) to get it

```bash
yc iam key create --service-account-name my-robot --output key.json
cat key.json | base64
```
### Credentials
  - `YANDEX_CLOUD_IAM_TOKEN`: The base64 encoded json which contains information about iam token of service account with `dns.admin` permissions
  - `YANDEX_CLOUD_FOLDER_ID`: The string id of folder (aka project) in Yandex Cloud

### Additional configuration
  - `YANDEX_CLOUD_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `YANDEX_CLOUD_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `YANDEX_CLOUD_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 60)


## Zilore (`zilore`)
### Credentials
  - `ZILORE_ACCESS_KEY`: Access key

### Additional configuration
  - `ZILORE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ZILORE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ZILORE_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 120)
  - `ZILORE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## ZoneEdit (`zoneedit`)
### Credentials
  - `ZONEEDIT_USER`: User ID
  - `ZONEEDIT_AUTH_TOKEN`: Authentication token

### Additional configuration
  - `ZONEEDIT_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ZONEEDIT_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ZONEEDIT_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Zone.ee (`zoneee`)
### Credentials
  - `ZONEEE_API_USER`: API user
  - `ZONEEE_API_KEY`: API key

### Additional configuration
  - `ZONEEE_ENDPOINT`: API endpoint URL
  - `ZONEEE_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 5)
  - `ZONEEE_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 300)
  - `ZONEEE_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)


## Zonomi (`zonomi`)
### Credentials
  - `ZONOMI_API_KEY`: User API key

### Additional configuration
  - `ZONOMI_POLLING_INTERVAL`: Time between DNS propagation check in seconds (Default: 2)
  - `ZONOMI_PROPAGATION_TIMEOUT`: Maximum waiting time for DNS propagation in seconds (Default: 60)
  - `ZONOMI_TTL`: The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600)
  - `ZONOMI_HTTP_TIMEOUT`: API request timeout in seconds (Default: 30)
