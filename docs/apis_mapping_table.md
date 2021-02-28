# FIWARE Open APIs mapping table

These tables shows the mapping from FIWARE Open APIs to NGSI Go commands.

## STH-Comet API 

| STH-Comet API                                                                                              | NGSI Go commands                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| GET /version                                                                                               | version                                                                                                         |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&hLimit={n}&hOffset={n}                  | hget attr --hLimit {n} --type {entityType} --id {enttiyId} --attrName {attrName}                                |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?entityType={entityType}&lastN={n}                         | hget attr --lastN {n} --type {entityType} --id {enttiyId} --attrName {attrName}                                 |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&aggrMethod={method}&aggrPeriod={period} | hget attr --arrgMethod {method} --aggrPeriod {period} --type {entityType} --id {enttiyId} --attrName {attrName} |
| DELETE /STH/v1/contextEntities                                                                             | hdelete                                                                                                         |
| DELETE /STH/v1/contextEntities/type/{entityType}/id/{entityId}                                             | hdelete --type {entityType} --id {enttiyId}                                                                     |
| DELETE /STH/v1/contextEntities/type/{entityType}/id/{entityId}/attributes/{attrName}                       | hdelete --type {entityType} --id {enttiyId} --attrName {attrName}                                               |

-   [STH-Comet API - GitHub](https://github.com/telefonicaid/fiware-sth-comet/blob/master/apiary.apib)

## QuantumLeap API 

| QuantumLeap API                                    | NGSI Go commands                                                       |
| -------------------------------------------------- | ---------------------------------------------------------------------- |
| GET /v2/                                           | apis                                                                   |
| GET /version                                       | version                                                                |
| POST /config                                       | (not yet implemented)                                                  |
| GET /health                                        | health                                                                 |
| POST /notify                                       | (not yet implemented)                                                  |
| POST /subscribe                                    | (not yet implemented)                                                  |
| GET /v2/entities                                   | hget entities                                                          |
| GET /v2/entities/{entityId}/attrs/{attrName}       | hget attr --id {entityId} --attrName {attrName}                        |
| GET /v2/entities/{entityId}/attrs/{attrName}/value | hget attr --id {entityId} --attrName {attrName} --value                |
| GET /v2/entities/{entityId}                        | hget attrs --id {entityId}                                             |
| GET /v2/entities/{entityId}/value                  | hget attrs --id {entityId} --value                                     |
| GET /v2/types/{entityType}/attrs/{attrName}        | hget attr --sameType --type {entityType} --attrName {attrName}         |
| GET /v2/types/{entityType}/attrs/{attrName}/value  | hget attr --sameType --type {entityType} --attrName {attrName} --value |
| GET /v2/types/{entityType}                         | hget attrs --sameType --type {entityType}                              |
| GET /v2/types/{entityType}/value                   | hget attrs --sameType --type {entityType} --value                      |
| GET /v2/attrs/{attrName}                           | hget attr --nTypes --attrName {attrName}                               |
| GET /v2/attrs/{attrName}/value                     | hget attr --nTypes --attrName {attrName} --value                       |
| GET /v2/attrs                                      | hget attrs --nTypes                                                    |
| GET /v2/attrs/value                                | hget attrs --nTypes --value                                            |
| DELETE /v2/entities/{entityId}                     | hdelete entity --id {entityId}                                         |
| DELETE /v2/types/{entityType}                      | hdelete entities --type {entityType}                                   |

-   [QuantumLeap API - GitHub](https://github.com/smartsdk/ngsi-timeseries-api/blob/master/specification/quantumleap.yml)

## Cygnus API

| Cygnus API                                         | NGSI Go commands                                                       |
| -------------------------------------------------- | ---------------------------------------------------------------------- |
| GET /v1/version                                    | version                                                                |
| GET /v1/stats                                      | admin statistics                                                       |
| PUT /v1/stats                                      | admin statistics --delete                                              |
| GET /v1/namemappings                               | namemappings list                                                      |
| POST /v1/namemappings                              | namemappings create                                                    |
| PUT /v1/namemappings                               | namemappings update                                                    |
| DELETE /v1/namemappings                            | namemappings delete                                                    |
| GET /v1/groupingrules                              | groupingrules list                                                     |
| GET /v1/groupingrules                              | groupingrules get                                                      |
| POST /v1/groupingrules                             | groupingrules create                                                   |
| PUT /v1/groupingrules                              | groupingrules update                                                   |
| DELETE /v1/groupingrules                           | groupingrules delete                                                   |
| POST /notify                                       | (not yet implemented)                                                  |
| GET /v1/subscriptions                              | (not yet implemented)                                                  |
| POST /v1/subscriptions                             | (not yet implemented)                                                  |
| DELETE /v1/subscriptions                           | (not yet implemented)                                                  |
| GET /admin/log                                     | admin log                                                              |
| PUT /admin/log                                     | admin log --level {log_level}                                          |
| GET /v1/admin/metrics                              | admin metrics                                                          |
| DELETE /v1/admin/metrics                           | admin metrics --delete                                                 |
| GET /v1/admin/log/loggers                          | admin loggers list                                                     |
| GET /v1/admin/log/loggers?name={name}              | admin loggers get --name {name}                                        |
| POST /v1/admin/log/loggers                         | admin loggers                                                          |
| PUT /v1/admin/log/loggers                          | admin loggers                                                          |
| DELETE /v1/admin/log/loggers                       | admin loggers delete                                                   |
| DELETE /v1/admin/log/loggers?name={name}           | admin loggers delete --name {name}                                     | 
| GET /v1/admin/log/appenders                        | admin appenders list                                                   |
| GET /v1/admin/log/appenders?name={name}            | admin appenders get --name {name}                                      |
| POST /v1/admin/log/appenders                       | admin appenders                                                        |
| PUT /v1/admin/log/appenders                        | admin appenders                                                        |
| DELETE /v1/admin/log/appenders                     | admin appenders delete                                                 |
| DELETE /v1/admin/log/appenders?name={name}         | admin appenders delete --name {name}                                   |

-   [Cygnus API - GitHub](https://github.com/telefonicaid/fiware-cygnus/blob/master/doc/cygnus-common/installation_and_administration_guide/management_interface_v1.md)

## IoT Agent Provision API

| IoT Agent Provision API     | NGSI Go commands                |
| --------------------------- | ------------------------------- |
| GET /services               | services list                   |
| POST /services              | services create                 |
| PUT /serivces               | services update                 |
| DELETE /services            | services delete                 |
| GET /devices                | devices list                    |
| GET /devices/{device_id}    | devices get --id {device_id}    |
| POST /devices/{device_id}   | devices create --id {device_id} |
| PUT /devices/{device_id}    | devices update --id {device_id} |
| DELETE /devices/{device_id} | devices delete --id {device_id} |

-   [IoT Agent Provision API - GitHub](https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/apiary/iotagent.apib)

## Perseo FE

| PESEO FE API                  | NGSI Go commands             |
| ----------------------------- | ---------------------------- |
| POST /notices                 | (not yet implemented)        |
| GET /rules                    | rules list                   |
| GET /rules/{id}               | rules get --id {rulesId}     |
| POST /rules                   | rules create                 |
| DELETE /rules/{id}            | rules delete --id {rulesId}  |
| GET /verion                   | version                      |
| PUT /admin/log?level={level}  | admin log --level {level}    |
| GET /admin/log                | admin log                    |
| GET /admin/metrics            | admin metrics                |
| GET /admin/metrics?reset=true | admin emtrics --reset        |
| DELETE /admin/metrics         | admin metrics --delete       |

-   [Perseo FE API - GitHub](https://github.com/telefonicaid/perseo-fe/blob/master/documentation/api.md)

## Perseo CORE

| PESEO CORE API           | NGSI Go commands |
| ------------------------ | ---------------- |
| GET /perseo-core/version | version          |

## Keyrock API 

| Kerrock API                                                                                                                        | NGSI Go commands                                                                                                                 |
| ---------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| GET /v1/auth/tokens                                                                                                                | token                                                                                                                            |
| POST /v1/auth/tokens                                                                                                               | token                                                                                                                            |
| DELETE /v1/auth/tokens                                                                                                             | (not yet implemented                                                                                                             |
| GET /v1/applications                                                                                                               | application --aid {application_id} list                                                                                          |
| POST /v1/applications                                                                                                              | applicaiton create                                                                                                               |
| GET /v1/applications/{application_id}                                                                                              | application --aid {application_id} get                                                                                           |
| DELETE /v1/applications/{application_id}                                                                                           | application --aid {application_id} delete                                                                                        |
| PATCH /v1/applications/{application_id}                                                                                            | application --aid {application_id} update                                                                                        |
| GET /v1/users                                                                                                                      | users list                                                                                                                       |
| POST /v1/users                                                                                                                     | users create                                                                                                                     |
| GET /v1/users/{user_id}                                                                                                            | users --uid {user_id} get                                                                                                        |
| DELETE /v1/users/{user_id}                                                                                                         | users --uid {user_id} delete                                                                                                     |
| PATCH /v1/users/{user_id}                                                                                                          | users --uid {user_id} update                                                                                                     |
| GET /v1/organizations                                                                                                              | organizations --oid {organization_id} list                                                                                       |
| POST /v1/organizations                                                                                                             | organizations --oid {organization_id} create                                                                                     |
| GET /v1/organizations/{organization_id}                                                                                            | organizations --oid {organization_id} get                                                                                        |
| DELETE /v1/organizations/{organization_id}                                                                                         | organizations --oid {organization_id} delete                                                                                     |
| PATCH /v1/organizations/{organization_id}                                                                                          | organizations --oid {organization_id} update                                                                                     |
| GET /v1/applications/{application_id}/roles                                                                                        | applications --aid {application_id} role --rid {role_id} list                                                                    |
| POST /v1/applications/{application_id}/roles                                                                                       | applications --aid {application_id} role --rid {role_id} create                                                                  |
| GET /v1/applications/{application_id}/roles/{role_id}                                                                              | applications --aid {application_id} role --rid {role_id} get                                                                     |
| DELETE /v1/applications/{application_id}/roles/{role_id}                                                                           | applications --aid {application_id} role --rid {role_id} delete                                                                  |
| PATCH /v1/applications/{application_id}/roles/{role_id}                                                                            | applications --aid {application_id} role --rid {role_id} update                                                                  |
| GET /v1/applications/{application_id}/permissions                                                                                  | applications --aid {application_id} permissions list                                                                             |
| POST /v1/applications/{application_id}/permissions                                                                                 | applications --aid {application_id} permissions create                                                                           |
| GET /v1/applications/{application_id}/permissions/{permission_id}                                                                  | applications --aid {application_id} permissions --pid {permission_id} get                                                        |
| DELETE /v1/applications/{application_id}/permissions/{permission_id}                                                               | applications --aid {application_id} permissions --pid {permission_id} delete                                                     |
| PATCH /v1/applications/{application_id}/permissions/{permission_id}                                                                | applications --aid {application_id} permissions --pid {permission_id} update                                                     |
| GET /v1/applications/{application_id}/pep_proxies                                                                                  | applications --aid {application_id} pep list                                                                                     |
| POST /v1/applications/{application_id}/pep_proxies                                                                                 | applications --aid {application_id} pep create                                                                                   |
| DELETE /v1/applications/{application_id}/pep_proxies                                                                               | applications --aid {application_id} pep delete                                                                                   |
| PATCH /v1/applications/{application_id}/pep_proxies                                                                                | applications --aid {application_id} pep reset                                                                                    |
| GET /v1/applications/{application_id}/iot_agents                                                                                   | applications --aid {application_id} iota list                                                                                    |
| POST /v1/applications/{application_id}/iot_agents                                                                                  | applications --aid {application_id} iota create                                                                                  |
| GET /v1/applications/{application_id}/permissions/{iot_agent_id}                                                                   | applications --aid {application_id} iota -iid {iot_agent_id} get                                                                 |
| DELETE /v1/applications/{application_id}/permissions/{iot_agent_id}                                                                | applications --aid {application_id} iota -iid {iot_agent_id} delete                                                              |
| PATCH /v1/applications/{application_id}/permissions/{iot_agent_id}                                                                 | applications --aid {application_id} iota -iid {iot_agent_id} reset                                                               |
| GET /v1/applications/{application_id}/roles/{role_id}/permissions                                                                  | applications --aid {application_id} role --rid {role_id}s permissions --pid {permission_id}                                      |
| POST /v1/applications/{application_id}/roles/{role_id}/permissions/{permission_id}                                                 | applications --aid {application_id} role --rid {role_id}s assign                                                                 |
| DELETE /v1/applications/{application_id}/roles/{role_id}/permissions/{permission_id}                                               | applications --aid {application_id} role --rid {role_id}s unassign                                                               |
| GET /v1/applications/{application_id}/users                                                                                        | applications --aid {application_id} users --uid {user_id} list                                                                   |
| GET /v1/applications/{application_id}/users/{user_id}/roles                                                                        | applications --aid {application_id} users --uid {user_id} get                                                                    |
| PUT /v1/applications/{application_id}/users/{user_id}/roles/{role_id}                                                              | applications --aid {application_id} users --uid {user_id} assign --rid {role_id}                                                 |
| DELETE /v1/applications/{application_id}/users/{user_id}/roles/{role_id}                                                           | applications --aid {application_id} users --uid {user_id} unassign --rid {role_id}                                               |
| GET /v1/applications/{application_id}/organizations                                                                                | applications --aid {application_id} organizations --oid {organization_id} list                                                   |
| GET /v1/applications/{application_id}/organizations/{organization_id}/roles                                                        | applications --aid {application_id} organizations --oid {organization_id} get                                                    |  
| PUT /v1/applications/{application_id}/organizations/{organization_id}/roles/{role_id}/organization_roles/{organization_role_id}    | applications --aid {application_id} organizations --oid {organization_id} assign --rid {role_id} --orid {organization_role_id}   |
| DELETE /v1/applications/{application_id}/organizations/{organization_id}/roles/{role_id}/organization_roles/{organization_role_id} | applications --aid {application_id} organizations --oid {organization_id} unassign --rid {role_id} --orid {organization_role_id} |
| GET /v1/organizations/{organization_id}/users                                                                                      | organizations --oid {organization_id} users --uid {user_id} list                                                                 |
| GET /v1/organizations/{organization_id}/users/{user_id}/organization_roles                                                         | organizations --oid {organization_id} users --uid {user_id} get                                                                  |
| PUT /v1/organizations/{organization_id}/users/{user_id}/organization_roles/{organization_role_id                                   | organizations --oid {organization_id} users --uid {user_id} create --orid {organization_role_id}                                 |
| DELETE /v1/organizations/{organization_id}/users/{user_id}/organization_roles/{organization_role_id}                               | organizations --oid {organization_id} users --uid {user_id} delete --orid {organization_role_id}                                 |
| GET /v1/applications/{application_id}/trusted_applications                                                                         | applications --aid {application_id} trusted list                                                                                 |
| PUT /v1/applications/{application_id}/trusted_applications/{trustedApplicationId}                                                  | applications --aid {application_id} trusted add --tid {trustedApplicationId}                                                     |
| DELETE /v1/applications/{application_id}/trusted_applications/{trustedApplicationId}                                               | applications --aid {application_id} trusted delete --tid {trustedApplicationId}                                                  |
| GET /v1/service_providers/configs                                                                                                  | providers                                                                                                                        |

-   [Keyrock API - GitHub](https://github.com/ging/fiware-idm/blob/master/apiary.apib)

## WireCloud

| WireCloud API                                                                    | NGSI Go commands                                                       |
| -------------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| /api/features                                                                    | version                                                                |
| GET /api/preferences/platform                                                    |                                                                        |
| POST /api/preferences/platform                                                   |                                                                        |
| GET /api/workspaces                                                              | workspaces list                                                        |
| POST /api/workspaces                                                             | workspaces create                                                      |
| GET /api/workspace/{workspace_id}                                                | workspaces get --wid {workspace_id}                                    |
| POST /api/workspace/{workspace_id}                                               | workspaces update --wid {workspace_id}                                 |
| DELETE /api/workspace/{workspace_id}                                             | workspaces delete --wid {workspace_id}                                 |
| POST /api/workspace/{workspace_id}/preferences                                   |                                                                        |
| PATCH /api/workspace/{workspace_id}/wiring                                       |                                                                        |
| PUT /api/workspace/{workspace_id}/wiring                                         |                                                                        |
| POST /api/worksace/{workspace_id}/tabs                                           | tabs create --wid {workspace_id}                                       |
| GET /api/workspace/{workspace_id}/tab/{tab_id}                                   | tabs get --wid {workspace_id} --tid {tab_id}                           |
| UPDATE /api/workspace/{workspace_id}/tab/{tab_id}                                | tabs update --wid {workspace_id} --tid {tab_id}                        |
| DELETE /api/workspace/{workspace_id}/tab/{tab_id}                                | tabs delete --wid {workspace_id} --tid {tab_id}                        |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/preferences                      |                                                                        |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidgets                         | widgets                                                                |
| DELETE /api/workspace/{workspace_id}/tab/{tab_id}/iwidgets                       | widgets                                                                |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}             |                                                                        |
| GET /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/preferences  |                                                                        |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/preferences |                                                                        |
| GET /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/properties   |                                                                        |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/properties  |                                                                        |
| GET /api/workspace/{workspace_id}/operators/{operator_id}/variables              |                                                                        |
| GET /api/resources                                                               |                                                                        |
| POST /api/resources                                                              |                                                                        |
| GET /api/resource/{vendor}/{name}/{version}                                      |                                                                        |
| DELETE /api/resource/{vendor}/{name}/{version}                                   |                                                                        |
| DELETE DELETE /api/resource/{vendor}/{name}                                      |

-   [WireCloud API - GitHub](https://github.com/Wirecloud/wirecloud/blob/develop/docs/restapi/applicationmashup.apib)

## Scorpio API 

| Scorpio API                     | NGSI Go commands         |
| ------------------------------- | ------------------------ |
| GET /scorpio/v1/info/           | admin scorpio list       |
| GET /scorpio/v1/info/types      | admin scorpio types      |
| GET /scorpio/v1/info/localtypes | admin scorpio localtypes |
| GET /scorpio/v1/info/stats      | admin scorpio stats      |
| GET /scorpio/v1/info/health     | admin scorpio health     |
