package main

// Ormolu - Formatter for Haskell source code
// Homepage: https://github.com/tweag/ormolu

import (
	"fmt"
	
	"os/exec"
)

func installOrmolu() {
	// Método 1: Descargar y extraer .tar.gz
	ormolu_tar_url := "https://github.com/tweag/ormolu/archive/refs/tags/0.7.7.0.tar.gz"
	ormolu_cmd_tar := exec.Command("curl", "-L", ormolu_tar_url, "-o", "package.tar.gz")
	err := ormolu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ormolu_zip_url := "https://github.com/tweag/ormolu/archive/refs/tags/0.7.7.0.zip"
	ormolu_cmd_zip := exec.Command("curl", "-L", ormolu_zip_url, "-o", "package.zip")
	err = ormolu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ormolu_bin_url := "https://github.com/tweag/ormolu/archive/refs/tags/0.7.7.0.bin"
	ormolu_cmd_bin := exec.Command("curl", "-L", ormolu_bin_url, "-o", "binary.bin")
	err = ormolu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ormolu_src_url := "https://github.com/tweag/ormolu/archive/refs/tags/0.7.7.0.src.tar.gz"
	ormolu_cmd_src := exec.Command("curl", "-L", ormolu_src_url, "-o", "source.tar.gz")
	err = ormolu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ormolu_cmd_direct := exec.Command("./binary")
	err = ormolu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
}
