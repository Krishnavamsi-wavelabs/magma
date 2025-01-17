/*
   Copyright 2020 The Magma Authors.
   This source code is licensed under the BSD-style license found in the
   LICENSE file in the root directory of this source tree.
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
 */

#pragma once
#include <sstream>
#include <cstdint>

namespace magma5g {
#define PDU_SESSION_TYPE_LENGTH 1

#define PDU_ADDRESS_CONTENT_MAX_LENGTH 12
// PDUAddress IE Class
class PDUAddressMsg {
 public:
  uint8_t iei;
  uint8_t length;
  uint8_t type_val : 3;
  uint8_t address_info[PDU_ADDRESS_CONTENT_MAX_LENGTH];

  PDUAddressMsg();
  ~PDUAddressMsg();
  int EncodePDUAddressMsg(
      PDUAddressMsg* pdu_address, uint8_t iei, uint8_t* buffer, uint32_t len);
  int DecodePDUAddressMsg(
      PDUAddressMsg* pdu_address, uint8_t iei, uint8_t* buffer, uint32_t len);
};
}  // namespace magma5g
