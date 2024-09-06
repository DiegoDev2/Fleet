package main

// InstallPeerdeps - CLI to automatically install peerDeps
// Homepage: https://github.com/nathanhleung/install-peerdeps

import (
	"fmt"
	
	"os/exec"
)

func installInstallPeerdeps() {
	// Método 1: Descargar y extraer .tar.gz
	installpeerdeps_tar_url := "https://registry.npmjs.org/install-peerdeps/-/install-peerdeps-3.0.3.tgz"
	installpeerdeps_cmd_tar := exec.Command("curl", "-L", installpeerdeps_tar_url, "-o", "package.tar.gz")
	err := installpeerdeps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	installpeerdeps_zip_url := "https://registry.npmjs.org/install-peerdeps/-/install-peerdeps-3.0.3.tgz"
	installpeerdeps_cmd_zip := exec.Command("curl", "-L", installpeerdeps_zip_url, "-o", "package.zip")
	err = installpeerdeps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	installpeerdeps_bin_url := "https://registry.npmjs.org/install-peerdeps/-/install-peerdeps-3.0.3.tgz"
	installpeerdeps_cmd_bin := exec.Command("curl", "-L", installpeerdeps_bin_url, "-o", "binary.bin")
	err = installpeerdeps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	installpeerdeps_src_url := "https://registry.npmjs.org/install-peerdeps/-/install-peerdeps-3.0.3.tgz"
	installpeerdeps_cmd_src := exec.Command("curl", "-L", installpeerdeps_src_url, "-o", "source.tar.gz")
	err = installpeerdeps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	installpeerdeps_cmd_direct := exec.Command("./binary")
	err = installpeerdeps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
