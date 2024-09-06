package main

// Gron - Make JSON greppable
// Homepage: https://github.com/tomnomnom/gron

import (
	"fmt"
	
	"os/exec"
)

func installGron() {
	// Método 1: Descargar y extraer .tar.gz
	gron_tar_url := "https://github.com/tomnomnom/gron/archive/refs/tags/v0.7.1.tar.gz"
	gron_cmd_tar := exec.Command("curl", "-L", gron_tar_url, "-o", "package.tar.gz")
	err := gron_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gron_zip_url := "https://github.com/tomnomnom/gron/archive/refs/tags/v0.7.1.zip"
	gron_cmd_zip := exec.Command("curl", "-L", gron_zip_url, "-o", "package.zip")
	err = gron_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gron_bin_url := "https://github.com/tomnomnom/gron/archive/refs/tags/v0.7.1.bin"
	gron_cmd_bin := exec.Command("curl", "-L", gron_bin_url, "-o", "binary.bin")
	err = gron_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gron_src_url := "https://github.com/tomnomnom/gron/archive/refs/tags/v0.7.1.src.tar.gz"
	gron_cmd_src := exec.Command("curl", "-L", gron_src_url, "-o", "source.tar.gz")
	err = gron_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gron_cmd_direct := exec.Command("./binary")
	err = gron_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
