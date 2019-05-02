package gojq

var builtinFuncs = map[string]string{
	"not":    `def not: if . then false else true end;`,
	"map":    `def map(f): [.[] | f];`,
	"select": `def select(f): if f then . else empty end;`,
	"recurse": `
		def recurse: recurse(.[]?);
		def recurse(f): def r: ., (f | r); r;
		def recurse(f; cond): def r: ., (f | select(cond) | r); r;`,
	"while": `
		def while(cond; update):
			def _while: if cond then ., (update | _while) else empty end;
			_while;`,
	"until": `
		def until(cond; next):
			def _until: if cond then . else (next | _until) end;
			_until;`,
	"range": `
		def range($x): range(0; $x);
		def range($start; $end):
			$start | while(. < $end; . + 1);
		def range($start; $end; $step):
			if $step > 0 then $start|while(. < $end; . + $step)
			elif $step < 0 then $start|while(. > $end; . + $step)
			else empty end;`,
	"arrays":    `def arrays: select(type == "array");`,
	"objects":   `def objects: select(type == "object");`,
	"iterables": `def iterables: select(type |. == "array" or . == "object");`,
	"booleans":  `def booleans: select(type == "boolean");`,
	"numbers":   `def numbers: select(type == "number");`,
	"strings":   `def strings: select(type == "string");`,
	"nulls":     `def nulls: select(. == null);`,
	"values":    `def values: select(. != null);`,
	"scalars":   `def scalars: select(type |. != "array" and . != "object");`,
	"reverse":   `def reverse: [.[length - 1 - range(0;length)]];`,
}