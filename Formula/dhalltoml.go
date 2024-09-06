package main

// DhallToml - Convert between Dhall and Toml
// Homepage: https://github.com/dhall-lang/dhall-haskell/tree/main/dhall-toml

import (
	"fmt"
	
	"os/exec"
)

func installDhallToml() {
	// Método 1: Descargar y extraer .tar.gz
	dhalltoml_tar_url := "https://hackage.haskell.org/package/dhall-toml-1.0.3/dhall-toml-1.0.3.tar.gz"
	dhalltoml_cmd_tar := exec.Command("curl", "-L", dhalltoml_tar_url, "-o", "package.tar.gz")
	err := dhalltoml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhalltoml_zip_url := "https://hackage.haskell.org/package/dhall-toml-1.0.3/dhall-toml-1.0.3.zip"
	dhalltoml_cmd_zip := exec.Command("curl", "-L", dhalltoml_zip_url, "-o", "package.zip")
	err = dhalltoml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhalltoml_bin_url := "https://hackage.haskell.org/package/dhall-toml-1.0.3/dhall-toml-1.0.3.bin"
	dhalltoml_cmd_bin := exec.Command("curl", "-L", dhalltoml_bin_url, "-o", "binary.bin")
	err = dhalltoml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhalltoml_src_url := "https://hackage.haskell.org/package/dhall-toml-1.0.3/dhall-toml-1.0.3.src.tar.gz"
	dhalltoml_cmd_src := exec.Command("curl", "-L", dhalltoml_src_url, "-o", "source.tar.gz")
	err = dhalltoml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhalltoml_cmd_direct := exec.Command("./binary")
	err = dhalltoml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@9.8")
	exec.Command("latte", "install", "ghc@9.8").Run()
}
