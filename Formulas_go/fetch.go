package main

// Fetch - Download assets from a commit, branch, or tag of GitHub repositories
// Homepage: https://www.gruntwork.io/

import (
	"fmt"
	
	"os/exec"
)

func installFetch() {
	// Método 1: Descargar y extraer .tar.gz
	fetch_tar_url := "https://github.com/gruntwork-io/fetch/archive/refs/tags/v0.4.6.tar.gz"
	fetch_cmd_tar := exec.Command("curl", "-L", fetch_tar_url, "-o", "package.tar.gz")
	err := fetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fetch_zip_url := "https://github.com/gruntwork-io/fetch/archive/refs/tags/v0.4.6.zip"
	fetch_cmd_zip := exec.Command("curl", "-L", fetch_zip_url, "-o", "package.zip")
	err = fetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fetch_bin_url := "https://github.com/gruntwork-io/fetch/archive/refs/tags/v0.4.6.bin"
	fetch_cmd_bin := exec.Command("curl", "-L", fetch_bin_url, "-o", "binary.bin")
	err = fetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fetch_src_url := "https://github.com/gruntwork-io/fetch/archive/refs/tags/v0.4.6.src.tar.gz"
	fetch_cmd_src := exec.Command("curl", "-L", fetch_src_url, "-o", "source.tar.gz")
	err = fetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fetch_cmd_direct := exec.Command("./binary")
	err = fetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
