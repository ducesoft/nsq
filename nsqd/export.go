/*
 * Copyright (c) 2000, 2099, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package nsqd

func NewProtocol(nsq *NSQD) *protocolV2 {
	return &protocolV2{nsqd: nsq}
}
