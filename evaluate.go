
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
)

type Object struct {
  Tag string
  Value interface{}
}

type Env map[string]Object

// Evaluate read expressions and return a single result
func Evaluate(env *Env, xs cirru.List) (ret Object) {
  // debugPrint(xs, *env)
  if len(xs) == 0 {
    emptyArray := Object{"list", xs}
    return emptyArray
  }

  head := xs[0]

  if token, ok := head.(cirru.Token); ok {
    // debugPrint(token.Text)
    switch token.Text {
    case "--": ret = cirruComment (env, xs[1:])
    case "array": ret = cirruArray (env, xs[1:])
    case "block": ret = cirruBlock(env, xs[1:])
    case "bool": ret = cirruBool(env, xs[1:])
    case "call": ret = cirruCall(env, xs[1:])
    case "child": ret = cirruChild(env, xs[1:])
    case "code": ret = cirruCode(env, xs[1:])
    case "eval": ret = cirruEval(env, xs[1:])
    case "float": ret = cirruFloat(env, xs[1:])
    case "get": ret = cirruGet(env, xs[1:])
    case "int": ret = cirruInt(env, xs[1:])
    case "map": ret = cirruMap(env, xs[1:])
    case "print": ret = cirruPrint(env, xs[1:])
    case "regexp": ret = cirruRegexp(env, xs[1:])
    case "require": ret = cirruRequire(env, xs[1:])
    case "self": ret = cirruSelf(env, xs[1:])
    case "set": ret = cirruSet(env, xs[1:])
    case "string": ret = cirruString(env, xs[1:])
    case "type": ret = cirruType(env, xs[1:])
    case "under": ret = cirruUnder(env, xs[1:])
    default:
      stop(token.Text, "not found")
      ret = userCall(env, xs)
    }
    return
  }
  if headExpression, ok := head.(cirru.List); ok {
    debugPrint(headExpression)
    return
  }
  return
}

func userCall(env *Env, xs cirru.List) (ret Object) {
  return
}