package util

import supa "github.com/nedpals/supabase-go"

func Supabase() *supa.Client {
	env := LoadEnv()
	return supa.CreateClient(env.Url, env.Key)
}
