package main

// Pickle - PHP Extension installer
// Homepage: https://github.com/FriendsOfPHP/pickle

import (
	"fmt"
	
	"os/exec"
)

func installPickle() {
	// Método 1: Descargar y extraer .tar.gz
	pickle_tar_url := "https://github.com/FriendsOfPHP/pickle/releases/download/v0.7.11/pickle.phar"
	pickle_cmd_tar := exec.Command("curl", "-L", pickle_tar_url, "-o", "package.tar.gz")
	err := pickle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pickle_zip_url := "https://github.com/FriendsOfPHP/pickle/releases/download/v0.7.11/pickle.phar"
	pickle_cmd_zip := exec.Command("curl", "-L", pickle_zip_url, "-o", "package.zip")
	err = pickle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pickle_bin_url := "https://github.com/FriendsOfPHP/pickle/releases/download/v0.7.11/pickle.phar"
	pickle_cmd_bin := exec.Command("curl", "-L", pickle_bin_url, "-o", "binary.bin")
	err = pickle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pickle_src_url := "https://github.com/FriendsOfPHP/pickle/releases/download/v0.7.11/pickle.phar"
	pickle_cmd_src := exec.Command("curl", "-L", pickle_src_url, "-o", "source.tar.gz")
	err = pickle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pickle_cmd_direct := exec.Command("./binary")
	err = pickle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
