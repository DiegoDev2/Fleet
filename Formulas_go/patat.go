package main

// Patat - Terminal-based presentations using Pandoc
// Homepage: https://github.com/jaspervdj/patat

import (
	"fmt"
	
	"os/exec"
)

func installPatat() {
	// Método 1: Descargar y extraer .tar.gz
	patat_tar_url := "https://hackage.haskell.org/package/patat-0.12.0.0/patat-0.12.0.0.tar.gz"
	patat_cmd_tar := exec.Command("curl", "-L", patat_tar_url, "-o", "package.tar.gz")
	err := patat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	patat_zip_url := "https://hackage.haskell.org/package/patat-0.12.0.0/patat-0.12.0.0.zip"
	patat_cmd_zip := exec.Command("curl", "-L", patat_zip_url, "-o", "package.zip")
	err = patat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	patat_bin_url := "https://hackage.haskell.org/package/patat-0.12.0.0/patat-0.12.0.0.bin"
	patat_cmd_bin := exec.Command("curl", "-L", patat_bin_url, "-o", "binary.bin")
	err = patat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	patat_src_url := "https://hackage.haskell.org/package/patat-0.12.0.0/patat-0.12.0.0.src.tar.gz"
	patat_cmd_src := exec.Command("curl", "-L", patat_src_url, "-o", "source.tar.gz")
	err = patat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	patat_cmd_direct := exec.Command("./binary")
	err = patat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
