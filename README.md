# azsume

Short for Azure Assume this repository is a "tongue in cheek" AWSifying of Azure over complications...

## Credentials for this repo

Follow the ~/.aws/credentials format and create a file in `$HOME/.azure/credentials`

Then follow the format below:

```
[<Name of Service Principal>]
ARM_CLIENT_ID=<insert client id>
ARM_CLIENT_SECRET=<insert client secret>
ARM_TENANT_ID=<insert tenant id>
```

```
[my-app-sp]
ARM_CLIENT_ID="12345678-1234-1234-1234-1234567890ab
ARM_CLIENT_SECRET="1aBc3D4eF6gH8iJkL0mN2oPqRsT4uVxW5yZ7aB"
ARM_TENANT_ID="3d7cf3eb-bb94-61d8-4ac4-98c113d6b83f"
```

## Commands (So Far)

azsume list             - Will list your Service Principals with values to pick from
azsume set <SP Name>    - Will Set your ENV variables with the values provided in the `~/.azure/credentials`
azsume get <SP Name>    - Will Print out your Service Principal values

`More to come...`