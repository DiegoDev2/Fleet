package main

// AwsKeychain - Uses macOS keychain for storage of AWS credentials
// Homepage: https://github.com/pda/aws-keychain

import (
	"fmt"
	
	"os/exec"
)

func installAwsKeychain() {
	// Método 1: Descargar y extraer .tar.gz
	awskeychain_tar_url := "https://github.com/pda/aws-keychain/archive/refs/tags/v3.0.0.tar.gz"
	awskeychain_cmd_tar := exec.Command("curl", "-L", awskeychain_tar_url, "-o", "package.tar.gz")
	err := awskeychain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awskeychain_zip_url := "https://github.com/pda/aws-keychain/archive/refs/tags/v3.0.0.zip"
	awskeychain_cmd_zip := exec.Command("curl", "-L", awskeychain_zip_url, "-o", "package.zip")
	err = awskeychain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awskeychain_bin_url := "https://github.com/pda/aws-keychain/archive/refs/tags/v3.0.0.bin"
	awskeychain_cmd_bin := exec.Command("curl", "-L", awskeychain_bin_url, "-o", "binary.bin")
	err = awskeychain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awskeychain_src_url := "https://github.com/pda/aws-keychain/archive/refs/tags/v3.0.0.src.tar.gz"
	awskeychain_cmd_src := exec.Command("curl", "-L", awskeychain_src_url, "-o", "source.tar.gz")
	err = awskeychain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awskeychain_cmd_direct := exec.Command("./binary")
	err = awskeychain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
