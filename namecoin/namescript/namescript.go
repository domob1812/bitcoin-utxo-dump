package namescript

import "github.com/btcsuite/btcd/txscript" // decoding script to check for name scripts

const OP_NAME_REGISTER = txscript.OP_1
const OP_NAME_UPDATE = txscript.OP_2

func isData (op byte) bool {
  return op >= txscript.OP_DATA_1 && op <= txscript.OP_PUSHDATA4
}

func isNameRegister (opcodes []byte) bool {
  if len (opcodes) < 5 {
    return false
  }

  return (opcodes[0] == OP_NAME_REGISTER &&
      isData (opcodes[1]) && isData (opcodes[2]) &&
      opcodes[3] == txscript.OP_2DROP && opcodes[4] == txscript.OP_DROP)
}

func isNameUpdate (opcodes []byte) bool {
  if len (opcodes) < 5 {
    return false
  }

  return (opcodes[0] == OP_NAME_UPDATE &&
      isData (opcodes[1]) && isData (opcodes[2]) &&
      opcodes[3] == txscript.OP_2DROP && opcodes[4] == txscript.OP_DROP)
}

func IsNameOp (script []byte) bool {
  tokeniser := txscript.MakeScriptTokenizer (0, script)
  var opcodes []byte
  for tokeniser.Next () {
    opcodes = append (opcodes, tokeniser.Opcode ())
  }

  if tokeniser.Err () != nil {
    return false
  }

  return isNameRegister (opcodes) || isNameUpdate (opcodes)
}
