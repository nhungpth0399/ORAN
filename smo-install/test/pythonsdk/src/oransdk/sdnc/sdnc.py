#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# SPDX-License-Identifier: Apache-2.0
"""Onap Sdnc module."""

from typing import Dict
from oransdk.configuration import settings
from onapsdk.sdnc.sdnc_element import SdncElement

class OranSdnc(SdncElement):
    """SDNC library."""

    base_url = settings.SDNC_URL

    @classmethod
    def get_status(cls) -> str:
        """
        Get status of SDNC component.

        Returns:
           the status of the SDNC component

        """
        url = f"{cls.base_url}/apidoc/explorer/"
        status = cls.send_message('GET',
                                  'Get status of SDNC component',
                                  url)
        return status

    @classmethod
    def get_odu_oru_status(cls,
                           odu_node,
                           oru_node,
                           basic_auth: Dict[str, str]) -> dict:
        """
        Get status of SDNC component.

        Args:
           basic_auth: (Dict[str, str]) for example:{ 'username': 'bob', 'password': 'secret' }

        Returns:
           the status of the SDNC component

        """
        url = f"{cls.base_url}/rests/data/network-topology:network-topology/"\
              + f"topology=topology-netconf/node={odu_node}/yang-ext:mount/"\
              + f"o-ran-sc-du-hello-world:network-function/du-to-ru-connection={oru_node}"
        status = cls.send_message_json('GET',
                                       'Get status of Odu Oru connectivity',
                                       url,
                                       basic_auth=basic_auth)
        return status
