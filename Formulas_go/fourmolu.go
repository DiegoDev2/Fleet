package main

// Fourmolu - Formatter for Haskell source code
// Homepage: https://github.com/fourmolu/fourmolu

import (
	"fmt"
	
	"os/exec"
)

func installFourmolu() {
	// Método 1: Descargar y extraer .tar.gz
	fourmolu_tar_url := "https://github.com/fourmolu/fourmolu/archive/refs/tags/v0.16.2.0.tar.gz"
	fourmolu_cmd_tar := exec.Command("curl", "-L", fourmolu_tar_url, "-o", "package.tar.gz")
	err := fourmolu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fourmolu_zip_url := "https://github.com/fourmolu/fourmolu/archive/refs/tags/v0.16.2.0.zip"
	fourmolu_cmd_zip := exec.Command("curl", "-L", fourmolu_zip_url, "-o", "package.zip")
	err = fourmolu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fourmolu_bin_url := "https://github.com/fourmolu/fourmolu/archive/refs/tags/v0.16.2.0.bin"
	fourmolu_cmd_bin := exec.Command("curl", "-L", fourmolu_bin_url, "-o", "binary.bin")
	err = fourmolu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fourmolu_src_url := "https://github.com/fourmolu/fourmolu/archive/refs/tags/v0.16.2.0.src.tar.gz"
	fourmolu_cmd_src := exec.Command("curl", "-L", fourmolu_src_url, "-o", "source.tar.gz")
	err = fourmolu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fourmolu_cmd_direct := exec.Command("./binary")
	err = fourmolu_cmd_direct.Run()
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
