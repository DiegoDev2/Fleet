package main

// Nylon - Proxy server
// Homepage: https://github.com/smeinecke/nylon

import (
	"fmt"
	
	"os/exec"
)

func installNylon() {
	// Método 1: Descargar y extraer .tar.gz
	nylon_tar_url := "https://monkey.org/~marius/nylon/nylon-1.21.tar.gz"
	nylon_cmd_tar := exec.Command("curl", "-L", nylon_tar_url, "-o", "package.tar.gz")
	err := nylon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nylon_zip_url := "https://monkey.org/~marius/nylon/nylon-1.21.zip"
	nylon_cmd_zip := exec.Command("curl", "-L", nylon_zip_url, "-o", "package.zip")
	err = nylon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nylon_bin_url := "https://monkey.org/~marius/nylon/nylon-1.21.bin"
	nylon_cmd_bin := exec.Command("curl", "-L", nylon_bin_url, "-o", "binary.bin")
	err = nylon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nylon_src_url := "https://monkey.org/~marius/nylon/nylon-1.21.src.tar.gz"
	nylon_cmd_src := exec.Command("curl", "-L", nylon_src_url, "-o", "source.tar.gz")
	err = nylon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nylon_cmd_direct := exec.Command("./binary")
	err = nylon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
}
