package main

// Kubesess - Manage multiple kubernetes cluster at the same time
// Homepage: https://rentarami.se/posts/2022-08-05-kube-context-2/

import (
	"fmt"
	
	"os/exec"
)

func installKubesess() {
	// Método 1: Descargar y extraer .tar.gz
	kubesess_tar_url := "https://github.com/Ramilito/kubesess/archive/refs/tags/1.2.11.tar.gz"
	kubesess_cmd_tar := exec.Command("curl", "-L", kubesess_tar_url, "-o", "package.tar.gz")
	err := kubesess_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubesess_zip_url := "https://github.com/Ramilito/kubesess/archive/refs/tags/1.2.11.zip"
	kubesess_cmd_zip := exec.Command("curl", "-L", kubesess_zip_url, "-o", "package.zip")
	err = kubesess_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubesess_bin_url := "https://github.com/Ramilito/kubesess/archive/refs/tags/1.2.11.bin"
	kubesess_cmd_bin := exec.Command("curl", "-L", kubesess_bin_url, "-o", "binary.bin")
	err = kubesess_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubesess_src_url := "https://github.com/Ramilito/kubesess/archive/refs/tags/1.2.11.src.tar.gz"
	kubesess_cmd_src := exec.Command("curl", "-L", kubesess_src_url, "-o", "source.tar.gz")
	err = kubesess_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubesess_cmd_direct := exec.Command("./binary")
	err = kubesess_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: kubernetes-cli")
	exec.Command("latte", "install", "kubernetes-cli").Run()
}
