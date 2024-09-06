package main

// Shelltestrunner - Portable command-line tool for testing command-line programs
// Homepage: https://github.com/simonmichael/shelltestrunner

import (
	"fmt"
	
	"os/exec"
)

func installShelltestrunner() {
	// Método 1: Descargar y extraer .tar.gz
	shelltestrunner_tar_url := "https://hackage.haskell.org/package/shelltestrunner-1.10/shelltestrunner-1.10.tar.gz"
	shelltestrunner_cmd_tar := exec.Command("curl", "-L", shelltestrunner_tar_url, "-o", "package.tar.gz")
	err := shelltestrunner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shelltestrunner_zip_url := "https://hackage.haskell.org/package/shelltestrunner-1.10/shelltestrunner-1.10.zip"
	shelltestrunner_cmd_zip := exec.Command("curl", "-L", shelltestrunner_zip_url, "-o", "package.zip")
	err = shelltestrunner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shelltestrunner_bin_url := "https://hackage.haskell.org/package/shelltestrunner-1.10/shelltestrunner-1.10.bin"
	shelltestrunner_cmd_bin := exec.Command("curl", "-L", shelltestrunner_bin_url, "-o", "binary.bin")
	err = shelltestrunner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shelltestrunner_src_url := "https://hackage.haskell.org/package/shelltestrunner-1.10/shelltestrunner-1.10.src.tar.gz"
	shelltestrunner_cmd_src := exec.Command("curl", "-L", shelltestrunner_src_url, "-o", "source.tar.gz")
	err = shelltestrunner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shelltestrunner_cmd_direct := exec.Command("./binary")
	err = shelltestrunner_cmd_direct.Run()
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
