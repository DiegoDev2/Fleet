package main

// AtomistCli - Unified command-line tool for interacting with Atomist services
// Homepage: https://github.com/atomist/cli

import (
	"fmt"
	
	"os/exec"
)

func installAtomistCli() {
	// Método 1: Descargar y extraer .tar.gz
	atomistcli_tar_url := "https://registry.npmjs.org/@atomist/cli/-/cli-1.8.0.tgz"
	atomistcli_cmd_tar := exec.Command("curl", "-L", atomistcli_tar_url, "-o", "package.tar.gz")
	err := atomistcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atomistcli_zip_url := "https://registry.npmjs.org/@atomist/cli/-/cli-1.8.0.tgz"
	atomistcli_cmd_zip := exec.Command("curl", "-L", atomistcli_zip_url, "-o", "package.zip")
	err = atomistcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atomistcli_bin_url := "https://registry.npmjs.org/@atomist/cli/-/cli-1.8.0.tgz"
	atomistcli_cmd_bin := exec.Command("curl", "-L", atomistcli_bin_url, "-o", "binary.bin")
	err = atomistcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atomistcli_src_url := "https://registry.npmjs.org/@atomist/cli/-/cli-1.8.0.tgz"
	atomistcli_cmd_src := exec.Command("curl", "-L", atomistcli_src_url, "-o", "source.tar.gz")
	err = atomistcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atomistcli_cmd_direct := exec.Command("./binary")
	err = atomistcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: macos-term-size")
exec.Command("latte", "install", "macos-term-size")
}
