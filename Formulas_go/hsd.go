package main

// Hsd - Handshake Daemon & Full Node
// Homepage: https://handshake.org

import (
	"fmt"
	
	"os/exec"
)

func installHsd() {
	// Método 1: Descargar y extraer .tar.gz
	hsd_tar_url := "https://github.com/handshake-org/hsd/archive/refs/tags/v6.1.1.tar.gz"
	hsd_cmd_tar := exec.Command("curl", "-L", hsd_tar_url, "-o", "package.tar.gz")
	err := hsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hsd_zip_url := "https://github.com/handshake-org/hsd/archive/refs/tags/v6.1.1.zip"
	hsd_cmd_zip := exec.Command("curl", "-L", hsd_zip_url, "-o", "package.zip")
	err = hsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hsd_bin_url := "https://github.com/handshake-org/hsd/archive/refs/tags/v6.1.1.bin"
	hsd_cmd_bin := exec.Command("curl", "-L", hsd_bin_url, "-o", "binary.bin")
	err = hsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hsd_src_url := "https://github.com/handshake-org/hsd/archive/refs/tags/v6.1.1.src.tar.gz"
	hsd_cmd_src := exec.Command("curl", "-L", hsd_src_url, "-o", "source.tar.gz")
	err = hsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hsd_cmd_direct := exec.Command("./binary")
	err = hsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: unbound")
exec.Command("latte", "install", "unbound")
}
