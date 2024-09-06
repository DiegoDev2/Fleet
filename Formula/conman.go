package main

// Conman - Serial console management program supporting a large number of devices
// Homepage: https://github.com/dun/conman

import (
	"fmt"
	
	"os/exec"
)

func installConman() {
	// Método 1: Descargar y extraer .tar.gz
	conman_tar_url := "https://github.com/dun/conman/archive/refs/tags/conman-0.3.1.tar.gz"
	conman_cmd_tar := exec.Command("curl", "-L", conman_tar_url, "-o", "package.tar.gz")
	err := conman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conman_zip_url := "https://github.com/dun/conman/archive/refs/tags/conman-0.3.1.zip"
	conman_cmd_zip := exec.Command("curl", "-L", conman_zip_url, "-o", "package.zip")
	err = conman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conman_bin_url := "https://github.com/dun/conman/archive/refs/tags/conman-0.3.1.bin"
	conman_cmd_bin := exec.Command("curl", "-L", conman_bin_url, "-o", "binary.bin")
	err = conman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conman_src_url := "https://github.com/dun/conman/archive/refs/tags/conman-0.3.1.src.tar.gz"
	conman_cmd_src := exec.Command("curl", "-L", conman_src_url, "-o", "source.tar.gz")
	err = conman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conman_cmd_direct := exec.Command("./binary")
	err = conman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: freeipmi")
	exec.Command("latte", "install", "freeipmi").Run()
}
