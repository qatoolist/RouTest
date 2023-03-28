package models

type BeforeHook func(*Response) (*Response, error)