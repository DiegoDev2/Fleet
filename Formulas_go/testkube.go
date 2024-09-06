package main

// Testkube - Kubernetes-native framework for test definition and execution
// Homepage: https://testkube.io

import (
	"fmt"
	
	"os/exec"
)

func installTestkube() {
	// Método 1: Descargar y extraer .tar.gz
	testkube_tar_url := "https://github.com/kubeshop/testkube/archive/refs/tags/v2.1.12.tar.gz"
	testkube_cmd_tar := exec.Command("curl", "-L", testkube_tar_url, "-o", "package.tar.gz")
	err := testkube_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	testkube_zip_url := "https://github.com/kubeshop/testkube/archive/refs/tags/v2.1.12.zip"
	testkube_cmd_zip := exec.Command("curl", "-L", testkube_zip_url, "-o", "package.zip")
	err = testkube_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	testkube_bin_url := "https://github.com/kubeshop/testkube/archive/refs/tags/v2.1.12.bin"
	testkube_cmd_bin := exec.Command("curl", "-L", testkube_bin_url, "-o", "binary.bin")
	err = testkube_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	testkube_src_url := "https://github.com/kubeshop/testkube/archive/refs/tags/v2.1.12.src.tar.gz"
	testkube_cmd_src := exec.Command("curl", "-L", testkube_src_url, "-o", "source.tar.gz")
	err = testkube_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	testkube_cmd_direct := exec.Command("./binary")
	err = testkube_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: helm")
exec.Command("latte", "install", "helm")
	fmt.Println("Instalando dependencia: kubernetes-cli")
exec.Command("latte", "install", "kubernetes-cli")
}
