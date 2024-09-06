package main

// Linkerd - Command-line utility to interact with linkerd
// Homepage: https://linkerd.io

import (
	"fmt"
	
	"os/exec"
)

func installLinkerd() {
	// Método 1: Descargar y extraer .tar.gz
	linkerd_tar_url := "https://github.com/linkerd/linkerd2.git"
	linkerd_cmd_tar := exec.Command("curl", "-L", linkerd_tar_url, "-o", "package.tar.gz")
	err := linkerd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linkerd_zip_url := "https://github.com/linkerd/linkerd2.git"
	linkerd_cmd_zip := exec.Command("curl", "-L", linkerd_zip_url, "-o", "package.zip")
	err = linkerd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linkerd_bin_url := "https://github.com/linkerd/linkerd2.git"
	linkerd_cmd_bin := exec.Command("curl", "-L", linkerd_bin_url, "-o", "binary.bin")
	err = linkerd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linkerd_src_url := "https://github.com/linkerd/linkerd2.git"
	linkerd_cmd_src := exec.Command("curl", "-L", linkerd_src_url, "-o", "source.tar.gz")
	err = linkerd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linkerd_cmd_direct := exec.Command("./binary")
	err = linkerd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
