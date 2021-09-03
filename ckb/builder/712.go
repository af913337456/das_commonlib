package builder

/**
 * Copyright (C), 2019-2021
 * FileName: 712
 * Author:   LinGuanHong
 * Date:     2021/9/3 10:50
 * Description:
 */

var MMJsonA = `{
  "types": {
    "EIP712Domain": [
      {"name": "chainId", "type": "uint256"},
      {"name": "name", "type": "string"},
      {"name": "verifyingContract", "type": "address"},
      {"name": "version", "type": "string"}
    ],
    "Action": [
      {"name": "action", "type": "string"},
      {"name": "params", "type": "string"}
    ],
    "Cell": [
      {"name": "capacity", "type": "string"},
      {"name": "lock", "type": "string"},
      {"name": "type", "type": "string"},
      {"name": "data", "type": "string"},
      {"name": "extraData", "type": "string"}
    ],
    "Transaction": [
      {"name": "plainText", "type": "string"},
      {"name": "inputsCapacity", "type": "string"},
      {"name": "outputsCapacity", "type": "string"},
      {"name": "fee", "type": "string"},
      {"name": "action", "type": "Action"},
      {"name": "inputs", "type": "Cell[]"},
      {"name": "outputs", "type": "Cell[]"},
      {"name": "digest", "type": "bytes32"}
    ]
  },
  "primaryType": "Transaction",
  "domain": {
    "chainId": 1,
    "name": "da.systems",
    "verifyingContract": "0xb3dc32341ee4bae03c85cd663311de0b1b122955",
    "version": "1"
  },
  "message": {
    "plainText": {{ plainText }},
    "inputsCapacity": {{ inputsCapacity }},
    "outputsCapacity": {{ outputsCapacity }},
    "fee": {{ fee }},
    "action": {{ action }},
    "inputs": {{ inputs }},
    "outputs": {{ outputs }},
    "digest": %s
  }
}`

func CreateMMJsonB(txDigestHexStr string) string {
	return MMJsonA
}














