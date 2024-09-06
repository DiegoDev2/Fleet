package main

// PscPackage - Package manager for PureScript based on package sets
// Homepage: https://psc-package.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installPscPackage() {
	// Método 1: Descargar y extraer .tar.gz
	pscpackage_tar_url := "https://github.com/purescript/psc-package/archive/refs/tags/v0.6.2.tar.gz"
	pscpackage_cmd_tar := exec.Command("curl", "-L", pscpackage_tar_url, "-o", "package.tar.gz")
	err := pscpackage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pscpackage_zip_url := "https://github.com/purescript/psc-package/archive/refs/tags/v0.6.2.zip"
	pscpackage_cmd_zip := exec.Command("curl", "-L", pscpackage_zip_url, "-o", "package.zip")
	err = pscpackage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pscpackage_bin_url := "https://github.com/purescript/psc-package/archive/refs/tags/v0.6.2.bin"
	pscpackage_cmd_bin := exec.Command("curl", "-L", pscpackage_bin_url, "-o", "binary.bin")
	err = pscpackage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pscpackage_src_url := "https://github.com/purescript/psc-package/archive/refs/tags/v0.6.2.src.tar.gz"
	pscpackage_cmd_src := exec.Command("curl", "-L", pscpackage_src_url, "-o", "source.tar.gz")
	err = pscpackage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pscpackage_cmd_direct := exec.Command("./binary")
	err = pscpackage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.8")
exec.Command("latte", "install", "ghc@9.8")
	fmt.Println("Instalando dependencia: purescript")
exec.Command("latte", "install", "purescript")
}
