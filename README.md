<!---

Copyright (c) 2019 AT&T Intellectual Property.

Licensed under the Creative Commons License, Attribution 4.0 Intl.
(the"Documentation License"); you may not use this documentation
except incompliance with the Documentation License. You may obtain
a copy of the Documentation License at 

    https://creativecommons.org/licenses/by/4.0/

Unless required by applicable law or agreed to in writing, 
documentation distributed under the Documentation License is
distributed on an "AS IS"BASIS, WITHOUT WARRANTIES OR CONDITIONS
OF ANY KIND, either express or implied. See the Documentation
License for the specific language governing permissions and
limitations under the Documentation License.

-->

This is a comment, it will not be included)
[comment]: <> (in  the output file unless you use it in)
[comment]: <> (a reference style link.)


# RIC Integration
  
This repo contains RAN Intelligent Controller (RIC) deployments related files.


### Overview

The RIC deployment scripts are designed to deploy RIC components using helm charts. A deployment recipe yaml file that
contains parameter key:value pairs can be provided as a parameter for any deployment script in this repository. The
deployment recipe is acting as the helm override values.yaml file. The default parameters are set up to deploy a 
RIC instance using Linux Foundation repositories in a self-contained environment. 
