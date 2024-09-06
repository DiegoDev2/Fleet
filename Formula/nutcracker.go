package main

// Nutcracker - Proxy for memcached and redis
// Homepage: https://github.com/twitter/twemproxy

import (
	"fmt"
	
	"os/exec"
)

func installNutcracker() {
	// Método 1: Descargar y extraer .tar.gz
	nutcracker_tar_url := "https://github.com/twitter/twemproxy/archive/refs/tags/0.5.0.tar.gz"
	nutcracker_cmd_tar := exec.Command("curl", "-L", nutcracker_tar_url, "-o", "package.tar.gz")
	err := nutcracker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nutcracker_zip_url := "https://github.com/twitter/twemproxy/archive/refs/tags/0.5.0.zip"
	nutcracker_cmd_zip := exec.Command("curl", "-L", nutcracker_zip_url, "-o", "package.zip")
	err = nutcracker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nutcracker_bin_url := "https://github.com/twitter/twemproxy/archive/refs/tags/0.5.0.bin"
	nutcracker_cmd_bin := exec.Command("curl", "-L", nutcracker_bin_url, "-o", "binary.bin")
	err = nutcracker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nutcracker_src_url := "https://github.com/twitter/twemproxy/archive/refs/tags/0.5.0.src.tar.gz"
	nutcracker_cmd_src := exec.Command("curl", "-L", nutcracker_src_url, "-o", "source.tar.gz")
	err = nutcracker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nutcracker_cmd_direct := exec.Command("./binary")
	err = nutcracker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
}
