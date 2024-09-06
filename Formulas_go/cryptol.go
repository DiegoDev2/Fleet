package main

// Cryptol - Domain-specific language for specifying cryptographic algorithms
// Homepage: https://www.cryptol.net/

import (
	"fmt"
	
	"os/exec"
)

func installCryptol() {
	// Método 1: Descargar y extraer .tar.gz
	cryptol_tar_url := "https://hackage.haskell.org/package/cryptol-3.2.0/cryptol-3.2.0.tar.gz"
	cryptol_cmd_tar := exec.Command("curl", "-L", cryptol_tar_url, "-o", "package.tar.gz")
	err := cryptol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cryptol_zip_url := "https://hackage.haskell.org/package/cryptol-3.2.0/cryptol-3.2.0.zip"
	cryptol_cmd_zip := exec.Command("curl", "-L", cryptol_zip_url, "-o", "package.zip")
	err = cryptol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cryptol_bin_url := "https://hackage.haskell.org/package/cryptol-3.2.0/cryptol-3.2.0.bin"
	cryptol_cmd_bin := exec.Command("curl", "-L", cryptol_bin_url, "-o", "binary.bin")
	err = cryptol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cryptol_src_url := "https://hackage.haskell.org/package/cryptol-3.2.0/cryptol-3.2.0.src.tar.gz"
	cryptol_cmd_src := exec.Command("curl", "-L", cryptol_src_url, "-o", "source.tar.gz")
	err = cryptol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cryptol_cmd_direct := exec.Command("./binary")
	err = cryptol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.8")
exec.Command("latte", "install", "ghc@9.8")
	fmt.Println("Instalando dependencia: z3")
exec.Command("latte", "install", "z3")
}
