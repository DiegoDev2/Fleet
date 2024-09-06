package main

// Ncc - Compile a Node.js project into a single file
// Homepage: https://github.com/vercel/ncc

import (
	"fmt"
	
	"os/exec"
)

func installNcc() {
	// Método 1: Descargar y extraer .tar.gz
	ncc_tar_url := "https://registry.npmjs.org/@vercel/ncc/-/ncc-0.38.1.tgz"
	ncc_cmd_tar := exec.Command("curl", "-L", ncc_tar_url, "-o", "package.tar.gz")
	err := ncc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncc_zip_url := "https://registry.npmjs.org/@vercel/ncc/-/ncc-0.38.1.tgz"
	ncc_cmd_zip := exec.Command("curl", "-L", ncc_zip_url, "-o", "package.zip")
	err = ncc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncc_bin_url := "https://registry.npmjs.org/@vercel/ncc/-/ncc-0.38.1.tgz"
	ncc_cmd_bin := exec.Command("curl", "-L", ncc_bin_url, "-o", "binary.bin")
	err = ncc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncc_src_url := "https://registry.npmjs.org/@vercel/ncc/-/ncc-0.38.1.tgz"
	ncc_cmd_src := exec.Command("curl", "-L", ncc_src_url, "-o", "source.tar.gz")
	err = ncc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncc_cmd_direct := exec.Command("./binary")
	err = ncc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
