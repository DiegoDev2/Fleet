package main

// Ykman - Tool for managing your YubiKey configuration
// Homepage: https://developers.yubico.com/yubikey-manager/

import (
	"fmt"
	
	"os/exec"
)

func installYkman() {
	// Método 1: Descargar y extraer .tar.gz
	ykman_tar_url := "https://files.pythonhosted.org/packages/b5/50/9b446ca65124bbd7bc0e74304cda737248a6bceb602e5dba957114ab64df/yubikey_manager-5.5.1.tar.gz"
	ykman_cmd_tar := exec.Command("curl", "-L", ykman_tar_url, "-o", "package.tar.gz")
	err := ykman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ykman_zip_url := "https://files.pythonhosted.org/packages/b5/50/9b446ca65124bbd7bc0e74304cda737248a6bceb602e5dba957114ab64df/yubikey_manager-5.5.1.zip"
	ykman_cmd_zip := exec.Command("curl", "-L", ykman_zip_url, "-o", "package.zip")
	err = ykman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ykman_bin_url := "https://files.pythonhosted.org/packages/b5/50/9b446ca65124bbd7bc0e74304cda737248a6bceb602e5dba957114ab64df/yubikey_manager-5.5.1.bin"
	ykman_cmd_bin := exec.Command("curl", "-L", ykman_bin_url, "-o", "binary.bin")
	err = ykman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ykman_src_url := "https://files.pythonhosted.org/packages/b5/50/9b446ca65124bbd7bc0e74304cda737248a6bceb602e5dba957114ab64df/yubikey_manager-5.5.1.src.tar.gz"
	ykman_cmd_src := exec.Command("curl", "-L", ykman_src_url, "-o", "source.tar.gz")
	err = ykman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ykman_cmd_direct := exec.Command("./binary")
	err = ykman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
