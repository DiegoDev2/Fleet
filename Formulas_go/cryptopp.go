package main

// Cryptopp - Free C++ class library of cryptographic schemes
// Homepage: https://cryptopp.com/

import (
	"fmt"
	
	"os/exec"
)

func installCryptopp() {
	// Método 1: Descargar y extraer .tar.gz
	cryptopp_tar_url := "https://cryptopp.com/cryptopp890.zip"
	cryptopp_cmd_tar := exec.Command("curl", "-L", cryptopp_tar_url, "-o", "package.tar.gz")
	err := cryptopp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cryptopp_zip_url := "https://cryptopp.com/cryptopp890.zip"
	cryptopp_cmd_zip := exec.Command("curl", "-L", cryptopp_zip_url, "-o", "package.zip")
	err = cryptopp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cryptopp_bin_url := "https://cryptopp.com/cryptopp890.zip"
	cryptopp_cmd_bin := exec.Command("curl", "-L", cryptopp_bin_url, "-o", "binary.bin")
	err = cryptopp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cryptopp_src_url := "https://cryptopp.com/cryptopp890.zip"
	cryptopp_cmd_src := exec.Command("curl", "-L", cryptopp_src_url, "-o", "source.tar.gz")
	err = cryptopp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cryptopp_cmd_direct := exec.Command("./binary")
	err = cryptopp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
