package main

// DhallYaml - Convert between Dhall and YAML
// Homepage: https://github.com/dhall-lang/dhall-haskell/tree/main/dhall-yaml

import (
	"fmt"
	
	"os/exec"
)

func installDhallYaml() {
	// Método 1: Descargar y extraer .tar.gz
	dhallyaml_tar_url := "https://hackage.haskell.org/package/dhall-yaml-1.2.12/dhall-yaml-1.2.12.tar.gz"
	dhallyaml_cmd_tar := exec.Command("curl", "-L", dhallyaml_tar_url, "-o", "package.tar.gz")
	err := dhallyaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhallyaml_zip_url := "https://hackage.haskell.org/package/dhall-yaml-1.2.12/dhall-yaml-1.2.12.zip"
	dhallyaml_cmd_zip := exec.Command("curl", "-L", dhallyaml_zip_url, "-o", "package.zip")
	err = dhallyaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhallyaml_bin_url := "https://hackage.haskell.org/package/dhall-yaml-1.2.12/dhall-yaml-1.2.12.bin"
	dhallyaml_cmd_bin := exec.Command("curl", "-L", dhallyaml_bin_url, "-o", "binary.bin")
	err = dhallyaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhallyaml_src_url := "https://hackage.haskell.org/package/dhall-yaml-1.2.12/dhall-yaml-1.2.12.src.tar.gz"
	dhallyaml_cmd_src := exec.Command("curl", "-L", dhallyaml_src_url, "-o", "source.tar.gz")
	err = dhallyaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhallyaml_cmd_direct := exec.Command("./binary")
	err = dhallyaml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.6")
exec.Command("latte", "install", "ghc@9.6")
}
