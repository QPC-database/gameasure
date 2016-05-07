/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"flag"

	"github.com/acksin/gmeasure/lib"
)

func main() {
	var (
		e   = gmeasure.Event{}
		tid string
	)

	flag.StringVar(&e.Category, "category", "", "Event category")
	flag.StringVar(&e.Action, "action", "", "Event action")
	flag.StringVar(&e.Label, "label", "", "Event label")
	flag.StringVar(&e.Value, "value", "", "Event value")
	flag.StringVar(&e.Cid, "clientid", "", "Client ID")
	flag.StringVar(&tid, "trackingid", "", "Google Analytics Tracking ID. XX-XXXXXXX-X")
	flag.Parse()

	ga := gmeasure.GA{TrackingID: tid}
	ga.Event(e)
}