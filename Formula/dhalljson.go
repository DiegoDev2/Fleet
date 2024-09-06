package main

// DhallJson - Dhall to JSON compiler and a Dhall to YAML compiler
// Homepage: https://github.com/dhall-lang/dhall-haskell/tree/master/dhall-json

import (
	"fmt"
	
	"os/exec"
)

func installDhallJson() {
	// Método 1: Descargar y extraer .tar.gz
	dhalljson_tar_url := "https://hackage.haskell.org/package/dhall-json-1.7.12/dhall-json-1.7.12.tar.gz"
	dhalljson_cmd_tar := exec.Command("curl", "-L", dhalljson_tar_url, "-o", "package.tar.gz")
	err := dhalljson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhalljson_zip_url := "https://hackage.haskell.org/package/dhall-json-1.7.12/dhall-json-1.7.12.zip"
	dhalljson_cmd_zip := exec.Command("curl", "-L", dhalljson_zip_url, "-o", "package.zip")
	err = dhalljson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhalljson_bin_url := "https://hackage.haskell.org/package/dhall-json-1.7.12/dhall-json-1.7.12.bin"
	dhalljson_cmd_bin := exec.Command("curl", "-L", dhalljson_bin_url, "-o", "binary.bin")
	err = dhalljson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhalljson_src_url := "https://hackage.haskell.org/package/dhall-json-1.7.12/dhall-json-1.7.12.src.tar.gz"
	dhalljson_cmd_src := exec.Command("curl", "-L", dhalljson_src_url, "-o", "source.tar.gz")
	err = dhalljson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhalljson_cmd_direct := exec.Command("./binary")
	err = dhalljson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@9.6")
	exec.Command("latte", "install", "ghc@9.6").Run()
}
