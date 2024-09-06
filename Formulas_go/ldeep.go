package main

// Ldeep - LDAP enumeration utility
// Homepage: https://github.com/franc-pentest/ldeep

import (
	"fmt"
	
	"os/exec"
)

func installLdeep() {
	// Método 1: Descargar y extraer .tar.gz
	ldeep_tar_url := "https://files.pythonhosted.org/packages/92/29/be36b5460cc1f4c38e4d20730a747e91106fd58e64464fca23f729135523/ldeep-1.0.58.tar.gz"
	ldeep_cmd_tar := exec.Command("curl", "-L", ldeep_tar_url, "-o", "package.tar.gz")
	err := ldeep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ldeep_zip_url := "https://files.pythonhosted.org/packages/92/29/be36b5460cc1f4c38e4d20730a747e91106fd58e64464fca23f729135523/ldeep-1.0.58.zip"
	ldeep_cmd_zip := exec.Command("curl", "-L", ldeep_zip_url, "-o", "package.zip")
	err = ldeep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ldeep_bin_url := "https://files.pythonhosted.org/packages/92/29/be36b5460cc1f4c38e4d20730a747e91106fd58e64464fca23f729135523/ldeep-1.0.58.bin"
	ldeep_cmd_bin := exec.Command("curl", "-L", ldeep_bin_url, "-o", "binary.bin")
	err = ldeep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ldeep_src_url := "https://files.pythonhosted.org/packages/92/29/be36b5460cc1f4c38e4d20730a747e91106fd58e64464fca23f729135523/ldeep-1.0.58.src.tar.gz"
	ldeep_cmd_src := exec.Command("curl", "-L", ldeep_src_url, "-o", "source.tar.gz")
	err = ldeep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ldeep_cmd_direct := exec.Command("./binary")
	err = ldeep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
