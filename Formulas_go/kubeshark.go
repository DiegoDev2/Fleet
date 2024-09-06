package main

// Kubeshark - API Traffic Analyzer providing real-time visibility into Kubernetes network
// Homepage: https://www.kubeshark.co/

import (
	"fmt"
	
	"os/exec"
)

func installKubeshark() {
	// Método 1: Descargar y extraer .tar.gz
	kubeshark_tar_url := "https://github.com/kubeshark/kubeshark/archive/refs/tags/v52.3.79.tar.gz"
	kubeshark_cmd_tar := exec.Command("curl", "-L", kubeshark_tar_url, "-o", "package.tar.gz")
	err := kubeshark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubeshark_zip_url := "https://github.com/kubeshark/kubeshark/archive/refs/tags/v52.3.79.zip"
	kubeshark_cmd_zip := exec.Command("curl", "-L", kubeshark_zip_url, "-o", "package.zip")
	err = kubeshark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubeshark_bin_url := "https://github.com/kubeshark/kubeshark/archive/refs/tags/v52.3.79.bin"
	kubeshark_cmd_bin := exec.Command("curl", "-L", kubeshark_bin_url, "-o", "binary.bin")
	err = kubeshark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubeshark_src_url := "https://github.com/kubeshark/kubeshark/archive/refs/tags/v52.3.79.src.tar.gz"
	kubeshark_cmd_src := exec.Command("curl", "-L", kubeshark_src_url, "-o", "source.tar.gz")
	err = kubeshark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubeshark_cmd_direct := exec.Command("./binary")
	err = kubeshark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
