/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package extension

import (
	"dubbo.apache.org/dubbo-go/v3/filter"
)

var (
	filters                  = make(map[string]func() filter.Filter)
	rejectedExecutionHandler = make(map[string]func() filter.RejectedExecutionHandler)
)

// SetFilter sets the filter extension with @name
// For example: hystrix/metrics/token/tracing/limit/...
func SetFilter(name string, v func() filter.Filter) {
	filters[name] = v
}

// GetFilter finds the filter extension with @name
func GetFilter(name string) (filter.Filter, bool) {
	if filters[name] == nil {
		return nil, false
	}
	return filters[name](), true
}

// SetRejectedExecutionHandler sets the RejectedExecutionHandler with @name
func SetRejectedExecutionHandler(name string, creator func() filter.RejectedExecutionHandler) {
	rejectedExecutionHandler[name] = creator
}

// GetRejectedExecutionHandler finds the RejectedExecutionHandler with @name
func GetRejectedExecutionHandler(name string) filter.RejectedExecutionHandler {
	creator, ok := rejectedExecutionHandler[name]
	if !ok {
		panic("RejectedExecutionHandler for " + name + " is not existing, make sure you have import the package " +
			"and you have register it by invoking extension.SetRejectedExecutionHandler.")
	}
	return creator()
}
