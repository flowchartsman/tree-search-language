// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tsl

// TLS operators
const (
	ltOp         = "$lt"
	lteOp        = "$lte"
	gtOp         = "$gt"
	gteOp        = "$gte"
	eqOp         = "$eq"
	notEqOp      = "$ne"
	regexOp      = "$regex"
	notRegexOp   = "$nregex"
	likeOp       = "$like"
	notLikeOp    = "$nlike"
	inOp         = "$in"
	notInOp      = "$nin"
	betweenOp    = "$between"
	notBetweenOp = "$nbetween"
	notOp        = "$not"
	andOp        = "$and"
	orOp         = "$or"
	isNilOp      = "$nil"
	isNotNilOp   = "$nnil"
)

// opDic maps SQL'ish operators to TLS operators.
var opDic = map[string]string{
	"<":  ltOp,
	"<=": lteOp,
	">":  gtOp,
	">=": gteOp,
	"=":  eqOp,
	"!=": notEqOp,
	"<>": notEqOp,
	"~=": regexOp,
	"~!": notRegexOp,
}