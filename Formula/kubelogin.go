package main

// Kubelogin - OpenID Connect authentication plugin for kubectl
// Homepage: https://github.com/int128/kubelogin

import (
	"fmt"
	
	"os/exec"
)

func installKubelogin() {
	// Método 1: Descargar y extraer .tar.gz
	kubelogin_tar_url := "https://github.com/int128/kubelogin/archive/refs/tags/v1.29.0.tar.gz"
	kubelogin_cmd_tar := exec.Command("curl", "-L", kubelogin_tar_url, "-o", "package.tar.gz")
	err := kubelogin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubelogin_zip_url := "https://github.com/int128/kubelogin/archive/refs/tags/v1.29.0.zip"
	kubelogin_cmd_zip := exec.Command("curl", "-L", kubelogin_zip_url, "-o", "package.zip")
	err = kubelogin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubelogin_bin_url := "https://github.com/int128/kubelogin/archive/refs/tags/v1.29.0.bin"
	kubelogin_cmd_bin := exec.Command("curl", "-L", kubelogin_bin_url, "-o", "binary.bin")
	err = kubelogin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubelogin_src_url := "https://github.com/int128/kubelogin/archive/refs/tags/v1.29.0.src.tar.gz"
	kubelogin_cmd_src := exec.Command("curl", "-L", kubelogin_src_url, "-o", "source.tar.gz")
	err = kubelogin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubelogin_cmd_direct := exec.Command("./binary")
	err = kubelogin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: kubernetes-cli")
	exec.Command("latte", "install", "kubernetes-cli").Run()
}
