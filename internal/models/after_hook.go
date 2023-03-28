package models

type AfterHook func(*Route) (*Route, error)
