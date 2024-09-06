package main

// Echidna - Ethereum smart contract fuzzer
// Homepage: https://github.com/crytic/echidna

import (
	"fmt"
	
	"os/exec"
)

func installEchidna() {
	// Método 1: Descargar y extraer .tar.gz
	echidna_tar_url := "https://github.com/crytic/echidna/archive/refs/tags/v2.2.4.tar.gz"
	echidna_cmd_tar := exec.Command("curl", "-L", echidna_tar_url, "-o", "package.tar.gz")
	err := echidna_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	echidna_zip_url := "https://github.com/crytic/echidna/archive/refs/tags/v2.2.4.zip"
	echidna_cmd_zip := exec.Command("curl", "-L", echidna_zip_url, "-o", "package.zip")
	err = echidna_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	echidna_bin_url := "https://github.com/crytic/echidna/archive/refs/tags/v2.2.4.bin"
	echidna_cmd_bin := exec.Command("curl", "-L", echidna_bin_url, "-o", "binary.bin")
	err = echidna_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	echidna_src_url := "https://github.com/crytic/echidna/archive/refs/tags/v2.2.4.src.tar.gz"
	echidna_cmd_src := exec.Command("curl", "-L", echidna_src_url, "-o", "source.tar.gz")
	err = echidna_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	echidna_cmd_direct := exec.Command("./binary")
	err = echidna_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghc@9.4")
	exec.Command("latte", "install", "ghc@9.4").Run()
	fmt.Println("Instalando dependencia: haskell-stack")
	exec.Command("latte", "install", "haskell-stack").Run()
	fmt.Println("Instalando dependencia: truffle")
	exec.Command("latte", "install", "truffle").Run()
	fmt.Println("Instalando dependencia: crytic-compile")
	exec.Command("latte", "install", "crytic-compile").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libff")
	exec.Command("latte", "install", "libff").Run()
	fmt.Println("Instalando dependencia: secp256k1")
	exec.Command("latte", "install", "secp256k1").Run()
	fmt.Println("Instalando dependencia: slither-analyzer")
	exec.Command("latte", "install", "slither-analyzer").Run()
}
