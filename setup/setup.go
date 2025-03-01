package setup

import "fmt"

func Setup() string {
	fmt.Println("The setup is a simple alias addition in your ~/.basrh file.\nTo be able to use it. run 'source ~/.bashrc' to reload your terminal configurations (you just need to do it once).\nThis code is directed to unix terminals and only tested on Ubuntu.")
	fmt.Println("If the alias already exists, it will not be added again.")
	fmt.Println("The alias is 'godirscan' and it will run the dirscan program.")
	fmt.Println("The programm will look for a specific file your file system going from where you are to its children.")
	fmt.Println("For now, the code must be on your home directory, ~.")
	return addAlias()
}
