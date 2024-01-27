TITLE=Deserialize Bitcoin blockchain in Python
DESCRIPTION=A Bitcoin blockchain parser written in Python
DATE=06/10/2021
+++
A Bitcoin blockchain parser written in Python.

Code: <https://github.com/denniscmartin/bitcaviar-plus>

## Installation

Recommended:

```bash
pip install bitcaviar-plus
```

Manual:

```bash
python setup.py install
```

## Usage

Deserialize genesis block:

```python
from bitcaviar_plus.block import deserialize_block


def parse_genesis_block():
    with open('path/to/file/blk00000.dat', 'rb') as f:
        block = deserialize_block(f)
        print(block)
```

Deserialize entire blockchain:

```python
import os
from bitcaviar_plus.block import deserialize_block
from bitcaviar_plus.errors import InvalidMagicBytes


def parse_entire_blockchain():
    file_counter = -1
    while True:
        file_counter += 1
        file_name = 'path/to/file/blk{}.dat'.format(str(file_counter).zfill(5))
        with open(file_name, 'rb') as f:
            file_size = os.path.getsize(file_name)
            while f.tell() < file_size:
                try:
                    block = deserialize_block(f)
                except InvalidMagicBytes as e:
                    print(e)
```

```json
{
  "magic_number":"f9beb4d9",
  "size":"0000011d",
  "id":"000000000019d6...",
  "transaction_count":"01",
  "header":{
    "version":"00000001",
    "previous_block_id":"00000000000000...",
    "merkle_root":"4a5e1e4baab89f3a32...",
    "time":"495fab29",
    "bits":"1d00ffff",
    "nonce":"7c2bac1d"
  },
  "transactions":[
    {
      "version":"00000001",
      "input_count":"01",
      "output_count":"01",
      "lock_time":"00000000",
      "id":"4a5e1e4baab89f3a32518a8...",
      "inputs":[
        {
          "id":"0000000000000000000000...",
          "vout":"ffffffff",
          "script_sig_size":"4d",
          "script_sig":"04ffff001d01044554686520546...",
          "sequence":"ffffffff"
        }
      ],
      "outputs":[
        {
          "value":"000000012a05f200",
          "script_pub_key_size":"43",
          "script_pub_key":"4104678afdb0fe55482719..."
        }
      ]
    }
  ]
}
```

## Attribution
<https://github.com/ragestack/blockchain-parser/blob/master/blockchain-parser.py?ref=denniscmartin.com>

<https://github.com/bitcoinbook/bitcoinbook?ref=denniscmartin.com>

<https://learnmeabitcoin.com/?ref=denniscmartin.com>