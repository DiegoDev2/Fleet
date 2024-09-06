package main

// VercelCli - Command-line interface for Vercel
// Homepage: https://vercel.com/home

import (
	"fmt"
	
	"os/exec"
)

func installVercelCli() {
	// Método 1: Descargar y extraer .tar.gz
	vercelcli_tar_url := "https://registry.npmjs.org/vercel/-/vercel-37.4.0.tgz"
	vercelcli_cmd_tar := exec.Command("curl", "-L", vercelcli_tar_url, "-o", "package.tar.gz")
	err := vercelcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vercelcli_zip_url := "https://registry.npmjs.org/vercel/-/vercel-37.4.0.tgz"
	vercelcli_cmd_zip := exec.Command("curl", "-L", vercelcli_zip_url, "-o", "package.zip")
	err = vercelcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vercelcli_bin_url := "https://registry.npmjs.org/vercel/-/vercel-37.4.0.tgz"
	vercelcli_cmd_bin := exec.Command("curl", "-L", vercelcli_bin_url, "-o", "binary.bin")
	err = vercelcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vercelcli_src_url := "https://registry.npmjs.org/vercel/-/vercel-37.4.0.tgz"
	vercelcli_cmd_src := exec.Command("curl", "-L", vercelcli_src_url, "-o", "source.tar.gz")
	err = vercelcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vercelcli_cmd_direct := exec.Command("./binary")
	err = vercelcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
