package main

// PatchPackage - Fix broken node modules instantly
// Homepage: https://github.com/ds300/patch-package

import (
	"fmt"
	
	"os/exec"
)

func installPatchPackage() {
	// Método 1: Descargar y extraer .tar.gz
	patchpackage_tar_url := "https://registry.npmjs.org/patch-package/-/patch-package-8.0.0.tgz"
	patchpackage_cmd_tar := exec.Command("curl", "-L", patchpackage_tar_url, "-o", "package.tar.gz")
	err := patchpackage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	patchpackage_zip_url := "https://registry.npmjs.org/patch-package/-/patch-package-8.0.0.tgz"
	patchpackage_cmd_zip := exec.Command("curl", "-L", patchpackage_zip_url, "-o", "package.zip")
	err = patchpackage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	patchpackage_bin_url := "https://registry.npmjs.org/patch-package/-/patch-package-8.0.0.tgz"
	patchpackage_cmd_bin := exec.Command("curl", "-L", patchpackage_bin_url, "-o", "binary.bin")
	err = patchpackage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	patchpackage_src_url := "https://registry.npmjs.org/patch-package/-/patch-package-8.0.0.tgz"
	patchpackage_cmd_src := exec.Command("curl", "-L", patchpackage_src_url, "-o", "source.tar.gz")
	err = patchpackage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	patchpackage_cmd_direct := exec.Command("./binary")
	err = patchpackage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
