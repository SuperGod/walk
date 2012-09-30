// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"time"
)

import (
	"github.com/lxn/walk"
)

type DateEdit struct {
	AssignTo      **walk.DateEdit
	Name          string
	StretchFactor int
	Row           int
	RowSpan       int
	Column        int
	ColumnSpan    int
	Font          Font
	MinDate       time.Time
	MaxDate       time.Time
	Date          time.Time
	OnDateChanged walk.EventHandler
}

func (de DateEdit) Create(parent walk.Container) error {
	w, err := walk.NewDateEdit(parent)
	if err != nil {
		return err
	}

	return InitWidget(de, w, func() error {
		if err := w.SetRange(de.MinDate, de.MaxDate); err != nil {
			return err
		}

		if !de.Date.IsZero() {
			if err := w.SetValue(de.Date); err != nil {
				return err
			}
		}

		if de.OnDateChanged != nil {
			w.ValueChanged().Attach(de.OnDateChanged)
		}

		if de.AssignTo != nil {
			*de.AssignTo = w
		}

		return nil
	})
}

func (de DateEdit) CommonInfo() (name string, stretchFactor, row, rowSpan, column, columnSpan int) {
	return de.Name, de.StretchFactor, de.Row, de.RowSpan, de.Column, de.ColumnSpan
}

func (de DateEdit) Font_() *Font {
	return &de.Font
}