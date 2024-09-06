package main

// Gping - Ping, but with a graph
// Homepage: https://github.com/orf/gping

import (
	"fmt"
	
	"os/exec"
)

func installGping() {
	// Método 1: Descargar y extraer .tar.gz
	gping_tar_url := "https://github.com/orf/gping/archive/refs/tags/gping-v1.17.3.tar.gz"
	gping_cmd_tar := exec.Command("curl", "-L", gping_tar_url, "-o", "package.tar.gz")
	err := gping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gping_zip_url := "https://github.com/orf/gping/archive/refs/tags/gping-v1.17.3.zip"
	gping_cmd_zip := exec.Command("curl", "-L", gping_zip_url, "-o", "package.zip")
	err = gping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gping_bin_url := "https://github.com/orf/gping/archive/refs/tags/gping-v1.17.3.bin"
	gping_cmd_bin := exec.Command("curl", "-L", gping_bin_url, "-o", "binary.bin")
	err = gping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gping_src_url := "https://github.com/orf/gping/archive/refs/tags/gping-v1.17.3.src.tar.gz"
	gping_cmd_src := exec.Command("curl", "-L", gping_src_url, "-o", "source.tar.gz")
	err = gping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gping_cmd_direct := exec.Command("./binary")
	err = gping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: iputils")
	exec.Command("latte", "install", "iputils").Run()
}
