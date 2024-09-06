package main

// Pulp - Build tool for PureScript projects
// Homepage: https://github.com/purescript-contrib/pulp

import (
	"fmt"
	
	"os/exec"
)

func installPulp() {
	// Método 1: Descargar y extraer .tar.gz
	pulp_tar_url := "https://registry.npmjs.org/pulp/-/pulp-16.0.2.tgz"
	pulp_cmd_tar := exec.Command("curl", "-L", pulp_tar_url, "-o", "package.tar.gz")
	err := pulp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pulp_zip_url := "https://registry.npmjs.org/pulp/-/pulp-16.0.2.tgz"
	pulp_cmd_zip := exec.Command("curl", "-L", pulp_zip_url, "-o", "package.zip")
	err = pulp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pulp_bin_url := "https://registry.npmjs.org/pulp/-/pulp-16.0.2.tgz"
	pulp_cmd_bin := exec.Command("curl", "-L", pulp_bin_url, "-o", "binary.bin")
	err = pulp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pulp_src_url := "https://registry.npmjs.org/pulp/-/pulp-16.0.2.tgz"
	pulp_cmd_src := exec.Command("curl", "-L", pulp_src_url, "-o", "source.tar.gz")
	err = pulp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pulp_cmd_direct := exec.Command("./binary")
	err = pulp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bower")
exec.Command("latte", "install", "bower")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: purescript")
exec.Command("latte", "install", "purescript")
}
