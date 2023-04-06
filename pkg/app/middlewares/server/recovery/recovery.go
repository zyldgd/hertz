/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package recovery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery/stack"
)

// Recovery returns a middleware that recovers from any panic.
// By default, it will print the time, content, and stack information of the error and write a 500.
// Overriding the Config configuration, you can customize the error printing logic.
func Recovery(opts ...Option) app.HandlerFunc {
	cfg := newOptions(opts...)

	return func(c context.Context, ctx *app.RequestContext) {
		defer func() {
			if err := recover(); err != nil {
				stacks := stack.Stack(3)

				cfg.recoveryHandler(c, ctx, err, stacks)
			}
		}()
		ctx.Next(c)
	}
}
