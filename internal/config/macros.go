package config

type Macros struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	List        []string `yaml:"commands"`
}

//func (m *Macros) GetMacros(name string) *Macros {
//	for _, v := range c.Macros {
//		if v.Name == name {
//			return &v
//		}
//	}
//	return nil
//}
