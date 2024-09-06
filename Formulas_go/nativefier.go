package main

// Nativefier - Wrap web apps natively
// Homepage: https://github.com/nativefier/nativefier

import (
	"fmt"
	
	"os/exec"
)

func installNativefier() {
	// Método 1: Descargar y extraer .tar.gz
	nativefier_tar_url := "https://registry.npmjs.org/nativefier/-/nativefier-52.0.0.tgz"
	nativefier_cmd_tar := exec.Command("curl", "-L", nativefier_tar_url, "-o", "package.tar.gz")
	err := nativefier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nativefier_zip_url := "https://registry.npmjs.org/nativefier/-/nativefier-52.0.0.tgz"
	nativefier_cmd_zip := exec.Command("curl", "-L", nativefier_zip_url, "-o", "package.zip")
	err = nativefier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nativefier_bin_url := "https://registry.npmjs.org/nativefier/-/nativefier-52.0.0.tgz"
	nativefier_cmd_bin := exec.Command("curl", "-L", nativefier_bin_url, "-o", "binary.bin")
	err = nativefier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nativefier_src_url := "https://registry.npmjs.org/nativefier/-/nativefier-52.0.0.tgz"
	nativefier_cmd_src := exec.Command("curl", "-L", nativefier_src_url, "-o", "source.tar.gz")
	err = nativefier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nativefier_cmd_direct := exec.Command("./binary")
	err = nativefier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
