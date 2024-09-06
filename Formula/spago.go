package main

// Spago - PureScript package manager and build tool
// Homepage: https://github.com/purescript/spago

import (
	"fmt"
	
	"os/exec"
)

func installSpago() {
	// Método 1: Descargar y extraer .tar.gz
	spago_tar_url := "https://github.com/purescript/spago/archive/refs/tags/0.21.0.tar.gz"
	spago_cmd_tar := exec.Command("curl", "-L", spago_tar_url, "-o", "package.tar.gz")
	err := spago_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spago_zip_url := "https://github.com/purescript/spago/archive/refs/tags/0.21.0.zip"
	spago_cmd_zip := exec.Command("curl", "-L", spago_zip_url, "-o", "package.zip")
	err = spago_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spago_bin_url := "https://github.com/purescript/spago/archive/refs/tags/0.21.0.bin"
	spago_cmd_bin := exec.Command("curl", "-L", spago_bin_url, "-o", "binary.bin")
	err = spago_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spago_src_url := "https://github.com/purescript/spago/archive/refs/tags/0.21.0.src.tar.gz"
	spago_cmd_src := exec.Command("curl", "-L", spago_src_url, "-o", "source.tar.gz")
	err = spago_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spago_cmd_direct := exec.Command("./binary")
	err = spago_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghc@9.2")
	exec.Command("latte", "install", "ghc@9.2").Run()
	fmt.Println("Instalando dependencia: haskell-stack")
	exec.Command("latte", "install", "haskell-stack").Run()
	fmt.Println("Instalando dependencia: purescript")
	exec.Command("latte", "install", "purescript").Run()
}
