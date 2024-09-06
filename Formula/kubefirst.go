package main

// Kubefirst - GitOps Infrastructure & Application Delivery Platform for kubernetes
// Homepage: https://kubefirst.io/

import (
	"fmt"
	
	"os/exec"
)

func installKubefirst() {
	// Método 1: Descargar y extraer .tar.gz
	kubefirst_tar_url := "https://github.com/konstructio/kubefirst/archive/refs/tags/v2.6.0.tar.gz"
	kubefirst_cmd_tar := exec.Command("curl", "-L", kubefirst_tar_url, "-o", "package.tar.gz")
	err := kubefirst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubefirst_zip_url := "https://github.com/konstructio/kubefirst/archive/refs/tags/v2.6.0.zip"
	kubefirst_cmd_zip := exec.Command("curl", "-L", kubefirst_zip_url, "-o", "package.zip")
	err = kubefirst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubefirst_bin_url := "https://github.com/konstructio/kubefirst/archive/refs/tags/v2.6.0.bin"
	kubefirst_cmd_bin := exec.Command("curl", "-L", kubefirst_bin_url, "-o", "binary.bin")
	err = kubefirst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubefirst_src_url := "https://github.com/konstructio/kubefirst/archive/refs/tags/v2.6.0.src.tar.gz"
	kubefirst_cmd_src := exec.Command("curl", "-L", kubefirst_src_url, "-o", "source.tar.gz")
	err = kubefirst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubefirst_cmd_direct := exec.Command("./binary")
	err = kubefirst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
