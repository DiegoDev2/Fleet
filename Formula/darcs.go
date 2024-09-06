package main

// Darcs - Distributed version control system that tracks changes, via Haskell
// Homepage: https://darcs.net/

import (
	"fmt"
	
	"os/exec"
)

func installDarcs() {
	// Método 1: Descargar y extraer .tar.gz
	darcs_tar_url := "https://hackage.haskell.org/package/darcs-2.18.3/darcs-2.18.3.tar.gz"
	darcs_cmd_tar := exec.Command("curl", "-L", darcs_tar_url, "-o", "package.tar.gz")
	err := darcs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darcs_zip_url := "https://hackage.haskell.org/package/darcs-2.18.3/darcs-2.18.3.zip"
	darcs_cmd_zip := exec.Command("curl", "-L", darcs_zip_url, "-o", "package.zip")
	err = darcs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darcs_bin_url := "https://hackage.haskell.org/package/darcs-2.18.3/darcs-2.18.3.bin"
	darcs_cmd_bin := exec.Command("curl", "-L", darcs_bin_url, "-o", "binary.bin")
	err = darcs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darcs_src_url := "https://hackage.haskell.org/package/darcs-2.18.3/darcs-2.18.3.src.tar.gz"
	darcs_cmd_src := exec.Command("curl", "-L", darcs_src_url, "-o", "source.tar.gz")
	err = darcs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darcs_cmd_direct := exec.Command("./binary")
	err = darcs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@9.8")
	exec.Command("latte", "install", "ghc@9.8").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
