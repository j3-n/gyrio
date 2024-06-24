// Code generated from Pkl module `gyrio.KeybindConfig`. DO NOT EDIT.
package keybinds

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type KeybindConfig struct {
	Select string `pkl:"select"`

	Up string `pkl:"up"`

	Down string `pkl:"down"`

	Left string `pkl:"left"`

	Right string `pkl:"right"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a KeybindConfig
func LoadFromPath(ctx context.Context, path string) (ret *KeybindConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a KeybindConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*KeybindConfig, error) {
	var ret KeybindConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
