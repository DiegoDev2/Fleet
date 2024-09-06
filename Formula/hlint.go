package main

// Hlint - Haskell source code suggestions
// Homepage: https://github.com/ndmitchell/hlint

import (
	"fmt"
	
	"os/exec"
)

func installHlint() {
	// Método 1: Descargar y extraer .tar.gz
	hlint_tar_url := "https://hackage.haskell.org/package/hlint-3.8/hlint-3.8.tar.gz"
	hlint_cmd_tar := exec.Command("curl", "-L", hlint_tar_url, "-o", "package.tar.gz")
	err := hlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hlint_zip_url := "https://hackage.haskell.org/package/hlint-3.8/hlint-3.8.zip"
	hlint_cmd_zip := exec.Command("curl", "-L", hlint_zip_url, "-o", "package.zip")
	err = hlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hlint_bin_url := "https://hackage.haskell.org/package/hlint-3.8/hlint-3.8.bin"
	hlint_cmd_bin := exec.Command("curl", "-L", hlint_bin_url, "-o", "binary.bin")
	err = hlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hlint_src_url := "https://hackage.haskell.org/package/hlint-3.8/hlint-3.8.src.tar.gz"
	hlint_cmd_src := exec.Command("curl", "-L", hlint_src_url, "-o", "source.tar.gz")
	err = hlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hlint_cmd_direct := exec.Command("./binary")
	err = hlint_cmd_direct.Run()
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
