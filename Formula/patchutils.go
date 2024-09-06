package main

// Patchutils - Small collection of programs that operate on patch files
// Homepage: http://cyberelk.net/tim/software/patchutils/

import (
	"fmt"
	
	"os/exec"
)

func installPatchutils() {
	// Método 1: Descargar y extraer .tar.gz
	patchutils_tar_url := "http://cyberelk.net/tim/data/patchutils/stable/patchutils-0.4.2.tar.xz"
	patchutils_cmd_tar := exec.Command("curl", "-L", patchutils_tar_url, "-o", "package.tar.gz")
	err := patchutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	patchutils_zip_url := "http://cyberelk.net/tim/data/patchutils/stable/patchutils-0.4.2.tar.xz"
	patchutils_cmd_zip := exec.Command("curl", "-L", patchutils_zip_url, "-o", "package.zip")
	err = patchutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	patchutils_bin_url := "http://cyberelk.net/tim/data/patchutils/stable/patchutils-0.4.2.tar.xz"
	patchutils_cmd_bin := exec.Command("curl", "-L", patchutils_bin_url, "-o", "binary.bin")
	err = patchutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	patchutils_src_url := "http://cyberelk.net/tim/data/patchutils/stable/patchutils-0.4.2.tar.xz"
	patchutils_cmd_src := exec.Command("curl", "-L", patchutils_src_url, "-o", "source.tar.gz")
	err = patchutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	patchutils_cmd_direct := exec.Command("./binary")
	err = patchutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: docbook")
	exec.Command("latte", "install", "docbook").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
}
