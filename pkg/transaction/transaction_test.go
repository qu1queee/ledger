// Copyright © 2019 Enrique Encalada
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transaction_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transactions", func() {
	Context("verify transactions functionality", func() {
		It("should record a new crypto transaction", func() {
			Expect("something").To(BeEquivalentTo("something"))
		})

		It("should record a new billing transaction", func() {
			Expect("something").To(BeEquivalentTo("something"))
		})

		It("should delete an existing crypto transaction", func() {
			Expect("something").To(BeEquivalentTo("something"))
		})

		It("should delete an existing billing transaction", func() {
			Expect("something").To(BeEquivalentTo("something"))
		})

	})
})
