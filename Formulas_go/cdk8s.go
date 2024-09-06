package main

// Cdk8s - Define k8s native apps and abstractions using object-oriented programming
// Homepage: https://cdk8s.io/

import (
	"fmt"
	
	"os/exec"
)

func installCdk8s() {
	// Método 1: Descargar y extraer .tar.gz
	cdk8s_tar_url := "https://registry.npmjs.org/cdk8s-cli/-/cdk8s-cli-2.198.208.tgz"
	cdk8s_cmd_tar := exec.Command("curl", "-L", cdk8s_tar_url, "-o", "package.tar.gz")
	err := cdk8s_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdk8s_zip_url := "https://registry.npmjs.org/cdk8s-cli/-/cdk8s-cli-2.198.208.tgz"
	cdk8s_cmd_zip := exec.Command("curl", "-L", cdk8s_zip_url, "-o", "package.zip")
	err = cdk8s_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdk8s_bin_url := "https://registry.npmjs.org/cdk8s-cli/-/cdk8s-cli-2.198.208.tgz"
	cdk8s_cmd_bin := exec.Command("curl", "-L", cdk8s_bin_url, "-o", "binary.bin")
	err = cdk8s_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdk8s_src_url := "https://registry.npmjs.org/cdk8s-cli/-/cdk8s-cli-2.198.208.tgz"
	cdk8s_cmd_src := exec.Command("curl", "-L", cdk8s_src_url, "-o", "source.tar.gz")
	err = cdk8s_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdk8s_cmd_direct := exec.Command("./binary")
	err = cdk8s_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
